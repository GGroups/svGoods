package cat2nd

import (
	"encoding/json"

	log "github.com/cihub/seelog"

	redis "github.com/gomodule/redigo/redis"
)

const (
	C2ND_REDIS_KEY = "gpwm_cat2nds"
	REDIS_URL      = "redis://127.0.0.1:5101"
)

type ICat2nd interface {
	GetCat2ndItems(c2nd Cat2ndRequest, ret_c2nds *[]Cat2nd) error
	SetCat2ndItems(newlist []Cat2nd) error
}

type Cat2nd struct {
	CategoryId   int    `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	Cat2ndId     int    `json:"cat2ndId"`
	Cat2ndName   string `json:"cat2ndName"`
}

func (s Cat2nd) GetCat2ndItems(c2nd Cat2ndRequest, ret_c2nds *[]Cat2nd) error {
	c2, err := redis.DialURL(REDIS_URL)
	if err != nil {
		return err
	}
	defer c2.Close()
	values, _ := redis.Values(c2.Do("lrange", C2ND_REDIS_KEY, "0", "-1"))
	for _, v := range values {
		var obj Cat2nd
		err = json.Unmarshal(v.([]byte), &obj)
		if err != nil {
			log.Error("##format error", err.Error())
		}
		*ret_c2nds = append(*ret_c2nds, obj)
	}

	return nil
}

func (s Cat2nd) SetCat2ndItems(newlist []Cat2nd) error {
	c2, err := redis.DialURL("redis://127.0.0.1:5101")
	if err != nil {
		return err
	}
	defer c2.Close()

	_, err = c2.Do("del", C2ND_REDIS_KEY)
	if err != nil {
		log.Error("##", err.Error())
	}

	for _, cat := range newlist {
		strval, e1 := json.Marshal(cat)
		if e1 == nil {
			c2.Do("lpush", C2ND_REDIS_KEY, strval)
		} else {
			log.Error("##format error", e1.Error())
		}
	}

	return nil
}
