package config

import "github.com/crossplane/upjet/v2/pkg/config"

// referenceMap lists, per argument name, the resource whose external name
// (its Vaultwarden ID) the argument holds. Upjet generates a *Ref and a
// *Selector field for each of them.
var referenceMap = map[string]string{
	"organization_id":      "vaultwarden_organization",
	"collection_id":        "vaultwarden_organization_collection",
	"collection_ids":       "vaultwarden_organization_collection",
	"group_id":             "vaultwarden_organization_group",
	"organization_user_id": "vaultwarden_organization_user",
	"folder_id":            "vaultwarden_folder",
}

// References wires the ID arguments of every resource to the resources they
// identify.
func References() config.ResourceOption {
	return func(r *config.Resource) {
		for arg, target := range referenceMap {
			if _, ok := r.TerraformResource.Schema[arg]; !ok || r.Name == target {
				continue
			}
			r.References[arg] = config.Reference{TerraformName: target}
		}
		// An attachment can belong to any item kind, so the most common one is
		// the referenced kind; a secure note is set through itemId directly.
		if r.Name == "vaultwarden_item_attachment" {
			r.References["item_id"] = config.Reference{TerraformName: "vaultwarden_item_login"}
		}
	}
}
