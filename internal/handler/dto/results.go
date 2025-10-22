package dto

type ErrorResult struct {
	Error *string
}

type BaseResult struct {
	Result any
	Error  *string
}

func strPtr(str string) *string {
	return &str
}
