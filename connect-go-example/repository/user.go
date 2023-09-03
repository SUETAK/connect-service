package repository

import (
	"context"
	"errors"
	"example/repository/model"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	FindUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, email string) (*model.User, error)
}

type UserRepositoryImpl struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	err := r.db.NewSelect().Model(user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, email string) (*model.User, error) {
	// user がいるかどうかを確認する
	cnt, err := r.db.NewSelect().Model((*model.User)(nil)).Where("email = ?", email).Count(ctx)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("user already exists")
	}

	user := &model.User{
		Email:    email,
		Password: "initPassword",
	}
	_, err = r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
