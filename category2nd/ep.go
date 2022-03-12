package cat2nd

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type Cat2ndRequest struct {
	Type       string `json:"type"`
	CategoryID int    `json:"categoryID"`
}

type Cat2ndResponse struct {
	Cat2nd  []Cat2nd `json:"datas"`
	Msg     string   `json:"msg"`
	RetCode string   `json:"retode"`
}

func MakeCat2ndEndPoint(sv ICat2nd) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(Cat2ndRequest)
		if !ok {
			return Cat2ndResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(INPUTE_RROR + `not "wx"`)
		}
		return Cat2ndResponse{Cat2nd: sv.GetCat2ndItems(r), Msg: "ok", RetCode: "0"}, nil
	}
}
