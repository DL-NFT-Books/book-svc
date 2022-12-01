package models

import (
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

type (
	// CreateBookParams is a helper struct to be included when calling CreateBook request
	CreateBookParams struct {
		Banner      resources.Media
		ChainId     int32
		Description string
		File        resources.Media
		Price       string
		Title       string
		TokenName   string
		TokenSymbol string
	}

	// UpdateBookParams is a helper struct to be included when calling UpdateBook request
	UpdateBookParams struct {
		Id          int64
		Banner      *resources.Media `json:"banner,omitempty"`
		Description *string          `json:"description,omitempty"`
		File        *resources.Media `json:"file,omitempty"`
		Title       *string          `json:"title,omitempty"`
	}

	// ListBooksParams is a helper struct to be included when calling ListBooks request
	ListBooksParams requests.GetBooksRequest
)
