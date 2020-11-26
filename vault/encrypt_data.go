package vault

import (
	"encoding/base64"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func EncryptData(client *api.Client, mount, key, data string) (string, error) {
	logical := client.Logical()

	b64 := base64.StdEncoding.EncodeToString([]byte(data))

	secret, err := logical.Write(fmt.Sprintf("%s/encrypt/%s", mount, key), map[string]interface{}{
		"plaintext": b64,
	})
	if err != nil {
		return "", err
	}

	encryptedData := secret.Data["ciphertext"]
	return encryptedData.(string), nil
}
