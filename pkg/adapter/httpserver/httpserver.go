package httpserver

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	// "github.com/rs/cors"
	"github.com/KatanaSoftwareHope/go-commons-libs/core/config"
)

func New(r chi.Router, conf *config.Config) *http.Server {

	// handler := cors.Default().Handler(r)
	// handler := cors.AllowAll().Handler(r)
	handler := r

	SRV_PORT := os.Getenv("SRV_PORT")
	if SRV_PORT != "" {
		conf.PORT = SRV_PORT
	} else {
		conf.PORT = "3000"
	}

	srv := &http.Server{
		ReadTimeout:  10 * time.Second, // Aguarda 10 segundos
		WriteTimeout: 10 * time.Second, // Responde em 10 segundos
		Addr:         ":" + conf.PORT,
		Handler:      handler,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}

	return srv
}
