package response

type Meta struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
}

type ErrorResponseDefault struct {
	Meta
}

type SuccessResponseDefault struct {
	Meta 
	Data interface{} `json:"data,omitempty"`
}
