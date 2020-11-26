package vault

import (
	"errors"
	"fmt"

	"github.com/hashicorp/vault/api"
)

func CreateKey(client *api.Client, mount, key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	logical := client.Logical()

	_, err := logical.Write(fmt.Sprintf("%s/keys/%s", mount, key), nil)
	if err != nil {
		return err
	}

	return nil
}
