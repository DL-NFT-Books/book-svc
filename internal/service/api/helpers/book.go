package helpers

import (
	"net/http"
	"strconv"

	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

const tokenIdIncrementKey = "token_id_increment"

func GetBookByID(r *http.Request, id int64) (*data.Book, error) {
	return BooksQ(r).FilterActual().FilterByID(id).Get()
}

func GetBookListByRequest(r *http.Request, request *requests.GetBooksRequest) ([]data.Book, error) {
	return applyQBooksFilters(BooksQ(r), request).Select()
}

func NewBooksList(books []data.Book) ([]resources.Book, error) {
	bookList := make([]resources.Book, len(books))

	for i, book := range books {
		responseBook, err := NewBook(&book)
		if err != nil {
			return nil, err
		}

		bookList[i] = *responseBook
	}

	return bookList, nil
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
			TokenId:         int32(book.TokenId),
			DeployStatus:    book.DeployStatus,
			File:            media[1],
			Banner:          media[0],
		},
	}

	return &res, nil
}

func GenerateTokenID(r *http.Request) (int64, error) {
	tokenKV, err := KeyValueQ(r).Get(tokenIdIncrementKey)
	if err != nil {
		return 0, err
	}

	if tokenKV == nil {
		tokenKV = &data.KeyValue{
			Key:   tokenIdIncrementKey,
			Value: "0",
		}
	}

	return strconv.ParseInt(tokenKV.Value, 10, 64)
}

func applyQBooksFilters(q data.BookQ, request *requests.GetBooksRequest) data.BookQ {
	if request.Status != nil {
		q = q.FilterByDeployStatus(*request.Status)
	}

	q = q.Page(request.OffsetPageParams)
	q = q.FilterActual()

	return q
}
