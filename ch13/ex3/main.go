package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"time"
)

/*
Write a small web server that returns the current time formatted using the RFC3339 format
when you send it a GET command. You can use a third-party module, if you'd like.
*/
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	r := createChiRouter(logger)
	// or
	// r := createServeMux()
	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func createServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			rw.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		t := time.Now().Format(time.RFC3339)
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(t))
		return

	})
	return mux
}

func createChiRouter(logger *slog.Logger) chi.Router {
	r := chi.NewRouter().With(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			logger.Info("incoming IP", "ip", ip)
			handler.ServeHTTP(w, r)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		var out string
		if r.Header.Get("Accept") == "application/json" {
			out = buildJSON(now)
		} else {
			out = buildText(now)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(out))
	})
	return r
}

func buildText(now time.Time) string {
	return now.Format(time.RFC3339)
}

func buildJSON(now time.Time) string {
	timeOut := struct {
		DayOfWeek  string `json:"day_of_week"`
		DayOfMonth int    `json:"day_of_month"`
		Month      string `json:"month"`
		Year       int    `json:"year"`
		Hour       int    `json:"hour"`
		Minute     int    `json:"minute"`
		Second     int    `json:"second"`
	}{
		DayOfWeek:  now.Weekday().String(),
		DayOfMonth: now.Day(),
		Month:      now.Month().String(),
		Year:       now.Year(),
		Hour:       now.Hour(),
		Minute:     now.Minute(),
		Second:     now.Second(),
	}
	out, _ := json.Marshal(timeOut)
	return string(out)
}
