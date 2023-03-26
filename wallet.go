package main

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

const dir = "./keys"

func LoadKey(acc accounts.Account, passphrase string, dirs ...string) *keystore.KeyStore {
	dir := dir
	if len(dirs) > 0 {
		dir = dirs[0]
	}
	ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	Ensure(ks.Unlock(acc, passphrase))
	return ks
}

func GenerateKey(passphrase string, dirs ...string) (accounts.Account, error) {
	dir := dir
	if len(dirs) > 0 {
		dir = dirs[0]
	}

	return keystore.StoreKey(dir, passphrase, keystore.StandardScryptN, keystore.StandardScryptP)
}
