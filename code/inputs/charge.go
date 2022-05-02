package inputs

type ChargeCode struct {
	Name           string `json:"name"`
	Value          int    `json:"value" validate:"gt=0"`
	MaxCapacity    int    `json:"max_capacity" validate:"gt=0"`
	ExpirationDate string `json:"expiration_date"`
}
