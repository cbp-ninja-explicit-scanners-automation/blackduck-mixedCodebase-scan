package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/amitramachandran/zero1/handlers"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	serverCheck := handlers.NewHealthHandler(l)
	product := handlers.NewProduct(l)

	// sm := http.NewServeMux()
	sm := mux.NewRouter()

	workingDir, _ := os.Getwd()
	//file server to handle static images
	staticImagePath := fmt.Sprintf(workingDir + "/src/images/")
	imgFS := http.FileServer(http.Dir(staticImagePath))
	sm.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imgFS))

	//file server to serve the CSS file

	staticCSSPath := fmt.Sprintf(workingDir + "/src/css/")
	cssFS := http.FileServer(http.Dir(staticCSSPath))
	sm.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssFS))

	//file server to serve the js file

	staticJSPath := fmt.Sprintf(workingDir + "/src/js/")
	jsFS := http.FileServer(http.Dir(staticJSPath))
	sm.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsFS))

	// get method routers
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/health", serverCheck.HealthCheck)
	// getRouter.HandleFunc("/", product.GetProducts)
	getRouter.HandleFunc("/", product.GetTemplProducts)
	getRouter.HandleFunc("/{id:[0-9]+}", product.GetTemplProduct)
	getRouter.HandleFunc("/about", product.GetAbout)
	getRouter.HandleFunc("/product", product.GetTemplAddProduct)

	// post method routers
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	// postRouter.HandleFunc("/product/", product.AddProduct)
	postRouter.HandleFunc("/product", product.PostTemplAddProduct)
	// postRouter.Use(product.ProductMiddleware)

	// put method routers
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/product/{id:[0-9]+}", product.UpdateProduct)
	putRouter.Use(product.ProductMiddleware)

	serve := http.Server{
		Addr:        "0.0.0.0:9090",
		Handler:     sm,
		IdleTimeout: 30 * time.Second,
	}

	l.Println("Server is up and running")
	go func() {
		err := serve.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	shutDown := make(chan os.Signal)
	signal.Notify(shutDown, os.Kill)
	signal.Notify(shutDown, os.Interrupt)

	sig := <-shutDown
	l.Printf("Gracefully shutting down %s", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	serve.Shutdown(ctx)
}
