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

type TodoController struct{
	TodoRepo repo.ITodoRepo
	Repo repo.IRepo
	Helper helper.IHelper
}

func (controller *TodoController) AddItem(context *gin.Context) {
	var todoRequest request.TodoRequest
	if isSuccess := controller.Helper.BindRequest(context, &todoRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	userIdValue := context.GetString("userId")
	userId, err := strconv.ParseUint(userIdValue, 10, 32)
	if err != nil {
		logger.ErrorLog("An error occured while converting string to uint - Value:", userIdValue, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}

	todo := entities.Todo{UserId: uint(userId), Name: todoRequest.Name, Description: todoRequest.Description, Deadline: todoRequest.Deadline,
		IsCompleted: todoRequest.IsCompleted}

	transaction := controller.Repo.BeginTransaction()
	if isSuccess := controller.TodoRepo.Create(transaction, todo); !isSuccess{
		context.Error(customeError.TodoCouldntCreate)
		return
	}
	transaction.Commit()

	context.JSON(200, response.Success)
}

func (controller *TodoController) GetAllItems(context *gin.Context) {
	userId := context.GetString("userId")
	todoRecords, isSuccess := controller.TodoRepo.GetAll(userId)
	if !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	var todos []response.TodoModelResponse
	for _, todoRecord := range todoRecords {
		todo := response.TodoModelResponse{Id: todoRecord.Id, Name: todoRecord.Name, Description: todoRecord.Description,
			 Deadline: todoRecord.Deadline, IsCompleted: todoRecord.IsCompleted}
		todos = append(todos, todo)
	}

	context.JSON(200, response.TodosResponse{Todos: todos})
}

func (controller *TodoController) DeleteItem(context *gin.Context) {
	todoIdValue := context.Param("todoId")
	todoId, err := strconv.ParseUint(todoIdValue, 10, 32)
	if err != nil {
		logger.ErrorLog("An error occured while converting string to uint - Value:", todoIdValue, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}

	userIdValue := context.GetString("userId")
	userId, err := strconv.ParseUint(userIdValue, 10, 32)
	if err != nil {
		logger.ErrorLog("An error occured while converting string to uint - Value:", userIdValue, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}

	transaction := controller.Repo.BeginTransaction()
	if isSuccess := controller.TodoRepo.Delete(transaction, uint(todoId), uint(userId)); !isSuccess {
		context.Error(customeError.TodoCouldntDelete)
		return
	}
	transaction.Commit()

	context.JSON(200, response.Success)
}

func (controller *TodoController) UpdateItem(context *gin.Context) {
	var todoRequest request.TodoRequest
	if isSuccess := controller.Helper.BindRequest(context, &todoRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	userIdValue := context.GetString("userId")
	userId, err := strconv.ParseUint(userIdValue, 10, 32)
	if err != nil {
		logger.ErrorLog("An error occured while converting string to uint - Value:", userIdValue, "- Error:", err.Error())
		context.Error(customeError.SomethingWentWrong)
		return
	}

	todo := entities.Todo{Id: todoRequest.Id, UserId: uint(userId),  Name: todoRequest.Name, Description: todoRequest.Description,
		 Deadline: todoRequest.Deadline, IsCompleted: todoRequest.IsCompleted}

	transaction := controller.Repo.BeginTransaction()
	if isSuccess := controller.TodoRepo.Update(transaction, todo); !isSuccess {
		context.Error(customeError.TodoCouldntUpdate)
		return
	}
	transaction.Commit()

	context.JSON(200, response.Success)
}

