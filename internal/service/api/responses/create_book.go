package responses

import "gitlab.com/tokend/nft-books/book-svc/resources"

func NewSignCreateResponse(bookId, tokenId int64, signature resources.SignatureAttributes) resources.CreateSignatureResponse {
	createSignatureKey := resources.NewKeyInt64(1, resources.CREATE_SIGNATURES)
	signatureKey := resources.NewKeyInt64(1, resources.SIGNATURES)

	included := resources.Included{}
	included.Add(&resources.Signature{
		Key:        signatureKey,
		Attributes: signature,
	})

	return resources.CreateSignatureResponse{
		Data: resources.CreateSignature{
			Key: createSignatureKey,
			Attributes: resources.CreateSignatureAttributes{
				TokenId: int32(tokenId),
				BookId:  int32(bookId),
			},
			Relationships: resources.CreateSignatureRelationships{
				Signature: resources.Relation{
					Data: &signatureKey,
				},
			},
		},
		Included: included,
	}
}
