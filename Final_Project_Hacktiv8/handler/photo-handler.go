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

type PhotoHandlerIface interface {
	PhotoHandler(w http.ResponseWriter, r *http.Request)
}

type PhotoHandler struct {
	photoService service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) PhotoHandlerIface {
	return &PhotoHandler{photoService: photoService}
}

func (p *PhotoHandler) PhotoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userId := r.Header.Get("userId")

	switch r.Method {
	case http.MethodPost:
		p.addPhoto(w, r, userId)
	case http.MethodGet:
		p.getPhotos(w, r)
	case http.MethodPut:
		p.updatePhoto(w, r, id, userId)
	case http.MethodDelete:
		p.deletePhoto(w, r, id, userId)
	}
}

func (p *PhotoHandler) addPhoto(w http.ResponseWriter, r *http.Request, userId string) {
	if userId != "" {
		if idInt, err := strconv.Atoi(userId); err == nil {
			decoder := json.NewDecoder(r.Body)
			var photo model.PhotoRequest
			w.Header().Add("Content-Type", "application/json")
			if err := decoder.Decode(&photo); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(helper.CreateErrorResponse("error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := p.photoService.Add(ctx, &photo, idInt)

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

func (p *PhotoHandler) getPhotos(w http.ResponseWriter, r *http.Request) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	w.Header().Set("Content-Type", "application/json")
	photos, err := p.photoService.Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}
	json, _ := json.Marshal(photos)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (p *PhotoHandler) updatePhoto(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				decoder := json.NewDecoder(r.Body)
				var editPhoto model.PhotoRequest
				w.Header().Add("Content-Type", "application/json")
				if err := decoder.Decode(&editPhoto); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(helper.CreateErrorResponse("error decoding json body"))
					return
				}

				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := p.photoService.Update(ctx, idInt, userInt, &editPhoto)

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

func (p *PhotoHandler) deletePhoto(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				w.Header().Add("Content-Type", "application/json")
				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := p.photoService.Delete(ctx, idInt, userInt)

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
