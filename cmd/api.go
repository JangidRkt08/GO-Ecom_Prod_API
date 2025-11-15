package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/jangidRkt08/go-Ecom_Prod-API/internal/adapters/postgresql/sqlc"
	"github.com/jangidRkt08/go-Ecom_Prod-API/internal/orders"
	"github.com/jangidRkt08/go-Ecom_Prod-API/internal/products"
)
type application struct{
	config config
	db *pgx.Conn
	
}
// mount -> router
func (app *application) mount() http.Handler{
	// fibre, chi, gorilla mux
	 r := chi.NewRouter()

	//  USER -> handler GET /products -> service getProducts -> repo SELECT * FROM products

  // A good base middleware stack
  r.Use(middleware.RequestID)  //important for ratelimiting
  r.Use(middleware.RealIP)	   // import for rate limiting  and analytics and trscing
  r.Use(middleware.Logger)		// important for logging
  r.Use(middleware.Recoverer)		// important for recover from Crases

  // Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

  r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("All Good for now..."))
  })

//   ---- PRODUCTS ----
  productService := products.NewService(repo.New(app.db))
  productHandler:= products.NewHandler(productService)
  r.Get("/products", productHandler.ListProducts)


//   ---- ORDERS ----
  orderService := orders.NewService(repo.New(app.db),app.db)
  ordersHandlers:= orders.NewHandler(orderService)
  r.Post("/orders",ordersHandlers.PlaceOrder)

	return r
	
}
// run ->gracefull shutdown

func (app *application) run(h http.Handler) error{
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
		WriteTimeout: time.Second*30,
		ReadTimeout: time.Second*10,
		IdleTimeout: time.Minute,
	}
	log.Printf("Listening on %s", app.config.addr)
	return srv.ListenAndServe()
}


type config struct{
	addr string
	db dbConfig
}

type dbConfig struct{
	dsn string        //domain string to connect to database
}