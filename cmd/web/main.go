package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ikapfw/bnb-booking-web-app/pkg/config"
	"github.com/ikapfw/bnb-booking-web-app/pkg/handlers"
	"github.com/ikapfw/bnb-booking-web-app/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main app func
func main() {
	// change app to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil{
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.HandleFunc("/divide", Divide)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello, World!")
		
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	// _ = http.ListenAndServe(portNumber, nil)
}

// add 2 int and return sum
// func addValues(x, y int) int {
// 	return x + y
// }

// divide page handler
// func Divide(w http.ResponseWriter, r *http.Request){
// 	f, err := divideValues(100.0, 0.0)

// 	if err != nil{
// 		fmt.Fprintf(w, "Cannot divide by 0")
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f equals %f", 100.0, 0.0, f))
// }

// divide 2 float values and return the result
// func divideValues(x, y float32) (float32, error){
// 	if y <= 0{
// 		err := errors.New("cannot divide by 0")
// 		return 0, err
// 	}

// 	result := x / y

// 	return result, nil
// }