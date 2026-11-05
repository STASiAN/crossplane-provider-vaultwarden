package vaultwarden

import (
	"fmt"

	"github.com/STASiAN/crossplane-provider-vaultwarden/third_party/terraform-provider-vaultwarden/internal/vaultwarden/crypt"
	"github.com/STASiAN/crossplane-provider-vaultwarden/third_party/terraform-provider-vaultwarden/internal/vaultwarden/encryptedstring"
	"github.com/STASiAN/crossplane-provider-vaultwarden/third_party/terraform-provider-vaultwarden/internal/vaultwarden/symmetrickey"
)

// encryptString encrypts a value with key. An empty value is returned as an empty string,
// because Bitwarden clients store empty fields as null instead of encrypting them.
func encryptString(value string, key symmetrickey.Key) (string, error) {
	if value == "" {
		return "", nil
	}
	return crypt.EncryptAsString([]byte(value), key)
}

// decryptString decrypts an encrypted string with key. An empty value decrypts to an empty string.
func decryptString(value string, key symmetrickey.Key) (string, error) {
	if value == "" {
		return "", nil
	}
	encrypted, err := encryptedstring.NewFromEncryptedValue(value)
	if err != nil {
		return "", fmt.Errorf("failed to parse encrypted value: %w", err)
	}
	decrypted, err := crypt.Decrypt(encrypted, &key)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
