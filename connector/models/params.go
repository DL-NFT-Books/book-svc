package models

import (
	"github.com/dl-nft-books/book-svc/internal/service/api/requests"
	"github.com/dl-nft-books/book-svc/resources"
)

type (

	// UpdateBookParams is a helper struct to be included when calling UpdateBook request
	UpdateBookParams resources.UpdateBook
	// ListBooksParams is a helper struct to be included when calling ListBooks request
	ListBooksParams requests.ListBooksRequest
)
