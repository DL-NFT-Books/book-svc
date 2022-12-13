package handlers

import (
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/responses"
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewListBooksRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	books, err := helpers.GetBookListByRequest(r, &request)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get books")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if books == nil {
		ape.Render(w, resources.BookListResponse{
			Data: make([]resources.Book, 0),
		})
		return
	}

	data, err := helpers.NewBooksList(books)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to form up book list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	links := responses.CreateLinks(
		r.URL, request.OffsetPageParams)

	count, err := helpers.GetBooksCount(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to form up book list response")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if *count <= (request.OffsetPageParams.PageNumber+1)*request.OffsetPageParams.Limit {
		links.Next = ""
	}
	ape.Render(w, resources.BookListResponse{
		Data:  data,
		Links: links,
	})
}
