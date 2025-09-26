package usecase

import (
	"context"

	"github.com/Clausia-Ifest/clausia-server/internal/domain/dto"
	"github.com/google/uuid"
)

func (u *UUser) Auth(ctx context.Context, req dto.SignInRequest) (*dto.SignInResponse, error) {
	tx, err := u.tx.Begin(ctx, false)
	if err != nil {
		return nil, err
	}

	params := dto.GetUserParams{
		Email: req.Email,
	}

	_user, err := u.ru.Get(ctx, tx.E, params)
	if err != nil {
		return nil, err
	}

	if err := u.bcrypt.ComparePassword(_user.Password, req.Password); err != nil {
		return nil, err
	}

	accessToken, err := u.token.Encode(_user.ID, _user.FullName, _user.Role.String(), _user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.SignInResponse{
		AccessToken: accessToken,
		User:        _user.ParseDTO(),
	}, nil
}

func (u *UUser) Self(ctx context.Context, userID uuid.UUID) (*dto.User, error) {
	tx, err := u.tx.Begin(ctx, false)
	if err != nil {
		return nil, err
	}

	params := dto.GetUserParams{
		ID: userID,
	}

	_user, err := u.ru.Get(ctx, tx.E, params)
	if err != nil {
		return nil, err
	}

	return _user.ParseDTO(), nil
}
