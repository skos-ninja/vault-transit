package vault

import (
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
)

const mountType = "transit"

func CreateMount(client *api.Client, mountPath string) error {
	// Ensure our mount path ends with a trailing slash
	if !strings.HasSuffix(mountPath, "/") {
		mountPath += "/"
	}

	sys := client.Sys()
	mounts, err := sys.ListMounts()
	if err != nil {
		return err
	}

	if v, ok := mounts[mountPath]; ok {
		if v.Type != mountType {
			return fmt.Errorf("mount of wrong type: %s", v.Type)
		}

		return nil
	}

	return sys.Mount(mountPath, &api.MountInput{
		Type: mountType,
	})
}
