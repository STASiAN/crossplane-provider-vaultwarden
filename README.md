# crossplane-provider-vaultwarden

`crossplane-provider-vaultwarden` is a [Crossplane](https://crossplane.io/) provider
for [Vaultwarden](https://github.com/dani-garcia/vaultwarden), the
Bitwarden-compatible server. It is built with
[Upjet](https://github.com/crossplane/upjet) code generation from the
[ottramst/vaultwarden](https://registry.terraform.io/providers/ottramst/vaultwarden/latest)
Terraform provider (v0.8.0) and exposes XRM-conformant managed resources for
the Vaultwarden API.

## Installation

Install the provider with Crossplane v2.x:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: crossplane-provider-vaultwarden
spec:
  package: ghcr.io/stasian/crossplane-provider-vaultwarden:v0.1.0
```

## Configuration

Create a credentials secret and a `ProviderConfig`. The credentials JSON keys
mirror the Terraform provider attributes: `endpoint` (required),
`email`, `master_password`, `client_id`, `client_secret`, `admin_token`,
`max_retry_wait`.

The provider authenticates in up to two ways, which can be combined:

- **As a user**, with `email` and `master_password`. Required for resources
  that act on behalf of a user (folders, items, organizations, collections,
  groups, ...), because their data is encrypted client-side with the user's
  keys. If the user has two-step login enabled, also set the user's personal
  API key (`client_id` and `client_secret`).
- **As the server admin**, with `admin_token`. Required for `User` and
  `AccountRegister`, which use the `/admin` API. The admin page must be enabled
  on the server.

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: vaultwarden-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "endpoint": "https://vault.example.com",
      "email": "admin@example.com",
      "master_password": "<master password>",
      "admin_token": "<admin token>"
    }
---
apiVersion: vaultwarden.m.crossplane.io/v1beta1
kind: ClusterProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: vaultwarden-creds
      namespace: crossplane-system
      key: credentials
```

Then create managed resources, e.g.:

```yaml
apiVersion: vaultwarden.m.crossplane.io/v1alpha1
kind: Organization
metadata:
  name: acme
  namespace: default
spec:
  forProvider:
    name: ACME
  providerConfigRef:
    kind: ClusterProviderConfig
    name: default
---
apiVersion: vaultwarden.m.crossplane.io/v1alpha1
kind: OrganizationCollection
metadata:
  name: infra
  namespace: default
spec:
  forProvider:
    name: Infrastructure
    organizationIdRef:
      name: acme
  providerConfigRef:
    kind: ClusterProviderConfig
    name: default
```

Sensitive arguments, such as an item's password or an attachment's content,
are read from Kubernetes secrets (`passwordSecretRef`, `contentSecretRef`, ...).
ID arguments can be set directly or resolved from other managed resources with
`*Ref` / `*Selector` fields. See [`examples/`](examples/) and
[`examples-generated/`](examples-generated/) for more.

## Resource coverage

All kinds live in the `vaultwarden.crossplane.io` group (cluster-scoped, legacy)
and `vaultwarden.m.crossplane.io` (namespaced, Crossplane v2 style):

| Kind | Terraform resource |
| --- | --- |
| `AccountRegister` | `vaultwarden_account_register` |
| `Folder` | `vaultwarden_folder` |
| `ItemAttachment` | `vaultwarden_item_attachment` |
| `ItemLogin` | `vaultwarden_item_login` |
| `ItemSecureNote` | `vaultwarden_item_secure_note` |
| `Organization` | `vaultwarden_organization` |
| `OrganizationAPIKeyRotation` | `vaultwarden_organization_api_key_rotation` |
| `OrganizationCollection` | `vaultwarden_organization_collection` |
| `OrganizationCollectionGroup` | `vaultwarden_organization_collection_group` |
| `OrganizationCollectionUser` | `vaultwarden_organization_collection_user` |
| `OrganizationGroup` | `vaultwarden_organization_group` |
| `OrganizationGroupMember` | `vaultwarden_organization_group_member` |
| `OrganizationPolicy` | `vaultwarden_organization_policy` |
| `OrganizationUser` | `vaultwarden_organization_user` |
| `User` | `vaultwarden_user` |

`ItemAttachment` reads its content from `contentSecretRef` or
`contentBase64SecretRef`; its `file` argument is a path on the provider's
filesystem and is of little use in a cluster.

## Development

```bash
make submodules   # initialize the build submodule
make generate     # regenerate APIs/controllers from the Terraform schema
make build        # build binaries and packages
make run          # run the controller out-of-cluster against $KUBECONFIG
```

The Terraform provider runs in-process (upjet no-fork mode with the Terraform
Plugin Framework). Upstream keeps its code under `internal/`, which Go does not
allow importing from another module, so `make generate` copies the sources of
`TERRAFORM_PROVIDER_VERSION` into
[`third_party/terraform-provider-vaultwarden`](third_party/terraform-provider-vaultwarden)
(MPL-2.0, see its `LICENSE`) with
[`hack/sync-tf-provider.sh`](hack/sync-tf-provider.sh). To upgrade the
Terraform provider, bump `TERRAFORM_PROVIDER_VERSION` in the `Makefile` and run
`make generate`.

## Releasing

Push a `v*` tag; the [`publish`](.github/workflows/publish.yml) workflow builds
the multi-arch package, pushes it to
`ghcr.io/stasian/crossplane-provider-vaultwarden` (and to
`xpkg.upbound.io/stasian` when the `XPKG_ACCESS_ID` / `XPKG_TOKEN` secrets are
set) and creates the GitHub release.

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/STASiAN/crossplane-provider-vaultwarden/issues).

## Licensing

crossplane-provider-vaultwarden is under the Apache 2.0 license. The vendored
Terraform provider in `third_party/` is under the Mozilla Public License 2.0.
