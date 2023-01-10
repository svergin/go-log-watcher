package boundary

import (
	"io"
	"net/http"
	"strings"

	"github.com/svergin/go-log-watcher/internal/logwatch"
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
	// w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// l := Log{
	// 	Lines: []string{"eine Zeile", "zweite Zeile", "dritte Zeile"},
	// }
	// json.NewEncoder(w).Encode(l)
	t, err := logwatch.Start("/var/log/syslog")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer t.Cleanup()
	defer t.Stop()

	for line := range t.Lines {
		io.Copy(w, strings.NewReader(line.Text))
	}

	w.WriteHeader(http.StatusOK)
}
