package auth

import (
	"go-web-boilerplate/interfaces"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/common"
	"go-web-boilerplate/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", c.register)
	auth.Post("/login", c.login)
	auth.Put("/edit", c.Shared.Middleware.AuthMiddleware, c.edit)
	auth.Post("/forgotPassword", c.forgotPassword)
	auth.Post("/resetPassword", c.resetPassword)
}

// All godoc
// @Tags Auth
// @Summary Register new user
// @Description Put all mandatory parameter
// @Param CreateUserRequest body dto.CreateUserRequest true "CreateUserRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateUserResponse
// @Failure 200 {object} dto.CreateUserResponse
// @Router /auth/register [post]
func (c *Controller) register(ctx *fiber.Ctx) error {
	var (
		req dto.CreateUserRequest
		res dto.CreateUserResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("register user with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.RegisterUser(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Login user
// @Description Put all mandatory parameter
// @Param LoginRequest body dto.LoginRequest true "LoginRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.LoginResponse
// @Failure 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (c *Controller) login(ctx *fiber.Ctx) error {
	var (
		req dto.LoginRequest
		res dto.LoginResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("login user with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.Login(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Edit user
// @Description Put all mandatory parameter
// @Param Authorization header string true "Authorization"
// @Param EditUserRequest body dto.EditUserRequest true "EditUserRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.EditUserResponse
// @Failure 200 {object} dto.EditUserResponse
// @Router /auth/edit [put]
func (c *Controller) edit(ctx *fiber.Ctx) error {
	var (
		req     dto.EditUserRequest
		res     dto.EditUserResponse
		payload dto.EditUserPayload
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	payload = dto.EditUserPayload{
		ID:       ctx.Locals("id").(float64),
		Email:    ctx.Locals("email").(string),
		Fullname: req.Fullname,
	}

	c.Shared.Logger.Infof("edit user with payload: %s", payload)

	res, err = c.Interfaces.AuthViewService.EditUser(payload)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Forgot password
// @Description Put all mandatory parameter
// @Param ForgotPasswordRequest body dto.ForgotPasswordRequest true "ForgotPasswordRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ForgotPasswordRequest
// @Router /auth/forgotPassword [post]
func (c *Controller) forgotPassword(ctx *fiber.Ctx) error {
	var (
		req dto.ForgotPasswordRequest
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("forgot password request for email: %s", req.Email)

	err = c.Interfaces.AuthViewService.ForgotPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

// All godoc
// @Tags Auth
// @Summary Reset password
// @Description Put all mandatory parameter
// @Param ResetPasswordRequest body dto.ResetPasswordRequest true "ResetPasswordRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ResetPasswordRequest
// @Router /auth/resetPassword [post]
func (c *Controller) resetPassword(ctx *fiber.Ctx) error {
	var (
		req dto.ResetPasswordRequest
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	err = c.Interfaces.AuthViewService.ResetPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}