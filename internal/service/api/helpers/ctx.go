package helpers

import (
	"context"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"net/http"

	"gitlab.com/tokend/nft-books/book-svc/internal/config"
	"gitlab.com/tokend/nft-books/doorman/connector"
	networker "gitlab.com/tokend/nft-books/network-svc/connector"

	"gitlab.com/distributed_lab/logan/v3"
	s3connector "gitlab.com/tokend/nft-books/blob-svc/connector/api"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	mimeTypesCtxKey
	deploySignatureCtxKey
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

func CtxDeploySignature(entry *config.DeploySignatureConfig) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, deploySignatureCtxKey, entry)
	}
}

func DeploySignatureConfig(r *http.Request) *config.DeploySignatureConfig {
	return r.Context().Value(deploySignatureCtxKey).(*config.DeploySignatureConfig)
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

func Networker(r *http.Request) networker.Connector {
	return r.Context().Value(networkerConnectorCtxKey).(networker.Connector)
}
