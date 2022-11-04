package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/requests"
)

func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewUpdateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	book, err := helpers.GetBookByID(r, req.ID)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if book == nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	banner := req.Data.Attributes.Banner
	file := req.Data.Attributes.File

	err = helpers.CheckMediaTypes(r, banner.Attributes.MimeType, file.Attributes.MimeType)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err = helpers.SetMediaLinks(r, &banner, &file); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&banner, &file)
	if media == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	bookToUpdate := data.Book{
		ID:          req.ID,
		Title:       req.Data.Attributes.Title,
		Description: req.Data.Attributes.Description,
		Banner:      media[0],
		File:        media[1],
	}

	err = helpers.BooksQ(r).Update(bookToUpdate)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, http.StatusNoContent)
}
