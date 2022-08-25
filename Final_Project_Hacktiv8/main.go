package main

import (
	"fmt"
	"golang-crud-sql/context"
	"golang-crud-sql/handler"
	"golang-crud-sql/helper"
	"golang-crud-sql/model"
	"golang-crud-sql/repository"
	"golang-crud-sql/service"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	var cfg *model.Config

	cfg, err := helper.GetConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db := context.Connect(cfg.Database.Host, cfg.Database.Source)
	defer db.Close()

	userRepo := repository.NewUserRepo(db)
	photoRepo := repository.NewPhotoRepo(db)
	commentRepo := repository.NewCommentRepo(db)
	socialMediaRepo := repository.NewSocialMediaRepo(db)
	userService := service.NewUserSvc(userRepo)
	photoService := service.NewPhotoSvc(photoRepo)
	commentService := service.NewCommentSvc(commentRepo)
	socialMediaService := service.NewSocialMediaSvc(socialMediaRepo)
	userHandler := handler.NewUserHandler(userService)
	photoHandler := handler.NewPhotoHandler(photoService)
	commentHandler := handler.NewCommentHandler(commentService)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService)

	r := mux.NewRouter()
	r.Use(handler.IsAuthorized)
	r.HandleFunc("/users", userHandler.UserHandler)
	r.HandleFunc("/users/login", userHandler.UserHandler)
	r.HandleFunc("/users/register", userHandler.UserHandler)
	r.HandleFunc("/photos", photoHandler.PhotoHandler)
	r.HandleFunc("/photos/{id}", photoHandler.PhotoHandler)
	r.HandleFunc("/comments", commentHandler.CommentHandler)
	r.HandleFunc("/comments/{id}", commentHandler.CommentHandler)
	r.HandleFunc("/socialmedias", socialMediaHandler.SocialMediaHandler)
	r.HandleFunc("/socialmedias/{id}", socialMediaHandler.SocialMediaHandler)

	fmt.Println("Now listening on port:" + cfg.Server.Port)
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:" + cfg.Server.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
