import * as cdk from "aws-cdk-lib/core";
import { Construct } from "constructs";
import { EnvConfig } from "../../config/env-config";

export interface PollenStackProps extends cdk.StackProps {
  readonly config: EnvConfig;
}

export class PollenStack extends cdk.Stack {
  constructor(scope: Construct, id: string, props: PollenStackProps) {
    super(scope, id, props);
  }
}
