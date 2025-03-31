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

	authService := services.NewAuthService(userStore)

	userService := services.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(authService, userService)

	router := httprouter.New()

	router.POST("/user/login", userHandler.UserLogin)
	router.GET("/user/:id", userHandler.GetById)
	router.POST("/user/new", userHandler.CreateUser)

	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
