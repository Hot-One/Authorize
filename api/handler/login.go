package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Procedure json
// @Param login body models.CreateUser true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Login(c *gin.Context) {
	var login models.CreateUser

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Username: login.Username})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "User does not exist", http.StatusInternalServerError, nil)
			return
		}
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}
	// if resp.Password != login.Password {
	// 	h.handlerResponse(c, "Wrong password", http.StatusInternalServerError, nil)
	// 	return
	// }
	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": resp.Id,
	}, time.Hour*360, h.cfg.SecretKey)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}
