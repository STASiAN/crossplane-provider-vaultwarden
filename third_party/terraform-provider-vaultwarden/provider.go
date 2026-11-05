// Package tfvaultwarden exposes the vendored ottramst/vaultwarden Terraform
// provider, whose upstream code lives under internal/, to the rest of this
// module. Regenerate the vendored sources with hack/sync-tf-provider.sh.
package tfvaultwarden

import (
	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	vwprovider "github.com/STASiAN/crossplane-provider-vaultwarden/third_party/terraform-provider-vaultwarden/internal/provider"
)

// New returns a constructor of the Vaultwarden Terraform Plugin Framework
// provider.
func New(version string) func() fwprovider.Provider {
	return vwprovider.New(version)
}
