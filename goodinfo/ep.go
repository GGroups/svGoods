package goodinfo

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	CM "github.com/GGroups/svGoods/comm"
	CME "github.com/GGroups/svGoods/comm_err"
)

type GoodInfRequest struct {
	Type   string `json:"type"`
	GoodID int    `json:"goodID"`
}

type GoodInfResponse struct {
	GoodInf GoodInf `json:"goodInfo"`
	Msg     string  `json:"msg"`
	RetCode string  `json:"retode"`
}

type WGoodInfRequest struct {
	Type     string    `json:"type"`
	GoodInfs []GoodInf `json:"goodInfs"`
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
		var goodinfo GoodInf
		e1 := sv.GetGoodInf(r.GoodID, &goodinfo)
		if e1 != nil {
			return GoodInfResponse{GoodInf: GoodInf{}, Msg: "ok", RetCode: "0"}, e1
		} else {
			return GoodInfResponse{GoodInf: goodinfo, Msg: "ok", RetCode: "0"}, nil
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
