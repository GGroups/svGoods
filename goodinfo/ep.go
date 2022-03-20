package cat2nd

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	CM "github.com/GGroups/svGoods/comm"
	CME "github.com/GGroups/svGoods/comm_err"
)

type GoodInfRequest struct {
	Type       string `json:"type"`
	CategoryID int    `json:"categoryID"`
}

type GoodInfResponse struct {
	GoodInf []GoodInf `json:"datas"`
	Msg     string    `json:"msg"`
	RetCode string    `json:"retode"`
}

type WGoodInfRequest struct {
	Type     string    `json:"type"`
	GoodInfs []GoodInf `json:"cat2nds"`
}

func MakeGoodInfEndPoint(sv IGoodInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(GoodInfRequest)
		if !ok {
			return GoodInfResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(CME.INPUTE_RROR + `not "wx"`)
		}
		var obj2nds []GoodInf
		e1 := sv.GetGoodInfItems(r, &obj2nds)
		if e1 != nil {
			return GoodInfResponse{GoodInf: nil, Msg: "ok", RetCode: "0"}, e1
		} else {
			return GoodInfResponse{GoodInf: obj2nds, Msg: "ok", RetCode: "0"}, nil
		}
	}
}

func MakeWGoodInfEndPoint(sv IGoodInf) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(WGoodInfRequest)
		if !ok {
			return GoodInfResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(CME.INPUTE_RROR + `not "wx"`)
		}
		e := sv.SetGoodInfItems(r.GoodInfs)
		if e == nil {
			return CM.WCommResponse{Msg: CME.KEY_OK, RetCode: CME.OKEY_CODE}, nil
		} else {
			return CM.WCommResponse{Msg: CME.KEY_ERR, RetCode: CME.SVR_ERROR}, nil
		}
	}
}
