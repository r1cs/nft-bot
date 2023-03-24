package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/laizy/web3/utils"
	"github.com/r1cs/nft-bot/binding"
	"math/big"
	"net/http"
	"net/url"
	"os"
)

type A struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}
type Info struct {
	Attributes []A `json:"attributes"`
}

func store() {
	data, err := os.ReadFile("ipfs_data.txt")
	utils.Ensure(err)

	luckNumber := make([]float64, 0)

	var info Info
	newData := make([]byte, 0)
	i := 0
	for _, b := range data {
		newData = append(newData, b)
		if b == '}' {
			i++
		}
		if i == 3 {
			i = 0
			if err := json.Unmarshal(newData, &info); err != nil {
				panic(err)
			}
			if info.Attributes[0].Value == "Archive Lv 1.0" {
				luckNumber = append(luckNumber, info.Attributes[1].Value.(float64))
			}
			newData = newData[:0]
		}
	}
	fmt.Println(luckNumber)
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
	store()
}

func runNFT() {
	limit := 1

	proxyURL, err := url.Parse("http://localhost:7890")
	if err != nil {
		panic(err)
	}

	// 创建 HTTP 客户端
	httpClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}}

	c, err := rpc.DialOptions(context.Background(), "https://bsc-dataseed1.defibit.io/", rpc.WithHTTPClient(httpClient))
	utils.Ensure(err)
	client := ethclient.NewClient(c)
	start, err := client.BlockNumber(context.Background())
	utils.Ensure(err)
	fmt.Println(start)

	//airdrop, err := binding.NewISeaDrop(common.HexToAddress("0x00005ea00ac477b1030ce78506496e8c2de24bf5"), client)
	//utils.Ensure(err)

	pmoDrop, err := binding.NewIERC721A(common.HexToAddress("0xd8DE41A2AE0E1D0e863E494F45B869789963b35d"), client)
	utils.Ensure(err)
	totalSupply, err := pmoDrop.TotalSupply(nil)
	utils.Ensure(err)
	fmt.Println(totalSupply)
	esIterator, err := pmoDrop.FilterTransfer(&bind.FilterOpts{Start: start}, []common.Address{common.Address{}}, nil, nil)
	utils.Ensure(err)
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
