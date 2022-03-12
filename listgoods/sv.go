package listgoods

const (
	INPUTE_RROR = "#input format Error#"
)

type ICoupons interface {
	GetCouponsItems(name string) []Coupons
}

type Coupons struct {
	ValidateType int    `json:"validityType"`
	ValidityTime int    `json:"validityTime"`
	Now          int    `json:"now"`
	End          int    `json:"end"`
	Limitprice   int    `json:"limitPrice"`
	CouponPrice  int    `json:"couponPrice"`
	Key          string `json:"key"`
	Remark       string `json:"remark"`
	Couponid     string `json:"couponId"`
}

func (s Coupons) GetCouponsItems(t string) []Coupons {
	ss := []Coupons{
		{ValidateType: 2, ValidityTime: 1, Now: 0, End: 1, Limitprice: 10, CouponPrice: 3, Key: "1", Remark: "不能叠使用", Couponid: "1"},
	}
	return ss
}
