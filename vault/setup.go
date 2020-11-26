package vault

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/vault/api"
)

var (
	httpClient = &http.Client{
		Timeout: 10 * time.Second,
	}
)

func Connect() (*api.Client, error) {
	addr := os.Getenv(api.EnvVaultAddress)
	client, err := api.NewClient(&api.Config{Address: addr, HttpClient: httpClient})
	if err != nil {
		return nil, err
	}

	token := os.Getenv(api.EnvVaultToken)
	if token == "" {
		return nil, errors.New("missing token")
	}
	client.SetToken(token)

	return client, nil
}
