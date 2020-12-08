package domain

import (
	"context"
	"errors"
	"log"

	"github.com/deathwofl/wine-reviews/graph/model"
	"github.com/deathwofl/wine-reviews/pkg/middleware"
)

func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	confirmEmail, err := d.UserService.UserbyEmail(input.Email)
	if confirmEmail.Email != "" {
		return nil, errors.New("email already in use")
	}

	confirmUsername, err := d.UserService.UserbyUsername(input.Username)
	if confirmUsername.Username != "" {
		return nil, errors.New("username already in use")
	}

	user := &model.User{
		Username:  input.Username,
		Lastname:  input.LastName,
		FirstName: input.FirstName,
		Email:     input.Email,
	}

	err = middleware.HashPassword(input.Password, user)
	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return nil, errors.New("Error hashing password")
	}

	_, err = d.UserService.CreateUser(user)
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return nil, errors.New("Error creating user")
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return nil, errors.New("Error generating token")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := d.UserService.UserbyEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password are wrong")
	}

	err = middleware.ComparePassword(input.Password, user)
	if err != nil {
		log.Printf("Error while comparing password: %v", err)
		return nil, errors.New("email or password are wrong 2")
	}

	token, err := middleware.GenerateToken(user)
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
