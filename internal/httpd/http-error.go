package httpd

type NotFound struct {
	Code    string `json:"code" binding:"required" example:"404"`
	Message string `json:"message" binding:"required" example:"Not found"`
}

type BadParam struct {
	Code    string `json:"code" binding:"required" example:"400"`
	Message string `json:"message" binding:"required" example:"Bad param"`
}

type ValidationError struct {
	Code    string `json:"code" binding:"required" example:"422"`
	Message string `json:"message" binding:"required" example:"Validation error"`
}

type InternalError struct {
	Code    string `json:"code" binding:"required" example:"500"`
	Message string `json:"message" binding:"required" example:"Internal error"`
}
