package main

import (
	"fmt"

	"transit/vault"

	"github.com/spf13/cobra"
)

var (
	cmd = &cobra.Command{
		Use:  "transit",
		RunE: runE,
		Args: cobra.ExactArgs(1),
	}

	mountPath = "transit"
	key       = ""
)

func init() {
	cmd.Flags().StringVar(&mountPath, "mount", mountPath, "")
	cmd.Flags().StringVar(&key, "key", key, "")
}

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func runE(cmd *cobra.Command, args []string) error {
	data := args[0]

	client, err := vault.Connect()
	if err != nil {
		return err
	}

	err = vault.CreateMount(client, mountPath)
	if err != nil {
		return err
	}

	err = vault.CreateKey(client, mountPath, key)
	if err != nil {
		return err
	}

	encryptedData, err := vault.EncryptData(client, mountPath, key, data)
	if err != nil {
		return err
	}
	fmt.Println(encryptedData)

	decryptedData, err := vault.DecryptData(client, mountPath, key, encryptedData)
	if err != nil {
		return err
	}
	if data != decryptedData {
		return fmt.Errorf("data doesn't match: %s", decryptedData)
	}

	fmt.Println("Data verified")
	return nil
}
