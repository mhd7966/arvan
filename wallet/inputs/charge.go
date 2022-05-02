package inputs

type Charge struct {
	PhoneNumber string `json:"phone_number"`
	Code        string `json:"code"`
	CodeType    int    `json:"code_type" default:"1"`
}
