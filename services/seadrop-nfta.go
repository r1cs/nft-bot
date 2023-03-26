package services

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/r1cs/nft-bot/binding"
)

type SeaDropNftAService struct {
	c       *ethclient.Client
	seaDrop *binding.ISeaDrop
	erc721A *binding.IERC721A
}

//func runBot(client *ethclient.Client, seaDrop *binding.ISeaDrop, erc721A *binding.IERC721A) {
//	seaDrop.MintPublic()
//}
