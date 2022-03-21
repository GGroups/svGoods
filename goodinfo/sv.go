package goodinfo

import (
	"encoding/json"
	"strconv"

	log "github.com/cihub/seelog"

	redis "github.com/gomodule/redigo/redis"
)

const (
	GOODINFO_REDIS_KEY = "gpwm_goodinfos"
	REDIS_URL          = "redis://127.0.0.1:5101"
)

type IGoodInf interface {
	GetGoodInf(goodId int, ret_c2nds *GoodInf) error
	SetGoodInfItems(newlist []GoodInf) error
}

type GoodInf struct {
	GoodID      int      `json:"goodID"`
	GoodDesc    string   `json:"goodDesc"`
	SizeList    []string `json:"sizeList"`
	ColorList   []string `json:"colorList"`
	GoodPicList []string `json:"goodPicList"`
}

func (s GoodInf) GetGoodInf(goodId int, goodinfo *GoodInf) error {
	c2, err := redis.DialURL(REDIS_URL)
	if err != nil {
		return err
	}
	defer c2.Close()
	redis_key := GOODINFO_REDIS_KEY + strconv.Itoa(goodId)
	val, _ := c2.Do("get", redis_key)
	err = json.Unmarshal(val.([]byte), goodinfo)
	if err != nil {
		log.Error("##format error", err.Error())
	}

	return nil
}

func (s GoodInf) SetGoodInfItems(newlist []GoodInf) error {
	c2, err := redis.DialURL("redis://127.0.0.1:5101")
	if err != nil {
		return err
	}
	defer c2.Close()

	for _, cat := range newlist {
		strval, e1 := json.Marshal(cat)
		if e1 == nil {
			redis_key := GOODINFO_REDIS_KEY + strconv.Itoa(cat.GoodID)
			c2.Do("set", redis_key, strval)
		} else {
			log.Error("##format error", e1.Error())
		}
	}

	return nil
}
