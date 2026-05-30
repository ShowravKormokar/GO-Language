package types

type DataRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price"`
	Qunatity    int64   `json:"quantity"`
	Status      bool    `json:"status"`
}
