package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	h "github.com/AntonyIS/url_shortener/api"
	mr "github.com/AntonyIS/url_shortener/repository/mongo"
	rr "github.com/AntonyIS/url_shortener/repository/redis"
	"github.com/AntonyIS/url_shortener/shortener"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	repo := chooseRepo()
	fmt.Println("REDIS", os.Getenv("URL_DB"))
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :5000")
		errs <- http.ListenAndServe(httpPort(), r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "5000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.NewRedisRepository(redisURL)
		if err != nil {
			log.Fatal(err)
		}
		return repo

	case "mongo":
		mongoURL := os.Getenv("MONGO_URL")
		mongodb := os.Getenv("MONGO_DB")
		mongoTimeout, _ := strconv.Atoi(os.Getenv("MONGO_TIMEOUT"))
		repo, err := mr.NewMongoRepository(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
