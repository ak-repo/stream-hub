package handler

import (
	"github.com/ak-repo/stream-hub/api/userspb"
	"github.com/ak-repo/stream-hub/pkg/errors"
	"github.com/ak-repo/stream-hub/pkg/helper"
	"github.com/ak-repo/stream-hub/pkg/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UsersHandler struct {
	client userspb.UserServiceClient
	zlog   *zap.Logger
}

func NewUsersHandler(cli userspb.UserServiceClient, zlog *zap.Logger) *UsersHandler {
	return &UsersHandler{
		client: cli,
		zlog:   zlog,
	}
}

// FindByEmail
func (h *UsersHandler) FindByEmail(ctx *fiber.Ctx) error {
	req := &userspb.FindByEmailRequest{}

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request payload",
		})
	}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.FindByEmail(gc, req)
	if err != nil {
		h.zlog.Error("grpc FindByEmail", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, "user fetched", resp.User)
}

// FindById
func (h *UsersHandler) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	req := &userspb.FindByIdRequest{Id: id}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.FindById(gc, req)
	if err != nil {
		h.zlog.Error("grpc FindById", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, "user found", resp.User)
}

// FindAllUsers (Admin)
func (h *UsersHandler) FindAllUsers(ctx *fiber.Ctx) error {
	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.FindAllUsers(gc, &userspb.Empty{})
	if err != nil {
		h.zlog.Error("grpc FindAllUsers", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, "users list", resp.Users)
}

// Block User
func (h *UsersHandler) BlockUser(ctx *fiber.Ctx) error {
	req := &userspb.BlockUserRequest{}

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request payload",
		})
	}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.BlockUser(gc, req)
	if err != nil {
		h.zlog.Error("grpc BlockUser", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, resp.Message, nil)
}

// Unblock User
func (h *UsersHandler) UnblockUser(ctx *fiber.Ctx) error {
	req := new(userspb.BlockUserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request payload",
		})
	}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.UnblockUser(gc, req)
	if err != nil {
		h.zlog.Error("grpc UnblockUser", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, resp.Message, nil)
}

// Ban User (Admin)
func (h *UsersHandler) BanUser(ctx *fiber.Ctx) error {
	req := new(userspb.BanUserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.BanUser(gc, req)
	if err != nil {
		h.zlog.Error("grpc BanUser", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, resp.Message, nil)
}

// Unban User (Admin)
func (h *UsersHandler) UnbanUser(ctx *fiber.Ctx) error {
	req := new(userspb.BanUserRequest)

	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	gc, cancel := helper.WithGRPCTimeout()
	defer cancel()

	resp, err := h.client.UnbanUser(gc, req)
	if err != nil {
		h.zlog.Error("grpc UnbanUser", zap.Error(err))
		code, body := errors.GRPCToFiber(err)
		return ctx.Status(code).JSON(body)
	}

	return response.Success(ctx, resp.Message, nil)
}
