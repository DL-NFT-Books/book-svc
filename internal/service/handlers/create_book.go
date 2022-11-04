package handlers

import (
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const defaultContractVersion = "1"

func CreateBook(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	err = helpers.CheckMediaTypes(r, req.Data.Attributes.Banner.Attributes.MimeType, req.Data.Attributes.File.Attributes.MimeType)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&req.Data.Attributes.Banner, &req.Data.Attributes.File)
	if media == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// TODO: get price && name from token contract (?)

	book := data.Book{
		Title:       req.Data.Attributes.Title,
		Description: req.Data.Attributes.Description,
		CreatedAt:   time.Now(),
		// mocked
		Price:           "100",
		ContractAddress: req.Data.Attributes.ContractAddress,
		// mocked
		ContractName: "contract name",
		// mocked
		ContractSymbol:  "contract symbol",
		ContractVersion: defaultContractVersion,
		Banner:          media[0],
		File:            media[1],
	}

	bookID, err := helpers.BooksQ(r).Insert(book)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	book.ID = bookID

	data, err := helpers.NewBook(&book)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, resources.BookResponse{
		Data: *data,
	})
}
