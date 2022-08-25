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

type SocialMediaHandlerIface interface {
	SocialMediaHandler(w http.ResponseWriter, r *http.Request)
}

type SocialMediaHandler struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaHandler(socialMediaService service.SocialMediaService) SocialMediaHandlerIface {
	return &SocialMediaHandler{socialMediaService: socialMediaService}
}

func (s *SocialMediaHandler) SocialMediaHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	userId := r.Header.Get("userId")

	switch r.Method {
	case http.MethodPost:
		s.addSocialMedia(w, r, userId)
	case http.MethodGet:
		s.getSocialMedias(w, r)
	case http.MethodPut:
		s.updateSocialMedia(w, r, id, userId)
	case http.MethodDelete:
		s.deleteSocialMedia(w, r, id, userId)
	}
}

func (s *SocialMediaHandler) addSocialMedia(w http.ResponseWriter, r *http.Request, userId string) {
	if userId != "" {
		if idInt, err := strconv.Atoi(userId); err == nil {
			decoder := json.NewDecoder(r.Body)
			var socialMedia model.SocialMediaRequest
			w.Header().Add("Content-Type", "application/json")
			if err := decoder.Decode(&socialMedia); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write(helper.CreateErrorResponse("error decoding json body"))
				return
			}

			ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancelfunc()
			result, err := s.socialMediaService.Add(ctx, &socialMedia, idInt)

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

func (s *SocialMediaHandler) getSocialMedias(w http.ResponseWriter, r *http.Request) {

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	w.Header().Set("Content-Type", "application/json")
	photos, err := s.socialMediaService.Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(helper.CreateErrorResponse(err.Error()))
		return
	}
	json, _ := json.Marshal(photos)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func (s *SocialMediaHandler) updateSocialMedia(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				decoder := json.NewDecoder(r.Body)
				var editSocialMedia model.SocialMediaRequest
				w.Header().Add("Content-Type", "application/json")
				if err := decoder.Decode(&editSocialMedia); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					w.Write(helper.CreateErrorResponse("error decoding json body"))
					return
				}

				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := s.socialMediaService.Update(ctx, idInt, userInt, &editSocialMedia)

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

func (s *SocialMediaHandler) deleteSocialMedia(w http.ResponseWriter, r *http.Request, id, userId string) {
	if id != "" && userId != "" {
		if idInt, err := strconv.Atoi(id); err == nil {
			if userInt, err := strconv.Atoi(userId); err == nil {
				w.Header().Add("Content-Type", "application/json")
				ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancelfunc()
				result, err := s.socialMediaService.Delete(ctx, idInt, userInt)

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
