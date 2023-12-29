package main

import (
	"context"
	"log"
	"main/zero1/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// hello := handlers.NewHelloHandler(l)
	product := handlers.NewProduct(l)

	// sm := http.NewServeMux()
	sm := mux.NewRouter()

	//file server to handle static images
	staticImagePath := "/Users/ext.amit.r/Documents/projectZero/zero1/src/images/"
	imgFS := http.FileServer(http.Dir(staticImagePath))
	sm.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imgFS))

	//file server to serve the CSS file
	staticCSSPath := "/Users/ext.amit.r/Documents/projectZero/zero1/src/css/"
	cssFS := http.FileServer(http.Dir(staticCSSPath))
	sm.PathPrefix("/css/").Handler(http.StripPrefix("/css/", cssFS))

	//file server to serve the js file
	staticJSPath := "/Users/ext.amit.r/Documents/projectZero/zero1/src/js/"
	jsFS := http.FileServer(http.Dir(staticJSPath))
	sm.PathPrefix("/js/").Handler(http.StripPrefix("/js/", jsFS))

	// get method routers
	getRouter := sm.Methods(http.MethodGet).Subrouter()
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

	// sm.HandleFunc("/", product.ServeHTTP)
	// sm.HandleFunc("/describe", hello.DescribeFunc)

	serve := http.Server{
		Addr:        "localhost:9090",
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
