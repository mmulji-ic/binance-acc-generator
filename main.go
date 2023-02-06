package main

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bip39 "github.com/cosmos/go-bip39"
	"github.com/tyler-smith/go-bip32"
)

func main() {
	// Generate a account from `gaiad keys add cosmos-test --keyring-backend file`
	// - name: cosmos-test
	//  type: local

	mnemonic := ""
	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	// seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "cosmostest")
	// if err != nil {
	// 	panic(err)
	// }
	seed := bip39.NewSeed(mnemonic, "cosmostest")

	masterKey, _ := bip32.NewMasterKey(seed)
	pubkey := secp256k1.PubKey{
		Key: masterKey.PublicKey().Key,
	}

	addr := sdk.AccAddress(pubkey.Address())
	fmt.Println(addr.String())
}
