package main

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"urlshortener/lib"
	"urlshortener/middleware"
)

type Shorten struct {
	URL string `json:"url"`
}

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func handleError(tmpl *template.Template, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		if wrapped.statusCode >= 400 {
			tmpl.ExecuteTemplate(w, "error.html", map[string]string{
				"ErrorMessage": http.StatusText(wrapped.statusCode),
				"Status":       strconv.Itoa(wrapped.statusCode),
			})
		}
	})
}

var ctx = context.Background()

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
	files := http.FileServer(http.Dir("./static"))

	r := http.NewServeMux()
	r.Handle("GET /static/", http.StripPrefix("/static", files))

	tmpl := template.Must(template.New("").ParseGlob("./templates/*"))
	redis := lib.GetRedisClient()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		key := strings.Split(r.URL.Path, "/")[1]

		val, err := redis.Get(ctx, key).Result()

		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		http.Redirect(w, r, val, http.StatusFound)
	})

	r.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		url := r.PostFormValue("url")

		key := lib.GenerateKey(redis, ctx)

		if key == nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err := redis.Set(ctx, *key, url, 0).Err()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpl.ExecuteTemplate(w, "shorten.html", Shorten{URL: *key})
	})

	r.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	s := http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(logger, handleError(tmpl, r)),
	}

	logger.Info("Server starting", slog.String("port", "8080"))

	s.ListenAndServe()
}
