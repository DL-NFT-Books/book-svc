package handlers

import (
	"math/big"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"gitlab.com/tokend/nft-books/book-svc/internal/data"
	"gitlab.com/tokend/nft-books/book-svc/internal/data/postgres"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/helpers"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/requests"
	"gitlab.com/tokend/nft-books/book-svc/internal/service/api/responses"
	"gitlab.com/tokend/nft-books/book-svc/internal/signature"
	"gitlab.com/tokend/nft-books/book-svc/resources"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewCreateBookRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Validating info
	banner := request.Data.Attributes.Banner
	file := request.Data.Attributes.File

	if err = helpers.CheckMediaTypes(r, banner.Attributes.MimeType, file.Attributes.MimeType); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Setting banner link
	if err = helpers.SetMediaLink(r, &banner); err != nil {
		logger.WithError(err).Error("failed to set banner link")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	// Setting file link
	if err = helpers.SetMediaLink(r, &file); err != nil {
		logger.WithError(err).Error("failed to set file link")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	media := helpers.MarshalMedia(&banner, &file)
	if media == nil {
		logger.Error("failed to marshal media")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	tokenPrice, ok := big.NewInt(0).SetString(request.Data.Attributes.Price, 10)
	if !ok {
		logger.Error("failed to cast price to big.Int")
		ape.RenderErr(w, problems.BadRequest(errors.New("failed to parse token price"))...)
		return
	}

	lastTokenContractID, err := helpers.GetLastTokenID(r)
	if err != nil {
		logger.WithError(err).Error("failed to get last token id")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Forming signature createInfo
	signatureConfig := helpers.DeploySignatureConfig(r)

	domainData := signature.EIP712DomainData{
		VerifyingAddress: signatureConfig.TokenFactoryAddress,
		ContractName:     signatureConfig.TokenFactoryName,
		ContractVersion:  signatureConfig.TokenFactoryVersion,
		ChainID:          signatureConfig.ChainId,
	}

	// if there is no voucher then passing null address and 0 amount
	voucher := "0x0000000000000000000000000000000000000000"
	voucherAmount := big.NewInt(0)

	if request.Data.Attributes.VoucherToken != nil && request.Data.Attributes.VoucherTokenAmount != nil {
		voucher = *request.Data.Attributes.VoucherToken
		voucherAmount = big.NewInt(*request.Data.Attributes.VoucherTokenAmount)
	}

	createInfo := signature.CreateInfo{
		TokenContractId:      lastTokenContractID + 1,
		TokenName:            request.Data.Attributes.TokenName,
		TokenSymbol:          request.Data.Attributes.TokenSymbol,
		PricePerOneToken:     tokenPrice,
		VoucherTokenContract: voucher,
		VoucherTokensAmount:  voucherAmount,
	}

	// Signing
	createSignature, err := signature.SignCreateInfo(&createInfo, &domainData, signatureConfig)
	if err != nil {
		logger.WithError(err).Error("failed to generate eip712 create signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	// Saving book to the database
	book := data.Book{
		Title:              request.Data.Attributes.Title,
		Description:        request.Data.Attributes.Description,
		CreatedAt:          time.Now(),
		Price:              request.Data.Attributes.Price,
		ContractAddress:    "mocked",
		ContractName:       request.Data.Attributes.TokenName,
		ContractSymbol:     request.Data.Attributes.TokenSymbol,
		ContractVersion:    signatureConfig.TokenFactoryVersion,
		Banner:             media[0],
		File:               media[1],
		Deleted:            false,
		TokenId:            createInfo.TokenContractId,
		DeployStatus:       resources.DeployPending,
		LastBlock:          0,
		VoucherToken:       createInfo.VoucherTokenContract,
		VoucherTokenAmount: *request.Data.Attributes.VoucherTokenAmount,
	}

	db := helpers.DB(r)
	var bookId int64

	if err = db.Transaction(func() error {
		// Inserting book
		bookId, err = db.Books().Insert(book)
		if err != nil {
			return errors.Wrap(err, "failed to save book")
		}

		// Updating last token id
		if err = db.KeyValue().Upsert(data.KeyValue{
			Key:   postgres.TokenIdIncrementKey,
			Value: strconv.FormatInt(createInfo.TokenContractId, 10),
		}); err != nil {
			return errors.Wrap(err, "failed to update last created token id")
		}

		return nil
	},
	); err != nil {
		logger.WithError(err).Error("failed to execute insertion tx")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, responses.NewSignCreateResponse(bookId, createInfo.TokenContractId, *createSignature))
}
