package dto

type TaskURI struct {
	GUID string `uri:"guid" binding:"required,uuid"`
}
