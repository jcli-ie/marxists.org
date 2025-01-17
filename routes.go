package main

import (
	"github.com/gin-gonic/gin"
	"marxists.org/controllers"
	"marxists.org/models/repository"
)

func Routes(router *gin.Engine) {
	router.StaticFile("style.css", "./www/styles/style.css")

	authorController := controllers.AuthorController{Repo: repository.AuthorRepository{Db: DB}}
	searchController := controllers.SearchController{Repo: repository.SearchRepository{Db: DB}}
	workController := controllers.WorkController{Repo: repository.WorkRepository{Db: DB}}

	router.GET("/", controllers.IndexHandler)
	router.GET("/author/:id/*page", authorController.AuthorById)
	router.GET("/search/:query/*page", searchController.Search)
	router.GET("/work/:id/*ch", workController.Work)
}
