import * as cdk from "aws-cdk-lib/core";
import { Template } from "aws-cdk-lib/assertions";
import { PollenStack } from "../lib/stacks/pollen-stack";
import { devConfig } from "../config/env-config";

test("PollenStack synthesizes without error", () => {
  const app = new cdk.App();
  const stack = new PollenStack(app, "PollenStack-Test", { config: devConfig });
  const template = Template.fromStack(stack);
  expect(template).toBeDefined();
});
