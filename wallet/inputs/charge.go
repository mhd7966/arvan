package inputs

type Charge struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Code        string `json:"code" binding:"required"`
	CodeType    int    `json:"code_type" binding:"required" default:"1"`
}
