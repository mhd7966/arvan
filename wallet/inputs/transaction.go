package inputs

type GetTransactionsPagination struct {
	Page   int `query:"page" default:"1"`
	Size   int `query:"size" default:"10"`
}
