import * as cdk from "aws-cdk-lib";
import * as apigwv2 from "aws-cdk-lib/aws-apigatewayv2";
import * as authorizers from "aws-cdk-lib/aws-apigatewayv2-authorizers";
import * as integrations from "aws-cdk-lib/aws-apigatewayv2-integrations";
import * as iam from "aws-cdk-lib/aws-iam";
import * as lambda from "aws-cdk-lib/aws-lambda";
import * as logs from "aws-cdk-lib/aws-logs";
import * as s3 from "aws-cdk-lib/aws-s3";
import { Construct } from "constructs";

/**
 * LambdaApi コンストラクタプロパティ
 *
 * @property envName - 環境名。リソースの命名に使用する
 * @property lambdaMemorySize - API Lambda 関数のメモリサイズ (MB)
 * @property logRetentionDays - CloudWatch Logs のロググループ保持日数
 * @property artifactsBucketName - Lambda ビルド成果物を格納する S3 バケット名
 * @property envVars - API Lambda に渡す追加の環境変数
 */
interface LambdaApiProps {
  readonly envName: string;
  readonly lambdaMemorySize: number;
  readonly logRetentionDays: number;
  readonly artifactsBucketName: string;
  readonly envVars?: { [key: string]: string };
}

/**
 * API Lambda + Lambda Authorizer + HTTP API (API Gateway v2) コンストラクト
 *
 * arm64 / provided.al2023 で API Lambda と Authorizer Lambda を定義し、
 * x-api-key ヘッダーを Secrets Manager で検証する Lambda Authorizer 付き HTTP API を構築する。
 */
export class LambdaApi extends Construct {
  /** API Lambda 関数 */
  readonly fn: lambda.Function;

  /** API Lambda 実行ロール */
  readonly executionRole: iam.Role;

  /** Authorizer Lambda 実行ロール */
  readonly authorizerRole: iam.Role;

  /** Authorizer Lambda 関数 */
  readonly authorizerFn: lambda.Function;

  /** HTTP API */
  readonly httpApi: apigwv2.HttpApi;

  constructor(scope: Construct, id: string, props: LambdaApiProps) {
    super(scope, id);

    const artifactsBucket = s3.Bucket.fromBucketName(
      this,
      "ArtifactsBucket",
      props.artifactsBucketName,
    );

    // ---- Authorizer Lambda ----

    this.authorizerRole = new iam.Role(this, "AuthorizerRole", {
      roleName: `${props.envName}-pollen-tracker-authorizer-role`,
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      // ref: https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AWSLambdaBasicExecutionRole.html
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole",
        ),
      ],
    });

    const authorizerLogGroup = new logs.LogGroup(this, "AuthorizerLogGroup", {
      logGroupName: `/aws/lambda/${props.envName}-pollen-tracker-authorizer`,
      retention: props.logRetentionDays as logs.RetentionDays,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });

    this.authorizerFn = new lambda.Function(this, "AuthorizerFunction", {
      functionName: `${props.envName}-pollen-tracker-authorizer`,
      description:
        "Lambda Authorizer — validates x-api-key header against Secrets Manager",
      runtime: lambda.Runtime.PROVIDED_AL2023,
      architecture: lambda.Architecture.ARM_64,
      handler: "bootstrap",
      // make upload-authorizer で s3://{artifactsBucketName}/artifacts/authorizer/bootstrap.zip にアップロードされたバイナリを参照する
      code: lambda.Code.fromBucket(
        artifactsBucket,
        "artifacts/authorizer/bootstrap.zip",
      ),
      role: this.authorizerRole,
      memorySize: 128,
      // Lambda Authorizer の標準タイムアウトは 10 秒。推奨は 1 秒以内だが保守的に 5 秒に設定
      timeout: cdk.Duration.seconds(5),
      environment: {
        ENV: props.envName,
        PROJECT: "pollen-tracker",
        SERVICE: "authorizer",
      },
      logGroup: authorizerLogGroup,
    });

    // ---- HTTP Lambda Authorizer ----

    const authorizer = new authorizers.HttpLambdaAuthorizer(
      "ApiKeyAuthorizer",
      this.authorizerFn,
      {
        // ref: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-lambda-authorizer.html
        identitySource: ["$request.header.x-api-key"],
        resultsCacheTtl: cdk.Duration.minutes(5),
        responseTypes: [authorizers.HttpLambdaResponseType.SIMPLE],
      },
    );

    // ---- API Lambda ----

    this.executionRole = new iam.Role(this, "ExecutionRole", {
      roleName: `${props.envName}-pollen-tracker-api-role`,
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
      // ref: https://docs.aws.amazon.com/ja_jp/aws-managed-policy/latest/reference/AWSLambdaBasicExecutionRole.html
      managedPolicies: [
        iam.ManagedPolicy.fromAwsManagedPolicyName(
          "service-role/AWSLambdaBasicExecutionRole",
        ),
      ],
    });

    const apiLogGroup = new logs.LogGroup(this, "ApiLogGroup", {
      logGroupName: `/aws/lambda/${props.envName}-pollen-tracker-api`,
      retention: props.logRetentionDays as logs.RetentionDays,
      removalPolicy: cdk.RemovalPolicy.DESTROY,
    });

    this.fn = new lambda.Function(this, "Function", {
      functionName: `${props.envName}-pollen-tracker-api`,
      description:
        "Receives HTTP requests via API Gateway v2 and serves pollen tracker API",
      runtime: lambda.Runtime.PROVIDED_AL2023,
      architecture: lambda.Architecture.ARM_64,
      handler: "bootstrap",
      // make upload-api で s3://{artifactsBucketName}/artifacts/api/bootstrap.zip にアップロードされたバイナリを参照する
      code: lambda.Code.fromBucket(
        artifactsBucket,
        "artifacts/api/bootstrap.zip",
      ),
      role: this.executionRole,
      // HTTP サーバーとして実装された Go バイナリを Lambda で動かすための Web Adapter
      // ref: https://github.com/awslabs/aws-lambda-web-adapter
      layers: [
        lambda.LayerVersion.fromLayerVersionArn(
          this,
          "LambdaWebAdapterLayer",
          "arn:aws:lambda:ap-northeast-1:753240598075:layer:LambdaAdapterLayerArm64:24",
        ),
      ],
      memorySize: props.lambdaMemorySize,
      // HTTP API (v2) の統合タイムアウト上限は 30 秒のため、1 秒のバッファを設けて 29 秒に設定する
      // ref: https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-quotas.html
      timeout: cdk.Duration.seconds(29),
      environment: {
        ENV: props.envName,
        PORT: "8080",
        PROJECT: "pollen-tracker",
        SERVICE: "api-server",
        // LWA の readiness check に /health を使用する (デフォルトの / はエラーハンドラを経由するため)
        AWS_LWA_READINESS_CHECK_PATH: "/health",
        ...props.envVars,
      },
      logGroup: apiLogGroup,
    });

    // ---- HTTP API ----

    this.httpApi = new apigwv2.HttpApi(this, "HttpApi", {
      apiName: `${props.envName}-pollen-tracker`,
      description: "HTTP API for pollen tracker",
      defaultAuthorizer: authorizer,
      defaultIntegration: new integrations.HttpLambdaIntegration(
        "Integration",
        this.fn,
      ),
    });

    new cdk.CfnOutput(scope, "ApiEndpointUrl", {
      value: this.httpApi.apiEndpoint,
      description: "API Gateway HTTP API endpoint URL",
    });
  }
}
