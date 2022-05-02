package handler

import (
	"errors"
	"go-mongo/model"
	"go-mongo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Dummy(c echo.Context) error {
	return c.JSON(http.StatusCreated, errors.New("hello world"))
}

// CreateTweet godoc
// @Summary Create a project
// @Description Create a project
// @ID create-project
// @Tags project
// @Accept  json
// @Produce  json
// @Param project body projectCreateRequest true "Project to create made of text and media"
// @Success 201 {object} projectResponse
// @Failure 404 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /project [post]
func (h *Handler) Create(c echo.Context) error {
	u := model.NewProject()
	req := &projectRegisterRequest{}
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.projectStore.Create(u); err != nil {
		return c.JSON(http.StatusNotFound, utils.NewError(err))
	}
	response := newProjectResponse(u)

	//cookie := new(http.Cookie)
	//cookie.Name = "Token"
	//cookie.Value = response.User.Token
	//cookie.Expires = time.Now().Add(24 * time.Hour)
	//c.SetCookie(cookie)

	//header('Access-Control-Allow-Origin', yourExactHostname);

	//c.Response().Header().Set(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
	//c.Response().Header().Add(echo.HeaderAccessControlAllowCredentials, "true")
	//c.Response().Header().Add(echo.HeaderAccessControlAllowOrigin, "http://localhost:3000")
	//c.Response().Header().Add(echo.HeaderAccessControlAllowHeaders, "Origin, X-Requested-With, Content-Type, Accept")
	//c.Response().Header().
	return c.JSON(http.StatusCreated, response)
}
