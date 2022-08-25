package handler

import (
	"context"
	"encoding/json"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"golang-crud-sql/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type UserHandlerIface interface {
	UserHandler(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandlerIface {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) UserHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if strings.HasSuffix(r.URL.Path, "users/register") {
		if r.Method == http.MethodPost {
			u.register(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(helper.CreateErrorResponse("invalid method"))
			return
		}

	} else if strings.HasSuffix(r.URL.Path, "users/login") {
		if r.Method == http.MethodPost {
			u.login(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write(helper.CreateErrorResponse("invalid method"))
			return
		}
	} else {
		switch r.Method {
		case http.MethodPut:
			u.updateUser(w, r, userId)
		case http.MethodDelete:
			u.deleteUser(w, r, userId)
		}
	}
}

func (u *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var loginUser model.UserLogin
	w.Header().Set("Content-Type", "application/json")

	if err := decoder.Decode(&loginUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.CreateErrorResponse("error decoding json body"))
		return
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	user, err := u.userService.Login(ctx, loginUser)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}

	validToken, err := GenerateJWT(user.Id, user.Username, user.Age)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse("failed to generate token"))
		return
	}

	var token model.Token
	token.Token = validToken
	json, _ := json.Marshal(token)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (u *UserHandler) getUsers(w http.ResponseWriter, _ *http.Request) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	w.Header().Set("Content-Type", "application/json")
	users, err := u.userService.GetUsers(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}
	json, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (u *UserHandler) getUserById(w http.ResponseWriter, _ *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			w.Header().Add("Content-Type", "application/json")
			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			user, err := u.userService.GetUserById(ctx, idInt)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(helper.CreateErrorResponse(err.Error()))
				return
			}

			if user.Id != 0 {
				jsonData, _ := json.Marshal(user)
				w.WriteHeader(http.StatusOK)
				w.Write(jsonData)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			w.Write(helper.CreateErrorResponse("user not found"))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("invalid parameter"))
	return
}

func (u *UserHandler) register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user model.UserRegistration
	w.Header().Add("Content-Type", "application/json")
	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(helper.CreateErrorResponse("error decoding json body"))
		return
	}

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	result, err := u.userService.Register(ctx, &user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}
	jsonData, _ := json.Marshal(&result)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
	return
}

func (u *UserHandler) updateUser(w http.ResponseWriter, r *http.Request, userId string) {
	if userId != "" {
		if idInt, err := strconv.Atoi(userId); err == nil {
			decoder := json.NewDecoder(r.Body)
			var editUser model.UserRequest
			w.Header().Add("Content-Type", "application/json")
			if err := decoder.Decode(&editUser); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(helper.CreateErrorResponse("error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := u.userService.UpdateUser(ctx, idInt, &editUser)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(helper.CreateErrorResponse(err.Error()))
				return
			}
			jsonData, _ := json.Marshal(&result)
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("user id not found"))
	return
}

func (u *UserHandler) deleteUser(w http.ResponseWriter, _ *http.Request, id string) {
	if id != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			w.Header().Add("Content-Type", "application/json")
			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := u.userService.DeleteUser(ctx, idInt)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(helper.CreateErrorResponse(err.Error()))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(result))
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("user id not found"))
	return
}
