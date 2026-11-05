package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func TestProviderConfiguresEveryResource(t *testing.T) {
	for name, pc := range map[string]func() map[string]string{
		"cluster": func() map[string]string {
			return kinds(GetProvider().Resources)
		},
		"namespaced": func() map[string]string {
			return kinds(GetProviderNamespaced().Resources)
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := pc()
			if len(got) != len(ExternalNameConfigs) {
				t.Fatalf("got %d resources, want %d", len(got), len(ExternalNameConfigs))
			}
			for tf, kind := range kindMap {
				if got[tf] != kind {
					t.Errorf("%s: got kind %q, want %q", tf, got[tf], kind)
				}
			}
		})
	}
}

func TestReferences(t *testing.T) {
	resources := GetProviderNamespaced().Resources
	cases := map[string]map[string]string{
		"vaultwarden_organization_collection": {
			"organization_id": "vaultwarden_organization",
		},
		"vaultwarden_organization_collection_group": {
			"organization_id": "vaultwarden_organization",
			"collection_id":   "vaultwarden_organization_collection",
			"group_id":        "vaultwarden_organization_group",
		},
		"vaultwarden_organization_group_member": {
			"organization_id":      "vaultwarden_organization",
			"group_id":             "vaultwarden_organization_group",
			"organization_user_id": "vaultwarden_organization_user",
		},
		"vaultwarden_item_login": {
			"organization_id": "vaultwarden_organization",
			"collection_ids":  "vaultwarden_organization_collection",
			"folder_id":       "vaultwarden_folder",
		},
		"vaultwarden_item_attachment": {
			"item_id": "vaultwarden_item_login",
		},
	}
	for res, want := range cases {
		r, ok := resources[res]
		if !ok {
			t.Fatalf("resource %s is not configured", res)
		}
		for arg, target := range want {
			if got := r.References[arg].TerraformName; got != target {
				t.Errorf("%s.%s: got reference to %q, want %q", res, arg, got, target)
			}
		}
	}
	if _, ok := resources["vaultwarden_organization"].References["organization_id"]; ok {
		t.Error("vaultwarden_organization must not reference itself")
	}
}

func kinds(resources map[string]*ujconfig.Resource) map[string]string {
	m := make(map[string]string, len(resources))
	for name, r := range resources {
		m[name] = r.Kind
	}
	return m
}
