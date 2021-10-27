// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"fmt"
	"encoding/json"
)

// The set of arguments for creating a StaticPage component resource.
type IamServiceAccountArgs struct {
	// The HTML content for index.html.
	NamespaceName pulumi.StringInput `pulumi:"namespaceName"`
	OidcIssuer pulumi.StringInput `pulumi:"oidcIssuer"`
	OidcProvider pulumi.StringInput `pulumi:"oidcProvider"`
	ServiceAccountName pulumi.StringInput `pulumi:"serviceAccountName`

}

// The StaticPage component resource.
type IamServiceAccount struct {
	pulumi.ResourceState
}

// NewStaticPage creates a new StaticPage component resource.
func NewIamServiceAccount(ctx *pulumi.Context,
	name string, args *IamServiceAccountArgs, opts ...pulumi.ResourceOption) (*IamServiceAccount, error) {
	if args == nil {
		args = &IamServiceAccountArgs{}
	}

	component := &IamServiceAccount{}
	err := ctx.RegisterComponentResource("iam:index:IamServiceAccount", name, component, opts...)
	if err != nil {
		return nil, err
	}


	assumeRolePolicyJSON := pulumi.All(args.NamespaceName, args.OidcIssuer, args.OidcProvider, args.ServiceAccountName).ApplyT(
		func(args []interface{}) (string, error) {
			ns := args[0].(string)
			issuer := args[1].(string)
			provider := args[2].(string)
			serviceAccountName := args[3].(string)
			policyJSON, err := json.Marshal(map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []interface{}{
					map[string]interface{}{
						"Effect": "Allow",
						"Principal": map[string]interface{}{
							"Federated": provider,
						},
						"Action": "sts:AssumeRoleWithWebIdentity",
						"Condition": map[string]interface{}{
							"StringEquals": map[string]interface{}{
								fmt.Sprintf("%s:sub", issuer): fmt.Sprintf("system:serviceaccount:%s:%s", ns, serviceAccountName),
							},
						},
					},
				},
			})
			if err != nil {
				return "", err
			}
			return string(policyJSON), nil
		},
	).(pulumi.StringOutput)

	iamRole, err := iam.NewRole(ctx, fmt.Sprintf("%s-role", name), &iam.RoleArgs{
		AssumeRolePolicy: assumeRolePolicyJSON,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, fmt.Errorf("error creating IAM role: %v", err)
	}

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"iamRole":     iamRole,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
