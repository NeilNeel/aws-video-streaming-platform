package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logJSON("info","hello endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
		"name_param": r.URL.Query().Get("name"),
	})
	

	w.Header().Set("Content-Type", "text/plain")
	name := r.URL.Query().Get("name")
	if name == ""{
		name = "stranger"
	}
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	logJSON("info","about endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "This is the about page.")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request){
	logJSON("info","not found endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"Page not found!")
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request){
	logJSON("info","method not allowed endpoint accessed", map[string]interface{}{
		"path": r.URL.Path,
		"method": r.Method,
	})
	
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w,"Method not allowed!")
}

func logJSON(level string,message string,data  map[string]interface{}){
	logEntry := map[string]interface{}{
		"time":time.Now().Format(time.RFC3339),
		"message":message,
		"level":level,
	}
	for k,v := range data{
		logEntry[k] = v
	}
	jsonByte, _ := json.Marshal(logEntry)
	fmt.Println(string(jsonByte))
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")

	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"status":"healthy",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
    logJSON("info", "version endpoint accessed", map[string]interface{}{
        "path": r.URL.Path,
        "method": r.Method,
    })
    
    w.Header().Set("Content-Type", "application/json")
    response := map[string]interface{}{
        "version": "1.0.0",
        "status": "running",
        "timestamp": time.Now().Format(time.RFC3339),
    }
    json.NewEncoder(w).Encode(response)
}

func main() {

	port := os.Getenv("PORT")
	if port == ""{
		port = "3000"
	}
	c := chi.NewRouter()
	c.Get("/hello", helloHandler)
	c.Get("/about", aboutHandler)
	c.Get("/health", healthHandler)
	c.Get("/version", versionHandler)

	c.NotFound(notFoundHandler)
	c.MethodNotAllowed(methodNotAllowedHandler)


	fmt.Printf("Server is running on http://localhost:%s\n", port)
	err := http.ListenAndServe(":"+port, c)
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		return
	}
}
