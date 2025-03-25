package app

import (
	"log"
	"muBlog/internal/api/handlers"
	"muBlog/internal/database"
	"muBlog/internal/services"
	"muBlog/internal/stores"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func App() {
	db := database.New()
	userStore := stores.NewUserStore(db)
	userService := services.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	router := httprouter.New()

	router.GET("/users/:id", userHandler.GetById)
	router.POST("/users/new", userHandler.CreateUser)


	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
