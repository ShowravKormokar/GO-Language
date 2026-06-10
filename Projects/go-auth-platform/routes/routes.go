package routes

import (
	"encoding/json"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
)

var serverStartTime = time.Now()

func healthHandler(rw http.ResponseWriter, rq *http.Request) {
	start := time.Now()

	// Memory usage
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	cpuCount := runtime.NumCPU()
	goroutines := runtime.NumGoroutine()

	// Version info
	version := "1.0.0v"

	response := map[string]interface{}{
		"success":    true,
		"message":    "Server is healthy and alive",
		"speed_ms":   time.Since(start).Milliseconds(),
		"uptime":     time.Since(serverStartTime).String(),
		"time":       time.Now().Format(time.RFC3339),
		"memory_mb":  m.Alloc / 1024 / 1024,
		"cpu_count":  cpuCount,
		"goroutines": goroutines,
		"version":    version,
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func RegisterRouter() *mux.Router {
	r := mux.NewRouter()

	// App health check
	r.HandleFunc("/health", healthHandler).Methods("GET")

	return r
}
