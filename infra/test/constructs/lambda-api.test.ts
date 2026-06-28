import * as cdk from "aws-cdk-lib/core";
import { Match, Template } from "aws-cdk-lib/assertions";
import { LambdaApi } from "../../lib/constructs/lambda-api";

describe("LambdaApi", () => {
  let template: Template;

  beforeAll(() => {
    const app = new cdk.App();
    const stack = new cdk.Stack(app, "TestStack");
    new LambdaApi(stack, "LambdaApi", {
      envName: "test",
      lambdaMemorySize: 128,
      logRetentionDays: 7,
      artifactsBucketName: "test-bucket",
    });
    template = Template.fromStack(stack);
  });

  describe("API Lambda", () => {
    test("provided.al2023 / arm64 / bootstrap を使用する", () => {
      template.hasResourceProperties("AWS::Lambda::Function", {
        FunctionName: "test-pollen-tracker-api",
        Runtime: "provided.al2023",
        Architectures: ["arm64"],
        Handler: "bootstrap",
      });
    });

    test("APP_ENV / APP_PORT / APP_PROJECT 環境変数が設定されている", () => {
      template.hasResourceProperties("AWS::Lambda::Function", {
        FunctionName: "test-pollen-tracker-api",
        Environment: {
          Variables: Match.objectLike({
            APP_ENV: "test",
            APP_PORT: "8080",
            APP_PROJECT: "pollen-tracker",
          }),
        },
      });
    });
  });

  describe("Authorizer Lambda", () => {
    test("provided.al2023 / arm64 / bootstrap を使用する", () => {
      template.hasResourceProperties("AWS::Lambda::Function", {
        FunctionName: "test-pollen-tracker-authorizer",
        Runtime: "provided.al2023",
        Architectures: ["arm64"],
        Handler: "bootstrap",
      });
    });
  });

  describe("HTTP API", () => {
    test("HTTP API が生成される", () => {
      template.hasResourceProperties("AWS::ApiGatewayV2::Api", {
        Name: "test-pollen-tracker",
        ProtocolType: "HTTP",
      });
    });

    test("Lambda Authorizer の identitySource が x-api-key ヘッダーである", () => {
      template.hasResourceProperties("AWS::ApiGatewayV2::Authorizer", {
        AuthorizerType: "REQUEST",
        IdentitySource: ["$request.header.x-api-key"],
        EnableSimpleResponses: true,
      });
    });
  });
});
