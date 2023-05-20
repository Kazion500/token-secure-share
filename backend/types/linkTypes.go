package types

type DataType struct {
	Link string `json:"link"`
}

type ErrorType struct {
	Error string `json:"error"`
}

type GType interface {
	string | ErrorType | DataType
}

type ResponseType[T any] struct {
	Success bool `json:"success"`
	Data    T    `json:"data"`
}

type RequestBody struct {
	Data string `json:"data"`
}
