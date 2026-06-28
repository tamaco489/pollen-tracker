import * as cdk from "aws-cdk-lib/core";
import { Construct } from "constructs";

import { EnvConfig } from "../../config/env-config";
import { LambdaApi } from "../constructs/lambda-api";
import { ManagedSecret } from "../constructs/managed-secret";

/**
 * PollenStack のコンストラクタプロパティ
 *
 * @property config - 環境設定。詳細は {@link EnvConfig} を参照
 */
export interface PollenStackProps extends cdk.StackProps {
  readonly config: EnvConfig;
}

/**
 * 花粉トラッカー API のメインスタック
 *
 * 各リソースは lib/constructs/ 配下の L3 カスタムコンストラクトに分割して組み立てる。
 */
export class PollenStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: PollenStackProps) {
    super(scope, id, props);

    const lambdaApi = new LambdaApi(this, "LambdaApi", {
      envName: props.config.envName,
      lambdaMemorySize: props.config.lambdaMemorySize,
      logRetentionDays: props.config.logRetentionDays,
      artifactsBucketName: props.config.artifactsBucketName,
    });

    const pollenApiKeySecret = new ManagedSecret(this, "PollenApiKeySecret", {
      secretName: `${props.config.envName}/pollen-tracker/pollen-api-key`,
      description: "Google Pollen API キー",
      lambdaRole: lambdaApi.executionRole,
    });

    const tursoSecret = new ManagedSecret(this, "TursoSecret", {
      secretName: `${props.config.envName}/pollen-tracker/turso`,
      description: "Turso 接続情報 (URL / AUTH_TOKEN)",
      lambdaRole: lambdaApi.executionRole,
    });

    const xApiKeySecret = new ManagedSecret(this, "XApiKeySecret", {
      secretName: `${props.config.envName}/pollen-tracker/api-key`,
      description: "Lambda Authorizer が検証する x-api-key",
      lambdaRole: lambdaApi.authorizerRole,
    });

    // ManagedSecret の ARN を Lambda 環境変数として注入
    lambdaApi.fn.addEnvironment(
      "SECRET_ARN_POLLEN",
      pollenApiKeySecret.secretArn,
    );
    lambdaApi.fn.addEnvironment("SECRET_ARN_TURSO", tursoSecret.secretArn);
    lambdaApi.authorizerFn.addEnvironment(
      "SECRET_ARN",
      xApiKeySecret.secretArn,
    );
  }
}
