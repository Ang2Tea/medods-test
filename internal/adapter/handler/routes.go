package handler

import (
	"context"
	"errors"

	"Ang2Tea/medods-test/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type fiberRoutes struct {
	auth usecase.IAuth
}

func NewFiberRoutes(auth usecase.IAuth) *fiberRoutes {
	return &fiberRoutes{
		auth: auth,
	}
}

func (f *fiberRoutes) RegisterRoutes(app *fiber.App) {
	app.Post("/auth", f.Auth)
	app.Post("/refresh", f.Refresh)
}

func (f *fiberRoutes) Auth(c *fiber.Ctx) error {
	ctx := context.Background()

	strUserID := c.Get("user-id")
	if strUserID == "" {
		return fiberError(c, fiber.StatusBadRequest, errors.New("user id is required"))
	}

	userID, err := uuid.Parse(strUserID)
	if err != nil {
		return fiberError(c, fiber.StatusBadRequest, err)
	}

	request := usecase.UserRequest{
		UserID:    userID,
		IPAddress: c.IP(),
	}

	token, err := f.auth.Register(ctx, request)
	if err != nil {
		return fiberError(c, fiber.StatusInternalServerError, err)
	}

	result := toFiber(*token)

	return c.Status(fiber.StatusCreated).JSON(result)
}

func (f *fiberRoutes) Refresh(c *fiber.Ctx) error {
	ctx := context.Background()

	refreshToken := c.Get("refresh-token")
	if refreshToken == "" {
		return fiberError(c, fiber.StatusBadRequest, errors.New("refresh token is required"))
	}

	request := usecase.RefreshRequest{
		RefreshToken: refreshToken,
		IPAddress:    c.IP(),
	}

	token, err := f.auth.Refresh(ctx, request)
	if err != nil {
		return fiberError(c, fiber.StatusInternalServerError, err)
	}

	result := toFiber(*token)

	return c.Status(fiber.StatusCreated).JSON(result)
}
