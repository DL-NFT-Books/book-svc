package connector

import (
	"github.com/dl-nft-books/doorman/resources"
)

func NewClaimsModel(address string, purpose string) resources.JwtClaims {
	model := resources.JwtClaims{
		Key: resources.Key{Type: resources.JWT_CLAIMS},
		Attributes: resources.JwtClaimsAttributes{
			EthAddress: address,
			Purpose: resources.Purpose{
				Type: purpose,
			},
		},
	}
	return model
}
