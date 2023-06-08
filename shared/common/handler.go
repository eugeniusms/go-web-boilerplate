package common

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

type (
	Response struct {
		Status string      `json:"status"`
		Error  string      `json:"error"`
		Data   interface{} `json:"data"`
	}
)

func DoCommonRequest(ctx *fiber.Ctx, body interface{}) error {
	err := ctx.BodyParser(body)
	if err != nil {
		return errors.New("failed to parse body")
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		return err
	}

	return nil
}

func DoCommonSuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(Response{
		Status: "SUCCESS",
		Data:   data,
	})
}

func DoCommonErrorResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Status: "FAILED",
		Error:  err.Error(),
	})
}