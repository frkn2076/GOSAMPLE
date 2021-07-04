package controllers

import (
	"app/GoSample/controllers/helper"
	"app/GoSample/controllers/models/request"
	"app/GoSample/controllers/models/response"
	"app/GoSample/db/entities"
	"app/GoSample/db/repo"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/customeError"

	"github.com/gin-gonic/gin"
)

type AccountController struct{
	AccountRepo repo.IAccountrepo
	Repo repo.IRepo
	Helper helper.IHelper
}

func (controller *AccountController) Register(context *gin.Context) {
	var accountRequest request.AccountRequest
	if isSuccess := controller.Helper.BindRequest(context, &accountRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	hashedPassword, isSuccess := controller.Helper.HashPassword(context, accountRequest.Password)
	if !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	account := entities.Account{UserName: accountRequest.UserName, Password: hashedPassword}

	if isUserNameExist := controller.AccountRepo.IsUserNameExist(account.UserName); isUserNameExist {
		context.Error(customeError.UserAlreadyExists)
		return
	}

	transaction := controller.Repo.BeginTransaction()
	userId := controller.AccountRepo.Create(transaction, account)
	if userId == 0 {
		context.Error(customeError.SomethingWentWrong)
		return
	}
	transaction.Commit()

	if isSuccess := controller.Helper.AddToSession(context, constant.UserName, accountRequest.UserName); !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	token, isSuccess := controller.Helper.GenerateToken(account.Id)
	if !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	context.JSON(200, response.AccountResponse{
			BaseResponse: response.BaseResponse{IsSuccess: true,},
			Token: token,
	})
}

func (controller *AccountController) Login(context *gin.Context) {
	var accountRequest request.AccountRequest
	if isSuccess := controller.Helper.BindRequest(context, &accountRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	account, isSuccess := controller.AccountRepo.FirstByUserName(accountRequest.UserName)
	if !isSuccess {
		context.Error(customeError.WrongCredentials)
		return
	}

	if isValidPassword := controller.Helper.CheckPasswordHash(accountRequest.Password, account.Password); !isValidPassword {
		context.Error(customeError.WrongCredentials)
		return
	}

	if isSuccess := controller.Helper.AddToSession(context, constant.UserName, accountRequest.UserName); !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	token, isSuccess := controller.Helper.GenerateToken(account.Id)
	if !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	context.JSON(200, response.AccountResponse{
		BaseResponse: response.BaseResponse{IsSuccess: true,},
		Token: token,
	})
}
