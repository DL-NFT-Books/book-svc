package models

import (
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type (
	// CreateBookParams is a helper struct to be included when calling CreateBook request
	CreateBookParams struct {
		Banner      resources.Media
		Description string
		File        resources.Media
		Price       string
		FloorPrice  string
		Title       string
		TokenName   string
		TokenSymbol string
	}

	// UpdateBookParams is a helper struct to be included when calling UpdateBook request
	UpdateBookParams struct {
		Id                 int64
		Banner             *resources.Media        `json:"banner,omitempty"`
		Description        *string                 `json:"description,omitempty"`
		File               *resources.Media        `json:"file,omitempty"`
		Title              *string                 `json:"title,omitempty"`
		ContractName       *string                 `json:"contract_name,omitempty"`
		ContractAddress    *string                 `json:"contract_address,omitempty"`
		DeployStatus       *resources.DeployStatus `json:"deploy_status,omitempty"`
		Symbol             *string                 `json:"symbol,omitempty"`
		Price              *string                 `json:"price,omitempty"`
		FloorPrice         *string                 `json:"floor_price,omitempty"`
		VoucherToken       *string                 `json:"voucher_token,omitempty"`
		VoucherTokenAmount *string                 `json:"voucher_token_amount,omitempty"`
	}

	// ListBooksParams is a helper struct to be included when calling ListBooks request
	ListBooksParams requests.ListBooksRequest
)
