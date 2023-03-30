package models

import (
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
)

type (

	// UpdateBookParams is a helper struct to be included when calling UpdateBook request
	UpdateBookParams struct {
		Id              int64                  `json:"id"`
		Banner          *resources.Media       `json:"banner,omitempty"`
		Description     *string                `json:"description,omitempty"`
		File            *resources.Media       `json:"file,omitempty"`
		ChainId         int64                  `json:"chain_id"`
		ContractAddress string                 `json:"contract_address"`
		DeployStatus    resources.DeployStatus `json:"deploy_status"`
	}
	// ListBooksParams is a helper struct to be included when calling ListBooks request
	ListBooksParams requests.ListBooksRequest
)
