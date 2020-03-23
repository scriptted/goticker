package route

import (
	"fmt"
	"net/http"
	"os"

	"github.com/scriptted/goticker/internal/config"

	"github.com/go-chi/chi"
)

// Mount routes
func Mount(r *chi.Mux, config *config.Config) error {
	fmt.Println(config.HTTP.PublicDir)
	root := config.HTTP.PublicDir
	fs := http.FileServer(http.Dir(root))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})

	return nil
}
