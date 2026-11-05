package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	tfprovider "github.com/STASiAN/crossplane-provider-vaultwarden/third_party/terraform-provider-vaultwarden"
)

const (
	resourcePrefix = "vaultwarden"
	modulePath     = "github.com/STASiAN/crossplane-provider-vaultwarden"
)

// tfProviderVersion is only reported to the embedded Terraform provider.
const tfProviderVersion = "crossplane"

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	return newProvider("vaultwarden.crossplane.io")
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("vaultwarden.m.crossplane.io",
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))
}

func newProvider(rootGroup string, opts ...ujconfig.ProviderOption) *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		append([]ujconfig.ProviderOption{
			ujconfig.WithRootGroup(rootGroup),
			// Every resource goes through the embedded Plugin Framework provider.
			ujconfig.WithIncludeList([]string{}),
			ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
			ujconfig.WithTerraformPluginFrameworkProvider(tfprovider.New(tfProviderVersion)()),
			ujconfig.WithFeaturesPackage("internal/features"),
			ujconfig.WithDefaultResourceOptions(
				ExternalNameConfigurations(),
				GroupKindOverrides(),
				References(),
			),
		}, opts...)...)

	pc.ConfigureResources()
	return pc
}
