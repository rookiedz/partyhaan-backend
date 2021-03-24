package services

//Data ...
type Data struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
}

//Empty ...
type Empty struct{}

func Success(data interface{}) Data {
	return Data{Status: "success", Data: data}
}

func Total(total int64, data interface{}) Data {
	return Data{Status: "success", Total: total, Data: data}
}

func NotFound() Data {
	return Data{Status: "success", Data: Empty{}}
}

func Failure(err error) Data {
	return Data{Status: "failure", Message: err.Error(), Data: Empty{}}
}

func Err(err error) Data {
	return Data{Status: "error", Message: err.Error(), Data: Empty{}}
}
