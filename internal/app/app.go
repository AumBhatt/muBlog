package app

import (
	"log"
	"muBlog/internal/api/handlers"
	"muBlog/internal/database"
	"muBlog/internal/services"
	"muBlog/internal/stores"
	"net/http"
)

func App() {
	db := database.New()
	userStore := stores.NewUserStore(db)
	userService := services.NewUserService(userStore)
	userHandler := handlers.NewUserHandler(userService)

	http.HandleFunc("GET /users/*", userHandler.GetById)

	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
