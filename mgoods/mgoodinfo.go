package mgoods

type MGoodPic struct {
	GoodId int    `json:"goodId"`
	Pic    string `json:"pic"`
}

func (r MGood) GetGoodPix(goodid int) []string {
	return []string{}
}

func (r MGood) UpdateGoodPix(goodid int, picList []string) error {
	return nil
}
