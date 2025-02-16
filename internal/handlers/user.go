package handlers

import "github.com/shanto-323/Library_v1.git/internal/storage"

type UserHandler struct {
	userStore storage.UserStorage
}

func NewUserhandler(userStorage storage.UserStorage) *UserHandler {
	return &UserHandler{
		userStore: userStorage,
	}
}

type GetAllUser()
