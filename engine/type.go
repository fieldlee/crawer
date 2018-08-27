package engine

type Request struct {
	Url string
	ParseFun func([]byte)ParseRequest
}

type ParseRequest struct {
	Requests []Request
	Items []interface{}
}

func NilParseFunc([]byte)ParseRequest  {
	return ParseRequest{}
}