import * as cdk from "aws-cdk-lib/core";
import { Template } from "aws-cdk-lib/assertions";
import { PollenStack } from "../lib/stacks/pollen-stack";
import { devConfig } from "../config/env-config";

describe("PollenStack", () => {
  let template: Template;

  beforeAll(() => {
    const app = new cdk.App();
    const stack = new PollenStack(app, "PollenStack-Test", {
      config: devConfig,
    });
    template = Template.fromStack(stack);
  });

  test("スタックが正常に合成される", () => {
    expect(template).toBeDefined();
  });

  test("3 つの Secrets Manager シークレットが生成される", () => {
    template.resourceCountIs("AWS::SecretsManager::Secret", 3);
  });

  test("Pollen API キー / Turso / x-api-key シークレットが生成される", () => {
    template.hasResourceProperties("AWS::SecretsManager::Secret", {
      Name: `${devConfig.envName}/pollen-tracker/pollen-api-key`,
    });
    template.hasResourceProperties("AWS::SecretsManager::Secret", {
      Name: `${devConfig.envName}/pollen-tracker/turso`,
    });
    template.hasResourceProperties("AWS::SecretsManager::Secret", {
      Name: `${devConfig.envName}/pollen-tracker/api-key`,
    });
  });

  test("2 つの Lambda 関数が生成される (API + Authorizer)", () => {
    template.resourceCountIs("AWS::Lambda::Function", 2);
  });

  test("HTTP API が 1 つ生成される", () => {
    template.resourceCountIs("AWS::ApiGatewayV2::Api", 1);
  });
});
