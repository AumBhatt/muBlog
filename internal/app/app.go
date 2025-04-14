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

	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(authService, userService)

	router := httprouter.New()

	router.POST("/auth/signup", authHandler.Signup)
	router.POST("/auth/login", authHandler.Login)

	router.GET("/user/:id", userHandler.GetById)

	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
