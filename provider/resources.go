// Copyright 2016-2018, Pulumi Corporation.
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

package megaport

import (
	"fmt"
	"github.com/BeStateless/pulumi-megaport/provider/pkg/version"
	"path/filepath"

	megaported "github.com/megaport/terraform-provider-megaport/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "megaport"
	// modules:
	mainMod     = "index" // the megaport module
	aws         = "aws"
	gcp         = "gcp"
	azure       = "azure"
	cloudRouter = "cloudRouter"
	port        = "port"
	vxc         = "vxc"
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(megaported.Provider())

	const gitHubOrg = "BeStateless"
	const pulumiVersion = "3.53.1"

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "megaport",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Stateless",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing megaport cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "megaport", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: fmt.Sprintf("https://github.com/%s/pulumi-megaport", gitHubOrg),
		// The GitHub Org for the provider - defaults to `terraform-providers`
		GitHubOrg: gitHubOrg,
		Config:    map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: tfbridge.MakeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"megaport_aws_connection":   {Tok: tfbridge.MakeResource(mainPkg, aws, "AwsConnection")},
			"megaport_gcp_connection":   {Tok: tfbridge.MakeResource(mainPkg, gcp, "GcpConnection")},
			"megaport_azure_connection": {Tok: tfbridge.MakeResource(mainPkg, azure, "AzureConnection")},
			"megaport_mcr":              {Tok: tfbridge.MakeResource(mainPkg, cloudRouter, "CloudRouter")},
			"megaport_port":             {Tok: tfbridge.MakeResource(mainPkg, port, "Port")},
			"megaport_vxc":              {Tok: tfbridge.MakeResource(mainPkg, vxc, "VxcConnection")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"megaport_aws_connection":   {Tok: tfbridge.MakeDataSource(mainPkg, aws, "getAwsConnection")},
			"megaport_azure_connection": {Tok: tfbridge.MakeDataSource(mainPkg, azure, "getAzureConnection")},
			"megaport_gcp_connection":   {Tok: tfbridge.MakeDataSource(mainPkg, gcp, "getGcpConnection")},
			"megaport_location":         {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLocation")},
			"megaport_locations":        {Tok: tfbridge.MakeDataSource(mainPkg, mainMod, "getLocations")},
			"megaport_mcr":              {Tok: tfbridge.MakeDataSource(mainPkg, cloudRouter, "getCloudRouter")},
			"megaport_partner_port":     {Tok: tfbridge.MakeDataSource(mainPkg, port, "getPartnerPort")},
			"megaport_port":             {Tok: tfbridge.MakeDataSource(mainPkg, port, "getPort")},
			"megaport_ports":            {Tok: tfbridge.MakeDataSource(mainPkg, port, "getPorts")},
			"megaport_vxc":              {Tok: tfbridge.MakeDataSource(mainPkg, vxc, "getVxcConnection")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": pulumiVersion,
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/%s/pulumi-%[1]s/sdk/", gitHubOrg, mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": pulumiVersion,
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
