package helpers

import (
	"context"
	"net/http"

	"gitlab.com/tokend/nft-books/doorman/connector"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	booksQCtxKey
	mimeTypesCtxKey
	doormanConnectorCtxKey
	documenterConnectorCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxBooksQ(entry data.BookQ) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, booksQCtxKey, entry)
	}
}

func CtxMimeTypes(entry *config.MimeTypes) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, mimeTypesCtxKey, entry)
	}
}

func MimeTypes(r *http.Request) *config.MimeTypes {
	return r.Context().Value(mimeTypesCtxKey).(*config.MimeTypes)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BooksQ(r *http.Request) data.BookQ {
	return r.Context().Value(booksQCtxKey).(data.BookQ).New()
}

func CtxDoormanConnector(entry connector.ConnectorI) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, doormanConnectorCtxKey, entry)
	}
}
func DoormanConnector(r *http.Request) connector.ConnectorI {
	return r.Context().Value(doormanConnectorCtxKey).(connector.ConnectorI)
}

func CtxDocumenterConnector(entry s3connector.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, documenterConnectorCtxKey, entry)
	}
}

func DocumenterConnector(r *http.Request) s3connector.Connector {
	return r.Context().Value(documenterConnectorCtxKey).(s3connector.Connector)
}
