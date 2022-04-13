module github.com/BeStateless/pulumi-megaport/provider

go 1.18

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 upstream-v2.10.1
	github.com/megaport/terraform-provider-megaport v0.2.4 => ../../terraform-provider-megaport
    github.com/megaport/megaportgo v0.1.9-beta => ../../megaportgo
)

require (
	github.com/megaport/megaportgo v0.1.9-beta
	github.com/megaport/terraform-provider-megaport v0.2.4
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.28.0
)
