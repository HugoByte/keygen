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
	"crypto/ed25519"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hugobyte/keygen/keystore/common"
	n "github.com/hugobyte/keygen/keystore/near"
	"github.com/spf13/cobra"
)

var (
	file string
	pass string
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypt command decrypts the keystore and returns the Private key",
	Long: `Decrypt command decrypts the keystore and returns the Private key 
	Prints  Private key will be of ed25519.PrivateKey to screen
	`,
	Run: func(cmd *cobra.Command, args []string) {

		private, err := DecryptKeyStore(file, pass)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		pKey, err := json.Marshal(private)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)

		}

		fmt.Println(string(pKey))

	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	decryptCmd.Flags().StringVarP(&pass, "pass", "p", "", "Password to Create KeyStore")
	decryptCmd.Flags().StringVarP(&file, "file", "f", "", "Keystore File or KeyStore File Path")
	decryptCmd.MarkFlagRequired("pass")
	decryptCmd.MarkFlagRequired("file")
}

func DecryptKeyStore(file string, password string) (ed25519.PrivateKey, error) {

	keystorejson, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	ksd, err := newKeyStoreData(keystorejson)
	if err != nil {
		return nil, err
	}
	privateKey, err := n.DecryptKey(ksd, password)
	if err != nil {
		return nil, err
	}

	return privateKey, nil

}

func newKeyStoreData(data []byte) (*common.KeyStoreData, error) {
	var ksData common.KeyStoreData
	if err := json.Unmarshal(data, &ksData); err != nil {
		return nil, err
	}

	return &ksData, nil
}