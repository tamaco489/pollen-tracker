import * as cdk from "aws-cdk-lib/core";
import { Construct } from "constructs";

import { EnvConfig } from "../../config/env-config";
import { LambdaApi } from "../constructs/lambda-api";

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

    new LambdaApi(this, "LambdaApi", {
      envName: props.config.envName,
      lambdaMemorySize: props.config.lambdaMemorySize,
      logRetentionDays: props.config.logRetentionDays,
      artifactsBucketName: props.config.artifactsBucketName,
      secretArn: props.config.secretArn,
    });
  }
}
