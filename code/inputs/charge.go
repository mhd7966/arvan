package inputs

type ChargeCode struct {
	Name           string `json:"name"`
	Value          int    `json:"value"`
	MaxCapacity    int    `json:"max_capacity"`
	ExpirationDate string `json:"expiration_date"`
}
