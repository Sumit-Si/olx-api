package main

import (
	"log"
	"net/http"
	"time"

	"github.com/sumit-si/olx-api/internal/config"
)

func main() {
	cfg :=config.MustLoad()

	// --------- BAD PRACTICE START (it is exposed to global - anyone can access it) ------------
	http.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) // 200 OK by default
		// w.Write([]byte("all ok"))	// return string
		w.Write([]byte(`{"status": "all ok"}`))
	})
	// --------- BAD PRACTICE END ------------

	// TO FIX IT: we can create a new Router(MUX in GO)
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		/*
			Order should be there otherwise you don't get what you want to achieve it
			1. Header - Content-Type
			2. WriteHeader - StatusOK
			3. Write - status: ok
		*/
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(`{"status": "ok"}`))
	})

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
