package handler

import (
	"go-mongo/model"
	"go-mongo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CreateTweet godoc
// @Summary Create an tweet
// @Description Create an tweet
// @ID create-tweet
// @Tags tweet
// @Accept  json
// @Produce  json
// @Param tweet body tweetCreateRequest true "Tweet to create made of text and media"
// @Success 201 {object} singleTweetResponse
// @Failure 404 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /tweets [post]
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
