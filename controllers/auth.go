package controllers

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	"github.com/imkrish7/ecoship-api/application"
	authv1 "github.com/imkrish7/ecoship-api/gen/proto/auth/v1"
	repositories "github.com/imkrish7/ecoship-api/repositories/auth"
)

type AuthServer struct {
	App      *application.Application
	authRepo *repositories.AuthRepository
}

// LoginService implements authv1connect.AuthServiceHandler.
func (auth *AuthServer) LoginService(ctx context.Context, req *connect.Request[authv1.LoginServiceRequest]) (*connect.Response[authv1.LoginServiceResponse], error) {

	log.Print(req.Msg.Email)
	log.Print(req.Msg.Password)

	res := connect.NewResponse(
		&authv1.LoginServiceResponse{
			Status:   true,
			Type:     "seller",
			NextStep: "kyc",
			Auth:     "krishna",
		},
	)
	return res, nil
}

func (auth *AuthServer) SignupService(ctx context.Context, req *connect.Request[authv1.SignupServiceRequest]) (*connect.Response[authv1.SignupServiceResponse], error) {

	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Name is required"))
	}

	if req.Msg.Email == "" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Email is required"))
	}

	if req.Msg.Type == "" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Type is required"))
	}

	if req.Msg.Password == "" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Password is required"))
	}

	if req.Msg.Phone == "" {
		return nil, connect.NewError(connect.CodeFailedPrecondition, errors.New("Phone is required"))
	}

	isUserExist, err := auth.authRepo.CreateUser(ctx, req.Msg.Name, req.Msg.Email, req.Msg.Type, req.Msg.Password, req.Msg.Phone)

	log.Print(isUserExist)
	log.Print(err)
	res := connect.NewResponse(
		&authv1.SignupServiceResponse{
			Status: true,
		},
	)
	return res, nil
}

func NewAuthServer(app *application.Application, authRepo *repositories.AuthRepository) *AuthServer {
	return &AuthServer{
		App:      app,
		authRepo: authRepo,
	}
}
