package app

import (
	"log"
	"muBlog/internal/api/handlers"
	"muBlog/internal/api/middlewares"
	"muBlog/internal/api/schemas"
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

	router.POST("/auth/signup", middlewares.ValidateRequest[schemas.SignupRequest](authHandler.Signup))
	router.POST("/auth/login", middlewares.ValidateRequest[schemas.LoginRequest](authHandler.Login))

	router.GET("/user/:id", userHandler.GetById)

	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
