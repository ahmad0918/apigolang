package models

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorMap struct {
	Message string `json:"-"`
}

type ErrorRes struct {
	Message  string      `json:"message"`
	Response int         `json:"response"`
	Result   interface{} `json:"result"`
}
