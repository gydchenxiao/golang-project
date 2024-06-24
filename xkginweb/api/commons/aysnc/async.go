package aysnc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AsyncTaskInput struct {
	TaskType  string      `json:"task_type"`
	TaskParam interface{} `json:"task_param"`
}

type AsyncTaskOutput struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func HandleAsyncTask(c *gin.Context) {
	var taskInput AsyncTaskInput
	if err := c.ShouldBindJSON(&taskInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 处理异步任务
	taskOutput := processAsyncTask(taskInput)

	c.JSON(http.StatusOK, taskOutput)
}

func processAsyncTask(taskInput AsyncTaskInput) AsyncTaskOutput {
	switch taskInput.TaskType {
	case "send_email":
		// 处理发送电子邮件任务
		go sendEmail(taskInput.TaskParam.(string))
		return AsyncTaskOutput{
			Message: "Task started",
			Data:    nil,
		}
	case "process_image":
		// 处理处理图片任务
		go processImage(taskInput.TaskParam.(string))
		return AsyncTaskOutput{
			Message: "Task started",
			Data:    nil,
		}
	case "generate_report":
		// 处理生成报表任务
		go generateReport(taskInput.TaskParam.(int))
		return AsyncTaskOutput{
			Message: "Task started",
			Data:    nil,
		}
	default:
		return AsyncTaskOutput{
			Message: "Unknown task type",
			Data:    nil,
		}
	}
}

func sendEmail(email string) {
	// 发送电子邮件的逻辑
}

func processImage(imageUrl string) {
	// 处理图片的逻辑
}

func generateReport(reportId int) {
	// 生成报表的逻辑
}
