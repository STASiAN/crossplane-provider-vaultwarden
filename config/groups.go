package config

import "github.com/crossplane/upjet/v2/pkg/config"

// kindMap places every resource in the root API group
// (vaultwarden.crossplane.io / vaultwarden.m.crossplane.io) with a fully
// descriptive kind, instead of upjet's default word-splitting which would
// produce groups like "organization" with kind "Collection" for
// vaultwarden_organization_collection.
var kindMap = map[string]string{
	"vaultwarden_account_register":              "AccountRegister",
	"vaultwarden_folder":                        "Folder",
	"vaultwarden_item_attachment":               "ItemAttachment",
	"vaultwarden_item_login":                    "ItemLogin",
	"vaultwarden_item_secure_note":              "ItemSecureNote",
	"vaultwarden_organization":                  "Organization",
	"vaultwarden_organization_api_key_rotation": "OrganizationAPIKeyRotation",
	"vaultwarden_organization_collection":       "OrganizationCollection",
	"vaultwarden_organization_collection_group": "OrganizationCollectionGroup",
	"vaultwarden_organization_collection_user":  "OrganizationCollectionUser",
	"vaultwarden_organization_group":            "OrganizationGroup",
	"vaultwarden_organization_group_member":     "OrganizationGroupMember",
	"vaultwarden_organization_policy":           "OrganizationPolicy",
	"vaultwarden_organization_user":             "OrganizationUser",
	"vaultwarden_user":                          "User",
}

// GroupKindOverrides moves all resources into the root group with explicit
// kinds.
func GroupKindOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		if k, ok := kindMap[r.Name]; ok {
			r.ShortGroup = ""
			r.Kind = k
		}
	}
}
