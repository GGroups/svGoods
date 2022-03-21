package mgoods

import (
	"github.com/jmoiron/sqlx"
)

type MGoodPic struct {
	GoodId int    `json:"goodId"`
	Pic    string `json:"pic"`
}

func (r MGood) GetGoodPix(goodid int) []string {
	return []string{}
}

func (r MGood) UpdateGoodPixDB(piclist []MGoodPic) error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}

	for _, pic := range piclist {
		_, err = db.Exec(SQL_INS_GDINFO, pic.GoodId, pic.Pic)
		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}

func (r MGood) UpdateGoodPixRedis(piclist []MGoodPic) error {
	db, err := sqlx.Open(LITE3, DB_FILE)
	if err != nil {
		return err
	}

	for _, pic := range piclist {
		_, err = db.Exec(SQL_INS_GDINFO, pic.GoodId, pic.Pic)
		if err != nil {
			return err
		}
	}

	db.Close()
	return nil
}
