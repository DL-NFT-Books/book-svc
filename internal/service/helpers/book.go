package helpers

import (
	"net/http"

	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

func GetBookByID(r *http.Request, id int64) (*data.Book, error) {
	return BooksQ(r).FilterActual().FilterByID(id).Get()
}

func GetBooksByPage(r *http.Request, page pgdb.OffsetPageParams) ([]data.Book, error) {
	return BooksQ(r).FilterActual().Page(page).Select()
}

func NewBooksList(books []data.Book) ([]resources.Book, resources.Included, error) {
	data := make([]resources.Book, len(books))
	included := resources.Included{}

	for i, book := range books {
		media, err := UnmarshalMedia(book.Banner, book.File)
		if err != nil {
			return nil, resources.Included{}, err
		}

		responseBook, err := NewBook(&book)
		if err != nil {
			return nil, resources.Included{}, err
		}

		data[i] = responseBook

		media[0].Key = resources.NewKeyInt64(book.ID, resources.BANNER)
		media[1].Key = resources.NewKeyInt64(book.ID, resources.FILE)
		included.Add(&media[0], &media[1])
	}

	return data, included, nil
}

func NewBook(book *data.Book) (resources.Book, error) {
	if book == nil {
		return resources.Book{}, nil
	}

	bannerKey := resources.NewKeyInt64(book.ID, resources.BANNER)
	documentKey := resources.NewKeyInt64(book.ID, resources.FILE)

	res := resources.Book{
		Key: resources.NewKeyInt64(book.ID, resources.BOOKS),
		Attributes: resources.BookAttributes{
			Title:           book.Title,
			Description:     book.Description,
			Price:           book.Price,
			ContractAddress: book.ContractAddress,
		},
		Relationships: resources.BookRelationships{
			Banner: resources.Relation{
				Data: &bannerKey,
			},
			File: resources.Relation{
				Data: &documentKey,
			},
		},
	}

	return res, nil
}
