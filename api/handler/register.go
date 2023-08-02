package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
)

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Register
// @Description Register
// @Tags Register
// @Accept json
// @Procedure json
// @Param register body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Register(c *gin.Context) {

	var createUser models.CreateUser
	var id string
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createUser.Password) < 7 {
		h.handlerResponse(c, "Password should inculude more than 7 elements", http.StatusBadRequest, errors.New("Password len should inculude more than 8 elements"))
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Username: createUser.Username})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.User().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "User already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "User already exist", http.StatusBadRequest, nil)
		return
	}
	resp, err = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})

	h.handlerResponse(c, "create user resposne", http.StatusCreated, resp)
}
