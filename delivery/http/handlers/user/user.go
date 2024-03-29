package user_handler

import (
	"github.com/labstack/echo/v4"
	param "hangout/param/http"
	"hangout/pkg/claims"
	"hangout/pkg/httperr"
	"net/http"
)

// Register godoc
//
//	@Summary		Register account
//	@Description	Create a new account for new user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.RegisterRequest	true	"Create User"
//	@Success		201		{object}	param.RegisterResponse
//	@Router			/signup [post]
func (h Handler) Register(c echo.Context) error {
	var req param.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateRegisterRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.userSvc.Register(c.Request().Context(), req)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

// Login godoc
//
//	@Summary		Login account
//	@Description	Login to user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		param.LoginRequest	true	"Login User"
//	@Success		201		{object}	param.LoginResponse
//	@Router			/login [post]
func (h Handler) Login(c echo.Context) error {
	var req param.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateLoginRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res, err := h.userSvc.Login(c.Request().Context(), req)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// UploadProfileImage godoc
//
//	@Summary		Upload profile image
//	@Description	User can set profile image
//	@Tags			profile-image
//	@Security		auth
//	@Accept			json
//	@Produce		json
//	@Param			image	formData file	true							"The image to upload"
//	@Param			image	body		param.SaveProfileImageRequest	false	"Request object"
//	@Success		201		{object}	param.SaveProfileImageResponse
//	@Router			/user/profile_img [post]
func (h Handler) UploadProfileImage(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)
	img, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	res, err := h.userSvc.SaveProfileImage(c.Request().Context(), param.SaveProfileImageRequest{}, img, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusCreated, res)
}

// GetPrimaryImage godoc
//
//	@Summary		Get primary profile image
//	@Description	User can get it's newly(primary) image uploaded
//	@Tags			profile-image
//	@Security		auth
//	@Accept			json
//	@Produce		json
//	@Param			image	body		param.GetPrimaryProfileImageRequest	false	"Request object"
//	@Success		200		{object}	param.GetPrimaryProfileImageResponse
//	@Router			/user/profile_img/primary [get]
func (h Handler) GetPrimaryImage(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)
	res, err := h.userSvc.GetPrimaryProfileImage(c.Request().Context(), param.GetPrimaryProfileImageRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// GetAllProfileImages godoc
//
//	@Summary		Get all profile images
//	@Description	User can get all uploaded image for it's account
//	@Tags			profile-image
//	@Security		auth
//	@Accept			json
//	@Produce		json
//	@Param			image	body		param.GetAllProfileImagesRequest	false	"Request object"
//	@Success		200		{object}	param.GetAllProfileImagesResponse
//	@Router			/user/profile_img [get]
func (h Handler) GetAllProfileImages(c echo.Context) error {
	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)
	res, err := h.userSvc.GetAllProfileImages(c.Request().Context(), param.GetAllProfileImagesRequest{}, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// DeleteProfileImage godoc
//
//	@Summary		Delete on profile image
//	@Description	User can delete a certain profile image
//	@Tags			profile-image
//	@Security		auth
//	@Accept			json
//	@Produce		json
//	@Param			image	body		param.DeleteProfileImageRequest	true	"ImageID"
//	@Success		200		{object}	param.DeleteProfileImageResponse
//	@Router			/user/profile_img [delete]
func (h Handler) DeleteProfileImage(c echo.Context) error {
	var req param.DeleteProfileImageRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if err := h.validator.ValidateDeleteProfileImageRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)
	res, err := h.userSvc.DeleteProfileImage(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}

// SetImageAsPrimary godoc
//
//	@Summary		Set an image as primary
//	@Description	User can change the primary image with this route
//	@Tags			profile-image
//	@Security		auth
//	@Accept			json
//	@Produce		json
//	@Param			image	body		param.SetImageAsPrimaryRequest	true	"imageID"
//	@Success		200		{object}	param.SetImageAsPrimaryResponse
//	@Router			/user/profile_img/primary [patch]
func (h Handler) SetImageAsPrimary(c echo.Context) error {
	var req param.SetImageAsPrimaryRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := h.validator.ValidateSetImageAsPrimaryRequest(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	claims := claims.GetClaimsFromEchoContext(c, h.authCfg)
	res, err := h.userSvc.SetImageAsPrimary(c.Request().Context(), req, claims.ID)
	if err != nil {
		code, msg := httperr.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, res)
}
