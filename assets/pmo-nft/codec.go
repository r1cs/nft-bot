package pmo_nft

import (
	"encoding/json"
	"github.com/r1cs/nft-bot/assets"
	"math/big"
	"os"
)

type PMO1EnumTokenIdChecker map[uint64]bool

func (self PMO1EnumTokenIdChecker) WantNext(tokenId *big.Int) bool {
	return self[tokenId.Uint64()+1]
}

func GetPMO1LuckCheckerEnum() PMO1EnumTokenIdChecker {
	var codec []uint64
	data, err := os.ReadFile(assets.PMO1LuckFile)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &codec); err != nil {
		panic(err)
	}

	// inject to map to speed up
	ret := make(map[uint64]bool)
	for _, n := range codec {
		ret[n] = true
	}
	return ret
}
