package controllers

import (
	"app/GoSample/controllers/helper"
	"app/GoSample/controllers/models/request"
	"app/GoSample/controllers/models/response"
	"app/GoSample/db/entities"
	"app/GoSample/db/repo"
	"app/GoSample/infra/constant"
	"app/GoSample/infra/customeError"
	"app/GoSample/infra/auth"

	"github.com/gin-gonic/gin"
)

type AccountController struct{}

func (u *AccountController) Register(context *gin.Context) {
	var accountRequest request.AccountRequest
	helper.BindRequest(context, &accountRequest)

	hashedPassword := helper.HashPassword(context, accountRequest.Password)

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

	helper.AddToSession(context, constant.UserName, accountRequest.UserName)

	token, err := auth.JWT.GenerateToken(userId)
	if err != nil {
		context.Error(customeError.WrongCredentials)
		return
	}

	context.JSON(200, response.BaseResponse{IsSuccess: true, Token: token})
}

func (u *AccountController) Login(context *gin.Context) {
	var accountRequest request.AccountRequest
	helper.BindRequest(context, &accountRequest)

	account, isSuccess := repo.Account.FirstByUserName(accountRequest.UserName)
	if !isSuccess {
		context.Error(customeError.WrongCredentials)
		return
	}

	isValidPassword := helper.CheckPasswordHash(accountRequest.Password, account.Password)

	if !isValidPassword {
		context.Error(customeError.WrongCredentials)
		return
	}

	helper.AddToSession(context, constant.UserName, accountRequest.UserName)

	token, err := auth.JWT.GenerateToken(account.Id)
	if err != nil {
		context.Error(customeError.SomethingWentWrong)
		return
	}

	context.JSON(200, response.BaseResponse{IsSuccess: true, Token: token})
}
