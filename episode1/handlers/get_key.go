package handlers

import (
	"fmt"
	"net/http"

	"github.com/arschles/go-in-5-minutes/episode1/storage"
	"github.com/gorilla/mux"
)

func GetKey(db storage.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, ok := mux.Vars(r)["key"]
		if !ok {
			http.Error(w, "missing key name in path", http.StatusBadRequest)
			return
		}
		val, err := db.Get(key)
		if err == storage.ErrNotFound {
			http.Error(w, "not found", http.StatusNotFound)
			return
		} else if err != nil {
			http.Error(w, fmt.Sprintf("error getting value from database: %s", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(val)
	})
}