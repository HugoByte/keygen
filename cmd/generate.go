/*
Copyright © 2021 HugoByte <hello@hugobyte.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/hugobyte/keygen/keystore/near"
	"github.com/spf13/cobra"
)

var password string

var out string

var privateKey string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates KeyStore form private key ",
	Long: `
	This Commmand Generates Keystore from the newly generated Publickey and Private Key pair if PrivateKey falg is not set. 
	If Private Key is provided Keystore will be generated by using it`,

	Run: func(cmd *cobra.Command, args []string) {

		if privateKey != "" {
			err := GenerateNewKeystoreFromSeed(out, privateKey, password)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := GenerateNewKeystore(out, password)
			if err != nil {
				fmt.Println(err)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&password, "secret", "s", "", "Password to Create KeyStore")
	generateCmd.Flags().StringVarP(&out, "out", "o", "keystore.json", "OutPut file path")
	generateCmd.MarkFlagRequired("pass")
	generateCmd.Flags().StringVarP(&privateKey, "privatekey", "p", "", "Private Key")

}

func GenerateNewKeystore(file string, pw string) error {

	///Generate New KeyPair

	keypair, err := near.NewKeyPair()
	if err != nil {
		return err
	}

	/// Genreate KeyStore from the Private Key obtained from Keypair

	err = near.EncryptKey(keypair, pw, file)
	if err != nil {
		return err
	}

	return nil
}

func GenerateNewKeystoreFromSeed(file string, privateKey string, password string) error {

	keypair, err := near.NewKeyPairFromPrivateKey(privateKey)

	if err != nil {
		return err
	}

	err = near.EncryptKey(keypair, password, file)
	if err != nil {
		return err
	}

	return nil
}
