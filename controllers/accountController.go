package controllers

import (
	"app/GoSample/controllers/helper"
	"app/GoSample/controllers/models/request"
	"app/GoSample/controllers/models/response"
	"app/GoSample/db/entities"
	"app/GoSample/db/repo"
	"app/GoSample/infra/auth"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/customeError"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (u *AccountController) Register(context *gin.Context) {
	var accountRequest request.AccountRequest
	if isSuccess := helper.BindRequest(context, &accountRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	hashedPassword, isSuccess := helper.HashPassword(context, accountRequest.Password)
	if !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	account := entities.Account{UserName: accountRequest.UserName, Password: hashedPassword}

	if isUserNameExist := repo.Account.IsUserNameExist(account.UserName); isUserNameExist {
		context.Error(customeError.UserAlreadyExists)
		return
	}

	transaction := repo.BeginTransaction()
	userId := repo.Account.Create(transaction, account)
	if userId == 0 {
		context.Error(customeError.WrongCredentials)
		return
	}
	transaction.Commit()

	if isSuccess := helper.AddToSession(context, constant.UserName, accountRequest.UserName); !isSuccess {
		context.Error(customeError.UserAlreadyExists)
		return
	}

	token, err := auth.JWT.GenerateToken(userId)
	if err != nil {
		context.Error(customeError.WrongCredentials)
		return
	}

	context.JSON(200, response.AccountResponse{
			response.BaseResponse{IsSuccess: true,},
			token,
	})
}

func (u *AccountController) Login(context *gin.Context) {
	var accountRequest request.AccountRequest
	if isSuccess := helper.BindRequest(context, &accountRequest); !isSuccess{
		context.Error(customeError.SomethingWentWrong)
		return
	}

	account, isSuccess := repo.Account.FirstByUserName(accountRequest.UserName)
	if !isSuccess {
		context.Error(customeError.WrongCredentials)
		return
	}

	if isValidPassword := helper.CheckPasswordHash(accountRequest.Password, account.Password); !isValidPassword {
		context.Error(customeError.WrongCredentials)
		return
	}

	if isSuccess := helper.AddToSession(context, constant.UserName, accountRequest.UserName); !isSuccess {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	token, err := auth.JWT.GenerateToken(account.Id)
	if err != nil {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	context.JSON(200, response.AccountResponse{
		response.BaseResponse{IsSuccess: true,},
		token,
	})
}
