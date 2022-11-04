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

func NewBooksList(books []data.Book) ([]resources.Book, error) {
	data := make([]resources.Book, len(books))

	for i, book := range books {
		responseBook, err := NewBook(&book)
		if err != nil {
			return nil, err
		}

		data[i] = *responseBook
	}

	return data, nil
}

func NewBook(book *data.Book) (*resources.Book, error) {
	if book == nil {
		return nil, nil
	}

	media, err := UnmarshalMedia(book.Banner, book.File)
	if err != nil {
		return nil, err
	}

	media[0].Key = resources.NewKeyInt64(book.ID, resources.BANNERS)
	media[1].Key = resources.NewKeyInt64(book.ID, resources.FILES)

	res := resources.Book{
		Key: resources.NewKeyInt64(book.ID, resources.BOOKS),
		Attributes: resources.BookAttributes{
			Title:           book.Title,
			Description:     book.Description,
			CreatedAt:       book.CreatedAt,
			Price:           book.Price,
			ContractAddress: book.ContractAddress,
			ContractName:    book.ContractName,
			ContractSymbol:  book.ContractSymbol,
			ContractVersion: book.ContractVersion,
			File:            media[0],
			Banner:          media[1],
		},
	}

	return &res, nil
}
