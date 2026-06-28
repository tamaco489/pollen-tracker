import * as cdk from "aws-cdk-lib/core";
import * as iam from "aws-cdk-lib/aws-iam";
import { Match, Template } from "aws-cdk-lib/assertions";
import { ManagedSecret } from "../../lib/constructs/managed-secret";

describe("ManagedSecret", () => {
  let template: Template;

  beforeAll(() => {
    const app = new cdk.App();
    const stack = new cdk.Stack(app, "TestStack");
    const role = new iam.Role(stack, "TestRole", {
      assumedBy: new iam.ServicePrincipal("lambda.amazonaws.com"),
    });
    new ManagedSecret(stack, "TestSecret", {
      secretName: "test/pollen-tracker/turso/config",
      description: "テスト用シークレット",
      lambdaRole: role,
    });
    template = Template.fromStack(stack);
  });

  test("指定した名前と説明で Secret が生成される", () => {
    template.hasResourceProperties("AWS::SecretsManager::Secret", {
      Name: "test/pollen-tracker/turso/config",
      Description: "テスト用シークレット",
    });
  });

  test("Lambda ロールに secretsmanager:GetSecretValue 権限が付与される", () => {
    template.hasResourceProperties("AWS::IAM::Policy", {
      PolicyDocument: {
        Statement: Match.arrayWith([
          Match.objectLike({
            Action: "secretsmanager:GetSecretValue",
            Effect: "Allow",
          }),
        ]),
      },
    });
  });
});
