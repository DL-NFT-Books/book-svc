package helpers

import (
	"context"
	"github.com/dl-nft-books/book-svc/internal/data"
	"net/http"

	"github.com/dl-nft-books/book-svc/internal/config"
	"github.com/dl-nft-books/doorman/connector"
	networker "github.com/dl-nft-books/network-svc/connector"

	s3connector "github.com/dl-nft-books/blob-svc/connector/api"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	mimeTypesCtxKey
	doormanConnectorCtxKey
	documenterConnectorCtxKey
	networkerConnectorCtxKey
	dbKey
)

func CtxDB(db data.DB) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbKey, db)
	}
}

func DB(r *http.Request) data.DB {
	return r.Context().Value(dbKey).(data.DB).New()
}

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
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

func CtxNetworkerConnector(entry networker.Connector) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, networkerConnectorCtxKey, entry)
	}
}
func Networker(r *http.Request) networker.Connector {
	return r.Context().Value(networkerConnectorCtxKey).(networker.Connector)
}
