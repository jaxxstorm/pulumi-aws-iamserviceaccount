// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Iamserviceaccount
{
    [IamserviceaccountResourceType("iam:index:IamServiceAccount")]
    public partial class IamServiceAccount : Pulumi.ComponentResource
    {
        /// <summary>
        /// The bucket resource.
        /// </summary>
        [Output("bucket")]
        public Output<Pulumi.Aws.S3.Bucket> Bucket { get; private set; } = null!;

        /// <summary>
        /// The website URL.
        /// </summary>
        [Output("websiteUrl")]
        public Output<string> WebsiteUrl { get; private set; } = null!;


        /// <summary>
        /// Create a IamServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamServiceAccount(string name, IamServiceAccountArgs args, ComponentResourceOptions? options = null)
            : base("iam:index:IamServiceAccount", name, args ?? new IamServiceAccountArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class IamServiceAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HTML content for index.html.
        /// </summary>
        [Input("indexContent", required: true)]
        public Input<string> IndexContent { get; set; } = null!;

        public IamServiceAccountArgs()
        {
        }
    }
}
