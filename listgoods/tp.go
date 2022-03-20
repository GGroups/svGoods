package listgoods

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	ERR "github.com/GGroups/svGoods/comm_err"
)

func CommEncodeResponse(c context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func GoodsListDecodeRequest(c context.Context, request *http.Request) (interface{}, error) {
	if request.Method != "POST" {
		return nil, errors.New("#must POST")
	}
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New(ERR.INPUTE_RROR + err.Error())
	}
	var obj GoodsListRequest
	err = json.Unmarshal(body, &obj)
	if err != nil {
		return nil, errors.New(ERR.INPUTE_RROR + err.Error())
	}
	return obj, nil
}
