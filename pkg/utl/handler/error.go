package handler

//Error - Error
type Error struct {
	Data       ValidationRes `json:"data"`
	Message    string        `json:"message"`
	ErrorCode  int           `json:"errorCode"`
	StatusCode int           `json:"statusCode"`
}

//ObjectError - ObjectError
type ObjectError struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Code  string `json:"code"`
}

//ValidationRes - ValidationRes
type ValidationRes struct {
	Errors []ObjectError `json:"errors"`
}

//MessageFlags - MessageFlags
var MessageFlags = map[int]string{
	200: "OK",
	201: "Created",
	202: "Accepted",
	204: "NoContent",
	400: "Bad request",
	401: "Unauthorized",
	403: "Forbidden",
	405: "Method Not Allowed",
	500: "Internal Server Error",
	502: "Bad gateway",
}

//GetMessage - GetMessage
func GetMessage(code int) string {
	return MessageFlags[code]
}
