// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iamserviceaccount

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IamServiceAccount struct {
	pulumi.ResourceState

	// The bucket resource.
	Bucket s3.BucketOutput `pulumi:"bucket"`
	// The website URL.
	WebsiteUrl pulumi.StringOutput `pulumi:"websiteUrl"`
}

// NewIamServiceAccount registers a new resource with the given unique name, arguments, and options.
func NewIamServiceAccount(ctx *pulumi.Context,
	name string, args *IamServiceAccountArgs, opts ...pulumi.ResourceOption) (*IamServiceAccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IndexContent == nil {
		return nil, errors.New("invalid value for required argument 'IndexContent'")
	}
	var resource IamServiceAccount
	err := ctx.RegisterRemoteComponentResource("iam:index:IamServiceAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type iamServiceAccountArgs struct {
	// The HTML content for index.html.
	IndexContent string `pulumi:"indexContent"`
}

// The set of arguments for constructing a IamServiceAccount resource.
type IamServiceAccountArgs struct {
	// The HTML content for index.html.
	IndexContent pulumi.StringInput
}

func (IamServiceAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamServiceAccountArgs)(nil)).Elem()
}

type IamServiceAccountInput interface {
	pulumi.Input

	ToIamServiceAccountOutput() IamServiceAccountOutput
	ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput
}

func (*IamServiceAccount) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccount)(nil))
}

func (i *IamServiceAccount) ToIamServiceAccountOutput() IamServiceAccountOutput {
	return i.ToIamServiceAccountOutputWithContext(context.Background())
}

func (i *IamServiceAccount) ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountOutput)
}

func (i *IamServiceAccount) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return i.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (i *IamServiceAccount) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountPtrOutput)
}

type IamServiceAccountPtrInput interface {
	pulumi.Input

	ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput
	ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput
}

type iamServiceAccountPtrType IamServiceAccountArgs

func (*iamServiceAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccount)(nil))
}

func (i *iamServiceAccountPtrType) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return i.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (i *iamServiceAccountPtrType) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountPtrOutput)
}

// IamServiceAccountArrayInput is an input type that accepts IamServiceAccountArray and IamServiceAccountArrayOutput values.
// You can construct a concrete instance of `IamServiceAccountArrayInput` via:
//
//          IamServiceAccountArray{ IamServiceAccountArgs{...} }
type IamServiceAccountArrayInput interface {
	pulumi.Input

	ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput
	ToIamServiceAccountArrayOutputWithContext(context.Context) IamServiceAccountArrayOutput
}

type IamServiceAccountArray []IamServiceAccountInput

func (IamServiceAccountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*IamServiceAccount)(nil))
}

func (i IamServiceAccountArray) ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput {
	return i.ToIamServiceAccountArrayOutputWithContext(context.Background())
}

func (i IamServiceAccountArray) ToIamServiceAccountArrayOutputWithContext(ctx context.Context) IamServiceAccountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountArrayOutput)
}

// IamServiceAccountMapInput is an input type that accepts IamServiceAccountMap and IamServiceAccountMapOutput values.
// You can construct a concrete instance of `IamServiceAccountMapInput` via:
//
//          IamServiceAccountMap{ "key": IamServiceAccountArgs{...} }
type IamServiceAccountMapInput interface {
	pulumi.Input

	ToIamServiceAccountMapOutput() IamServiceAccountMapOutput
	ToIamServiceAccountMapOutputWithContext(context.Context) IamServiceAccountMapOutput
}

type IamServiceAccountMap map[string]IamServiceAccountInput

func (IamServiceAccountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*IamServiceAccount)(nil))
}

func (i IamServiceAccountMap) ToIamServiceAccountMapOutput() IamServiceAccountMapOutput {
	return i.ToIamServiceAccountMapOutputWithContext(context.Background())
}

func (i IamServiceAccountMap) ToIamServiceAccountMapOutputWithContext(ctx context.Context) IamServiceAccountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamServiceAccountMapOutput)
}

type IamServiceAccountOutput struct {
	*pulumi.OutputState
}

func (IamServiceAccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamServiceAccount)(nil))
}

func (o IamServiceAccountOutput) ToIamServiceAccountOutput() IamServiceAccountOutput {
	return o
}

func (o IamServiceAccountOutput) ToIamServiceAccountOutputWithContext(ctx context.Context) IamServiceAccountOutput {
	return o
}

func (o IamServiceAccountOutput) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return o.ToIamServiceAccountPtrOutputWithContext(context.Background())
}

func (o IamServiceAccountOutput) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return o.ApplyT(func(v IamServiceAccount) *IamServiceAccount {
		return &v
	}).(IamServiceAccountPtrOutput)
}

type IamServiceAccountPtrOutput struct {
	*pulumi.OutputState
}

func (IamServiceAccountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IamServiceAccount)(nil))
}

func (o IamServiceAccountPtrOutput) ToIamServiceAccountPtrOutput() IamServiceAccountPtrOutput {
	return o
}

func (o IamServiceAccountPtrOutput) ToIamServiceAccountPtrOutputWithContext(ctx context.Context) IamServiceAccountPtrOutput {
	return o
}

type IamServiceAccountArrayOutput struct{ *pulumi.OutputState }

func (IamServiceAccountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IamServiceAccount)(nil))
}

func (o IamServiceAccountArrayOutput) ToIamServiceAccountArrayOutput() IamServiceAccountArrayOutput {
	return o
}

func (o IamServiceAccountArrayOutput) ToIamServiceAccountArrayOutputWithContext(ctx context.Context) IamServiceAccountArrayOutput {
	return o
}

func (o IamServiceAccountArrayOutput) Index(i pulumi.IntInput) IamServiceAccountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IamServiceAccount {
		return vs[0].([]IamServiceAccount)[vs[1].(int)]
	}).(IamServiceAccountOutput)
}

type IamServiceAccountMapOutput struct{ *pulumi.OutputState }

func (IamServiceAccountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IamServiceAccount)(nil))
}

func (o IamServiceAccountMapOutput) ToIamServiceAccountMapOutput() IamServiceAccountMapOutput {
	return o
}

func (o IamServiceAccountMapOutput) ToIamServiceAccountMapOutputWithContext(ctx context.Context) IamServiceAccountMapOutput {
	return o
}

func (o IamServiceAccountMapOutput) MapIndex(k pulumi.StringInput) IamServiceAccountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IamServiceAccount {
		return vs[0].(map[string]IamServiceAccount)[vs[1].(string)]
	}).(IamServiceAccountOutput)
}

func init() {
	pulumi.RegisterOutputType(IamServiceAccountOutput{})
	pulumi.RegisterOutputType(IamServiceAccountPtrOutput{})
	pulumi.RegisterOutputType(IamServiceAccountArrayOutput{})
	pulumi.RegisterOutputType(IamServiceAccountMapOutput{})
}
