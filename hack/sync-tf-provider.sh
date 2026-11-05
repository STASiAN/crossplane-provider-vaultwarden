#!/usr/bin/env bash
# Copies the Go sources of the ottramst/vaultwarden Terraform provider into
# third_party/terraform-provider-vaultwarden so that it can be embedded in the
# controller (upjet no-fork mode). The upstream provider keeps its code under
# internal/, which Go forbids importing from another module, hence the copy.
#
# Usage: hack/sync-tf-provider.sh <version without the leading v>
set -euo pipefail

VERSION="${1:?usage: $0 <terraform provider version>}"
REPO="${TERRAFORM_PROVIDER_REPO:-https://github.com/ottramst/terraform-provider-vaultwarden}"
ROOT="$(cd "$(dirname "$0")/.." && pwd)"
DEST="${ROOT}/third_party/terraform-provider-vaultwarden"
UPSTREAM_MODULE="github.com/ottramst/terraform-provider-vaultwarden"
LOCAL_MODULE="$(cd "${ROOT}" && go list -m)/third_party/terraform-provider-vaultwarden"

TMP="$(mktemp -d)"
trap 'rm -rf "${TMP}"' EXIT

git clone -q -c advice.detachedHead=false --depth 1 --branch "v${VERSION}" "${REPO}" "${TMP}/src"

rm -rf "${DEST}/internal"
mkdir -p "${DEST}"
cp "${TMP}/src/LICENSE" "${DEST}/LICENSE"
(cd "${TMP}/src" && find internal -name '*.go' ! -name '*_test.go' ! -path 'internal/vaultwarden/test/*') |
	while read -r f; do
		mkdir -p "${DEST}/$(dirname "${f}")"
		sed "s|\"${UPSTREAM_MODULE}/internal/|\"${LOCAL_MODULE}/internal/|g" "${TMP}/src/${f}" > "${DEST}/${f}"
	done
echo "${VERSION}" > "${DEST}/VERSION"
echo "synced ${UPSTREAM_MODULE} v${VERSION} into ${DEST}"
