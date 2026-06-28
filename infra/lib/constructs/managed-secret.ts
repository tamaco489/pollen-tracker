import * as iam from "aws-cdk-lib/aws-iam";
import * as secretsmanager from "aws-cdk-lib/aws-secretsmanager";
import { Construct } from "constructs";

/**
 * ManagedSecret コンストラクタプロパティ
 *
 * @property secretName - Secrets Manager に登録するシークレット名
 * @property description - シークレットの説明
 * @property lambdaRole - GetSecretValue 権限を付与する Lambda 実行ロール
 */
interface ManagedSecretProps {
  readonly secretName: string;
  readonly description: string;
  readonly lambdaRole: iam.IRole;
}

/**
 * Secrets Manager シークレット + IAM ポリシー付与コンストラクト
 *
 * Secret を生成し、指定した Lambda 実行ロールに secretsmanager:GetSecretValue 権限を付与する。
 */
export class ManagedSecret extends Construct {
  /** 生成されたシークレットの ARN */
  readonly secretArn: string;

  constructor(scope: Construct, id: string, props: ManagedSecretProps) {
    super(scope, id);

    const secret = new secretsmanager.Secret(this, "Secret", {
      secretName: props.secretName,
      description: props.description,
    });

    props.lambdaRole.addToPrincipalPolicy(
      new iam.PolicyStatement({
        actions: ["secretsmanager:GetSecretValue"],
        resources: [secret.secretArn],
      }),
    );

    this.secretArn = secret.secretArn;
  }
}
