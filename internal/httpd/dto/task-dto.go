package dto

type CreateTaskRequest struct {
	Type        string `json:"type" binding:"required" example:"TAKE_PHOTO"`
	Description string `json:"task_statement" binding:"required,min=3,max=1000" example:"Take a several photo of a car"`
	Price       string `json:"price_trx" binding:"required" example:"100.0"`
}

type DeleteTaskResponse struct {
	Success bool `json:"success" binding:"required" example:"true"`
}

type ShowTaskResponse struct {
	Id          string `json:"id" binding:"required,guid" example:"01a09bc4-29dc-7cb3-8966-93d058554dbb"`
	Status      string `json:"task_status" binding:"required" example:"PENDING_PAYMENT"`
	Type        string `json:"task_type" binding:"required" example:"TAKE_PHOTO"`
	Description string `json:"task_statement" binding:"required" example:"Take a several photo of a car"`
	Message     string `json:"hint_message" example:"Please make the payment within 1 hour"`
	Address     string `json:"escrow_address_trx" binding:"required"  example:"TU...123"`
	Price       string `json:"escrow_price_trx" binding:"required" example:"10.0"`
}
type ShowTasksResponse = []ShowTaskResponse
