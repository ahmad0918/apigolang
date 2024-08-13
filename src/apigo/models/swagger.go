package models

// SuccessResponse status 200
type SuccessResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Operation successful"`
	Response int         `json:"response" swaggertype:"integer" example:"200"`
	Result   interface{} `json:"result" swaggertype:"object"`
}

// CommonErrorResponse status 400
type CommonErrorResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Bad request"`
	Response int         `json:"response" swaggertype:"integer" example:"400"`
	Result   interface{} `json:"result" swaggertype:"object"`
}

// UnauthorizedResponse status 401
type UnauthorizedResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Unauthorized access"`
	Response int         `json:"response" swaggertype:"integer" example:"401"`
	Result   interface{} `json:"result" swaggertype:"object"`
}

// ForbiddenErrorResponse status 403
type ForbiddenErrorResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Forbidden access"`
	Response int         `json:"response" swaggertype:"integer" example:"403"`
	Result   interface{} `json:"result" swaggertype:"object"`
}

// DataNotFoundResponse status 404
type DataNotFoundResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Resource not found"`
	Response int         `json:"response" swaggertype:"integer" example:"404"`
	Result   interface{} `json:"result" swaggertype:"object"`
}

// InternalErrorResponse status 500
type InternalErrorResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Internal server error"`
	Response int         `json:"response" swaggertype:"integer" example:"500"`
	Result   interface{} `json:"result" swaggertype:"string" example:""`
}

// ServiceUnavailableResponse status 503
type ServiceUnavailableResponse struct {
	Message  string      `json:"message" swaggertype:"string" example:"Service Unavailable"`
	Response int         `json:"response" swaggertype:"integer" example:"503"`
	Result   interface{} `json:"result" swaggertype:"string" example:""`
}
