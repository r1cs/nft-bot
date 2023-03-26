package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/r1cs/nft-bot/binding"
	"github.com/r1cs/nft-bot/tools"
	"math/big"
)

type A struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}
type Info struct {
	Attributes []A `json:"attributes"`
}

var m map[uint64]bool = func() map[uint64]bool {

	m := make(map[uint64]bool)
	point := []uint64{1615, 1643, 1755, 1783, 1895, 1923, 2035, 2063, 2175, 2203, 2315, 1343, 2455, 2483}
	for _, p := range point {
		m[p-1] = true
	}
	return m
}()

func nextIsRare(tokenId *big.Int) bool {
	return m[tokenId.Uint64()]
}

func main() {

}

func runBot(client *ethclient.Client, seaDrop *binding.ISeaDrop, erc721A *binding.IERC721A) {

}

func Ensure(err error) {
	if err != nil {
		panic(any(err))
	}
}

func runNFT() {
	limit := 1
	client, err := tools.NewEthClientWithProxy("https://bsc-dataseed1.defibit.io/", "http://localhost:7890")

	start, err := client.BlockNumber(context.Background())
	Ensure(err)
	fmt.Println(start)

	//airdrop, err := binding.NewISeaDrop(common.HexToAddress("0x00005ea00ac477b1030ce78506496e8c2de24bf5"), client)
	//utils.Ensure(err)

	pmoDrop, err := binding.NewIERC721A(common.HexToAddress("0xd8DE41A2AE0E1D0e863E494F45B869789963b35d"), client)
	Ensure(err)
	totalSupply, err := pmoDrop.TotalSupply(nil)
	Ensure(err)
	fmt.Println(totalSupply)
	esIterator, err := pmoDrop.FilterTransfer(&bind.FilterOpts{Start: start}, []common.Address{common.Address{}}, nil, nil)
	Ensure(err)
	for esIterator.Next() {
		e := esIterator.Event
		log.Info("pmo minted", "id", e.TokenId, "from", e.From, "to", e.To)
		if nextIsRare(e.TokenId) {
			//try to buy, if did not reach limit
			if limit > 0 {
				//airdrop.MintPublic()
				//limit--

			}
		}
	}
}
