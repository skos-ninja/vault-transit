package vault

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func DecryptData(client *api.Client, mount, key, encryptedData string) (string, error) {
	logical := client.Logical()

	out, err := logical.Write(fmt.Sprintf("%s/decrypt/%s", mount, key), map[string]interface{}{
		"ciphertext": encryptedData,
	})
	if err != nil {
		return "", err
	}

	b64 := out.Data["plaintext"].(string)

	decoded, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
