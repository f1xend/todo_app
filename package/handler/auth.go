package handler

import (
	"net/http"

	"github.com/f1xend/todo-app"
	"github.com/gin-gonic/gin"
)

//	@Summary		SignUp
//	@Tags			auth
//	@description	create account
//	@ID				create-account
//	@Accept			json
//	@Produce		json
//	@Param			input	body		todo.User	true	"account info"
//	@Success		200		{integer}	integer		1
//	@Failure		400		{object}	errorResponce
//	@Failure		404		{object}	errorResponce
//	@Failure		500		{object}	errorResponce
//	@Failure		default	{object}	errorResponce
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" bindings:"required"`
	Password string `json:"password" bindings:"required"`
}

//	@Summary		SignIn
//	@Tags			auth
//	@description	login
//	@ID				login
//	@Accept			json
//	@Produce		json
//	@Param			input	body		signInInput	true	"credentials"
//	@Success		200		{string}	string		"token"
//	@Failure		400		{object}	errorResponce
//	@Failure		404		{object}	errorResponce
//	@Failure		500		{object}	errorResponce
//	@Failure		default	{object}	errorResponce
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
