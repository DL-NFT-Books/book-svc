package helpers

import (
	"context"
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	booksQCtxKey
	jwtCtxKey
	mimeTypesCtxKey
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

func CtxJWT(entry *config.JWT) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, jwtCtxKey, entry)
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

func JWT(r *http.Request) *config.JWT {
	return r.Context().Value(jwtCtxKey).(*config.JWT)
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func BooksQ(r *http.Request) data.BookQ {
	return r.Context().Value(booksQCtxKey).(data.BookQ).New()
}
