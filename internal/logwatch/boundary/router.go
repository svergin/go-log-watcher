package boundary

import (
	"encoding/json"
	"net/http"
)

type LogwatchHandler struct {
	mux http.ServeMux
}

type Log struct {
	Lines []string `json:"lines"`
}

func Provide() *LogwatchHandler {
	h := &LogwatchHandler{
		mux: *http.NewServeMux(),
	}
	h.mux.HandleFunc("/watch", h.handleLogFile)
	//TODO weitere Endppoints definieren
	return h
}

// ServeHTTP dispatches and executes health checks.
func (h *LogwatchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

func (h *LogwatchHandler) handleLogFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	l := Log{
		Lines: []string{"eine Zeile", "zweite Zeile", "dritte Zeile"},
	}
	json.NewEncoder(w).Encode(l)
	w.WriteHeader(http.StatusOK)
}
