package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider. Vaultwarden assigns identifiers (UUIDs) server-side, and the
// Terraform provider keeps them, including the composite ones such as
// `<organization_id>/<collection_id>/<group_id>`, in `id`, so every resource
// uses IdentifierFromProvider. This map also acts as the include list.
var ExternalNameConfigs = map[string]config.ExternalName{
	"vaultwarden_account_register":              config.IdentifierFromProvider,
	"vaultwarden_folder":                        config.IdentifierFromProvider,
	"vaultwarden_item_attachment":               config.IdentifierFromProvider,
	"vaultwarden_item_login":                    config.IdentifierFromProvider,
	"vaultwarden_item_secure_note":              config.IdentifierFromProvider,
	"vaultwarden_organization":                  config.IdentifierFromProvider,
	"vaultwarden_organization_api_key_rotation": config.IdentifierFromProvider,
	"vaultwarden_organization_collection":       config.IdentifierFromProvider,
	"vaultwarden_organization_collection_group": config.IdentifierFromProvider,
	"vaultwarden_organization_collection_user":  config.IdentifierFromProvider,
	"vaultwarden_organization_group":            config.IdentifierFromProvider,
	"vaultwarden_organization_group_member":     config.IdentifierFromProvider,
	"vaultwarden_organization_policy":           config.IdentifierFromProvider,
	"vaultwarden_organization_user":             config.IdentifierFromProvider,
	"vaultwarden_user":                          config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
