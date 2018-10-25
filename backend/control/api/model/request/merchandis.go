package request


type MerchandisRequest struct {
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantify"`
	ImageUrl string  `json:"imagerl"`
}

type MerchandisUpdateRequest struct {
	Mid      int64   `js`
	Name     string  `json:"name"`
	Weight   float64 `json:"weight"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantify"`
	ImageUrl string  `json:"imagerl"`
}