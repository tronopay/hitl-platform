package httpd

import (
	"fmt"
	"html/template"
	"net/http"

	// "tronopay/internal/entity"
	"tronopay/internal/service"

	// "tronopay/internal/httpd/dto"
	_ "tronopay/internal/httpd/dto"

	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
	qrcode "github.com/skip2/go-qrcode"
)

type Controller struct {
	pages    map[string]*template.Template
	router   *gin.Engine
	services *service.Services
}

func (controller *Controller) MainPage(c *gin.Context) {
	controller.router.SetHTMLTemplate(controller.pages["index.html"])
	c.HTML(http.StatusOK, "index", gin.H{
		"donate": controller.services.App.Donate,
	})
}

func (controller *Controller) GenerateQRCode(c *gin.Context) {
	value, ok := c.Params.Get("value")
	if !ok {
		// log
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	png, err := qrcode.Encode(value, qrcode.Highest, 400)
	if err != nil {
		// log
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}

	c.Data(http.StatusOK, "image/png", png)
}

// @Summary      Create a task
// @Description  To create a new task, you must provide the mandatory fields: description, type, price, and currency.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        description body string true "Detailed task description for the person performing the work" "Take a several photo of a car"
// @Param        type body string true "Type of the task" enums(TAKE_PHOTO,TAKE_VIDEO,OTHER)
// @Param        price body string true "Price of the task" "10.0"
// @Param        currency body string true "Currency of the price" "TRX"
// @Success      201  {object} dto.ShowTaskResponse
// @Failure		 422  {object} ValidationError
// @Failure      500  {object} InternalError
// @Router       /tasks [post]
func (controller *Controller) CreateTask(c *gin.Context) {
	notImplemented(c)

	// var createTask dto.CreateTaskRequest
	// if err := c.ShouldBindBodyWithJSON(&createTask); err != nil {
	// 	errMsg := fmt.Sprintf("create task - bad param error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
	// 	return
	// }

	// guid, err := uuid.NewV7()
	// if err != nil {
	// 	errMsg := fmt.Sprintf("create task - guid generation error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
	// 	return
	// }

	// taskNew := entity.Task{
	// 	Id:          guid,
	// 	Status:      entity.TaskStatusNew,
	// 	Type:        entity.TaskType(createTask.Type),
	// 	Description: createTask.Description,
	// }

	// err = controller.services.TaskService.Create(c.Request.Context(), taskNew)
	// if err != nil {
	// 	errMsg := fmt.Sprintf("create task error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
	// 	return
	// }

	// c.JSON(http.StatusCreated, dto.ShowTaskResponse{
	// 	Id:          taskNew.Id.String(),
	// 	Status:      string(taskNew.Status),
	// 	Type:        string(taskNew.Type),
	// 	Description: taskNew.Description,
	// 	Message:     "Please make the payment within 1 hour",
	// 	Address:     "Ts...123",
	// 	Price:       "10.0",
	// })
}

// @Summary      Show the list of the task
// @Description  Get the list of the tasks
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200  {object} dto.ShowTasksResponse
// @Failure      500  {object} InternalError
// @Router       /tasks [get]
func (controller *Controller) ShowTaskList(c *gin.Context) {
	notImplemented(c)

	// tasks, err := controller.services.TaskService.GetTaskList(c.Request.Context())
	// if err != nil {
	// 	errMsg := fmt.Sprintf("show task list error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
	// 	return
	// }

	// response := make([]dto.ShowTaskResponse, 0)
	// for _, task := range tasks {
	// 	response = append(response, dto.ShowTaskResponse{
	// 		Id:      task.Id.String(),
	// 		Type:    string(task.Type),
	// 		Status:  string(task.Status),
	// 		Message: "",
	// 		Address: "Tu...123",
	// 		Price:   "10.0",
	// 	})
	// }
	// c.JSON(http.StatusOK, response)

}

// @Summary      Show a task
// @Description  Get task info by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        guid path     string  true  "Task ID"
// @Success      200  {object} dto.ShowTaskResponse
// @Failure      400  {object} BadParam
// @Failure      404  {object} NotFound
// @Failure      500  {object} InternalError
// @Router       /tasks/{guid} [get]
func (controller *Controller) ShowTask(c *gin.Context) {
	notImplemented(c)

	// var uri dto.TaskURI
	// if err := c.ShouldBindUri(&uri); err != nil {
	// 	errMsg := fmt.Sprintf("show task - bad param error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusBadRequest, BadParam{Code: string(http.StatusBadRequest), Message: errMsg})
	// 	return
	// }

	// taskUUID, err := uuid.Parse(uri.GUID)
	// if err != nil {
	// 	errMsg := fmt.Sprintf("show task - parse UUID error: %v", err)
	// 	controller.services.Log.Warn(errMsg)
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": errMsg})
	// 	return
	// }

	// task, err := controller.services.TaskService.GetTaskById(c.Request.Context(), taskUUID)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusOK, dto.ShowTaskResponse{
	// 	Id:      task.Id.String(),
	// 	Type:    string(task.Type),
	// 	Status:  string(task.Status),
	// 	Message: "Please wait until the task is completed.",
	// 	Address: "Tu...123",
	// 	Price:   "10.0",
	// })
}

// @Summary      Delete the task
// @Description  Delete the task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        guid path     string  true  "Task GUID"
// @Success      202  {object} dto.DeleteTaskResponse
// @Failure      404  {object} NotFound
// @Failure      500  {object} InternalError
// @Router       /tasks/{guid} [delete]
func (controller *Controller) DeleteTask(c *gin.Context) {
	notImplemented(c)

	// var uri dto.TaskURI
	// if err := c.ShouldBindUri(&uri); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "GUID is not valid"})
	// 	return
	// }

	// taskUUID, err := uuid.Parse(uri.GUID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse UUID"})
	// 	return
	// }

	// err = controller.services.TaskService.Delete(c.Request.Context(), taskUUID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusAccepted, dto.DeleteTaskResponse{
	// 	Success: true,
	// })
}

func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented yet. Coming soon"})
}
