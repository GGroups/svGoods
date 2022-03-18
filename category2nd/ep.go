package cat2nd

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"

	CM "github.com/GGroups/svGoods/comm"
	CME "github.com/GGroups/svGoods/comm_err"
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

type WCat2ndRequest struct {
	Type    string   `json:"type"`
	Cat2nds []Cat2nd `json:"cat2nds"`
}

func MakeCat2ndEndPoint(sv ICat2nd) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(Cat2ndRequest)
		if !ok {
			return Cat2ndResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(CME.INPUTE_RROR + `not "wx"`)
		}
		var obj2nds []Cat2nd
		e1 := sv.GetCat2ndItems(r, obj2nds)
		if e1 != nil {
			return Cat2ndResponse{Cat2nd: nil, Msg: "ok", RetCode: "0"}, e1
		} else {
			return Cat2ndResponse{Cat2nd: obj2nds, Msg: "ok", RetCode: "0"}, nil
		}
	}
}

func MakeWCat2ndEndPoint(sv ICat2nd) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(WCat2ndRequest)
		if !ok {
			return Cat2ndResponse{}, nil
		}
		if r.Type != "wx" {
			return nil, errors.New(CME.INPUTE_RROR + `not "wx"`)
		}
		e := sv.SetCat2ndItems(r.Cat2nds)
		if e == nil {
			return CM.WCommResponse{Msg: CME.KEY_OK, RetCode: CME.OKEY_CODE}, nil
		} else {
			return CM.WCommResponse{Msg: CME.KEY_ERR, RetCode: CME.SVR_ERROR}, nil
		}
	}
}
