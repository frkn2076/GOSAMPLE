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

type AccountController struct{}

func (u *AccountController) Register(c *gin.Context) {
	var accountRequest request.AccountRequest
	helper.BindRequest(c, &accountRequest)

	hashedPassword := helper.HashPassword(c, accountRequest.Password)

	account := entities.Account{UserName: accountRequest.UserName, Password: hashedPassword}

	if isUserNameExist := repo.Account.IsUserNameExist(account.UserName); isUserNameExist {
		c.Error(customeError.UserAlreadyExists)
		return
	}

	tx := repo.BeginTransaction()
	repo.Account.Create(tx, account)
	tx.Commit()

	helper.AddToSession(c, constant.UserName, accountRequest.UserName)

	c.JSON(200, response.Success)
}

func (u *AccountController) Login(c *gin.Context) {
	var accountRequest request.AccountRequest
	helper.BindRequest(c, &accountRequest)

	hashedPassword := helper.HashPassword(c, accountRequest.Password)

	account := entities.Account{UserName: accountRequest.UserName, Password: hashedPassword}

	if isExist := repo.Account.IsExist(account); !isExist {
		c.Error(customeError.WrongCredentials)
		return
	}

	helper.AddToSession(c, constant.UserName, accountRequest.UserName)

	c.JSON(200, response.Success)
}
