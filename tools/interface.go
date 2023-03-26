package tools

import "math/big"

// EnumTokenIdChecker only used for those nft that tokenId is increment by 1
type EnumTokenIdChecker interface {
	// WantNext check current tokenId and evaluate next token's value
	WantNext(tokenId *big.Int) bool
}
