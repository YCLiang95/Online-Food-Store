package protocal

import "time"

type Order struct {
	Oid                   int64     `xorm:"pk autoincr 'oid'" json:"oid" `
	Uid                   int64     `json:"uid"`
	Status                int64       `json:"status"`
	OrderId               string    `json:"order_id"`
	CreateTime            time.Time `json:"create_time"`
	UpdateTime            time.Time
	TotalMerchandisePrice float64   `json:"total_merchandise_price"`
	TotalPaymentPrice     float64   `json:"total_payment_price"`
	PaymentMethod         int       `json:"payment_method"`
}

type OrderResponse struct {
	Order       Order         `json:"order"`
	Delivery    Delivery      `json:"delivery"`
	OrderDetail []*OrderDetail `json:"order_detail"`
}
type Delivery struct {
	Did        int64     `xorm:"pk autoincr 'did' json:"did""`
	State      string    `json:"state"`
	City       string    `json:"city"`
	Zipcode    int64       `json:"zipcode"`
	Address    string    `json:"address"`
	CreateTime time.Time `json:"create_time"`
	Status     int64      `json:"status"`
	Oid        int64     `json:"oid"`
	UpdateTime time.Time `json:"update_time"`
	DeliveryId string    `json:"delivery_id"`
}

type DeliveryRequest struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Zipcode int    `json:"zipcode"`
	Address string `json:"address"`
}

type OrderDetail struct {
	Odid             int64   `json:"odid"`
	Oid              int64   `json:"oid"`
	MerchandiseName  string  `json:"merchandise_name"`
	MerchandisePrice float64 `json:"merchandise_price"`
	MerchandiseCount int64     `json:"merchandise_count"`
	MerchandiseId    int64   `json:"merchandise_id"`
}

type OrderRequest struct {
	Uid            int64  `json:"uid"`
	OrderDetail    string `json:"order_detail"`
	PaymentType    int    `json:"payment_type"`
	DeliveryDetail string `json:"delivery_detail"`
}

type OrderReequestDetail struct {
	Mid   int64 `json:"mid"`
	Count int   `json:"count"`
}
