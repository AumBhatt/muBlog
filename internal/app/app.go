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
	db := database.New(stores.InitStores)
	userStore := stores.NewUserStore(db)
	postStore := stores.NewPostStore(db)

	authService := services.NewAuthService(userStore)
	userService := services.NewUserService(userStore)
	postService := services.NewPostService(postStore, userStore)

	authHandler := handlers.NewAuthHandler(authService, userService)
	userHandler := handlers.NewUserHandler(authService, userService)
	postHandler := handlers.NewPostHandler(postService)

	router := httprouter.New()

	router.POST("/auth/signup", middlewares.ValidateRequest[schemas.SignupRequest](authHandler.Signup))
	router.POST("/auth/login", middlewares.ValidateRequest[schemas.LoginRequest](authHandler.Login))

	router.GET("/user/:id", userHandler.GetById)

	// router.POST("/post/create", middlewares.ValidateRequest[schemas.CreatePostRequest](postHandler.Create))
	// router.POST("/post/react", middlewares.ValidateRequest[schemas.AddReactionRequest](postHandler.React))

	router.POST("/post/create", postHandler.Create)
	router.GET("/post/reactions/detailed/:postId", postHandler.GetReactionsByPostId)
	router.GET("/post/reactions/count/:postId", postHandler.GetReactionsCountByPostId)
	router.POST("/post/react", postHandler.React)

	log.Println("App running @ http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
