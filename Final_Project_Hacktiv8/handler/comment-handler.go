package handler

import (
	"context"
	"encoding/json"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"golang-crud-sql/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type CommentHandlerIface interface {
	CommentHandler(w http.ResponseWriter, r *http.Request)
}

type CommentHandler struct {
	commentService service.CommentService
}

func NewCommentHandler(commentService service.CommentService) CommentHandlerIface {
	return &CommentHandler{commentService: commentService}
}

func (c *CommentHandler) CommentHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userId := r.Header.Get("userId")

	switch r.Method {
	case http.MethodPost:
		c.addComment(w, r, userId)
	case http.MethodGet:
		c.getComments(w, r)
	case http.MethodPut:
		c.updateComment(w, r, id, userId)
	case http.MethodDelete:
		c.deleteComment(w, r, id, userId)
	}
}

func (c *CommentHandler) addComment(w http.ResponseWriter, r *http.Request, userId string) {
	if userId != "" {
		if idInt, err := strconv.Atoi(userId); err == nil {
			decoder := json.NewDecoder(r.Body)
			var comment model.CommentRequest
			w.Header().Add("Content-Type", "application/json")
			if err := decoder.Decode(&comment); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(helper.CreateErrorResponse("error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := c.commentService.Add(ctx, &comment, idInt)

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
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("user id not found"))
	return
}

func (c *CommentHandler) getComments(w http.ResponseWriter, r *http.Request) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	w.Header().Set("Content-Type", "application/json")
	photos, err := c.commentService.Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}
	json, _ := json.Marshal(photos)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (c *CommentHandler) updateComment(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				decoder := json.NewDecoder(r.Body)
				var editComment model.EditCommentRequest
				w.Header().Add("Content-Type", "application/json")
				if err := decoder.Decode(&editComment); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(helper.CreateErrorResponse("error decoding json body"))
					return
				}

				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := c.commentService.Update(ctx, idInt, userInt, &editComment)

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
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("invalid parameter"))
	return
}

func (c *CommentHandler) deleteComment(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				w.Header().Add("Content-Type", "application/json")
				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := c.commentService.Delete(ctx, idInt, userInt)

				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write(helper.CreateErrorResponse(err.Error()))
					return
				}
				w.WriteHeader(http.StatusOK)
				w.Write(helper.CreateErrorResponse(result))
				return
			}
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write(helper.CreateErrorResponse("invalid parameter"))
	return
}
