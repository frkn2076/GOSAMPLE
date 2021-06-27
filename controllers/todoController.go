package controllers

import (
	"strconv"

	"app/GoSample/controllers/helper"
	"app/GoSample/controllers/models/request"
	"app/GoSample/controllers/models/response"
	"app/GoSample/db/entities"
	"app/GoSample/db/repo"
	"app/GoSample/logger"
	"app/GoSample/infra/customeError"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func (u *TodoController) AddItem(context *gin.Context) {
	var todoRequest request.TodoRequest
	helper.BindRequest(context, &todoRequest)

	userIdValue := context.GetString("userId")
	userId, err := strconv.ParseUint(userIdValue, 10, 32)
	if err != nil {
		logger.ErrorLog("An error occured while converting string to uint - Value:", userIdValue, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}

	todo := entities.Todo{UserId: uint(userId), Name: todoRequest.Name, Description: todoRequest.Description, Deadline: todoRequest.Deadline,
		IsCompleted: todoRequest.IsCompleted}

	transaction := repo.BeginTransaction()
	repo.Todo.Create(transaction, todo)
	transaction.Commit()

	context.JSON(200, response.Success)
}

func (u *TodoController) GetAllItems(context *gin.Context) {
	var todoRequest request.TodoRequest
	helper.BindRequest(context, &todoRequest)

	userId := context.GetString("userId")
	todoRecords, err := repo.Todo.GetAll(userId)
	if err != nil {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	var todos []response.TodoModelResponse
	for _, todoRecord := range todoRecords {
		todo := response.TodoModelResponse{Name: todoRecord.Name, Description: todoRecord.Description, Deadline: todoRecord.Deadline,
			 IsCompleted: todoRecord.IsCompleted}
		todos = append(todos, todo)
	}

	context.JSON(200, response.TodosResponse{Todos: todos})
}

