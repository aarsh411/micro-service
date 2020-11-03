package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"time/signal"

	"github.com/aarsh411/micro-service/helloworld/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gg := handlers.NewGoodbye(l)
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/Goodbye", gg)
	s := &http.Server{
		Addr:         ":9900",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	gofunc(){
		err:= s.ListenAndServe()
		if err != nil{
			l.Fatal(error)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("received terminate",sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
