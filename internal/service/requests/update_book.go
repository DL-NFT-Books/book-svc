package requests

import "net/http"

type UpdateBookRequest struct {
	ID int64 `json:"id"`
	*CreateBookRequest
}

func NewUpdateBookRequest(r *http.Request) (UpdateBookRequest, error) {
	idReq, err := NewGetBookByIDRequest(r)
	if err != nil {
		return UpdateBookRequest{}, err
	}

	bookReq, err := NewCreateBookRequest(r)
	if err != nil {
		return UpdateBookRequest{}, err
	}

	return UpdateBookRequest{
		ID:                idReq.ID,
		CreateBookRequest: &bookReq,
	}, nil
}
