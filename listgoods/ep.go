package listgoods

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
)

type MainCouponsRequest struct {
	Type string `json:"type"`
}

type MainCouponsResponse struct {
	Coupons []Coupons `json:"datas"`
	Msg     string    `json:"msg"`
	RetCode string    `json:"retode"`
}

func MakeCouponsEndPoint(sv ICoupons) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(MainCouponsRequest)
		if !ok {
			return MainCouponsResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(INPUTE_RROR + `not "wx"`)
		}
		return MainCouponsResponse{Coupons: sv.GetCouponsItems(r.Type), Msg: "ok", RetCode: "0"}, nil
	}
}
