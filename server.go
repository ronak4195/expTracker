package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/rs/cors"
)

type UIComponent struct {
    Type       string                 `json:"type"`
    Properties map[string]interface{} `json:"properties"`
    Children   []UIComponent          `json:"children"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
    ui := UIComponent{
        Type: "container",
        Properties: map[string]interface{}{
            "direction": "vertical",
            "padding":   16,
        },
        Children: []UIComponent{
            {
                Type: "text",
                Properties: map[string]interface{}{
                    "text":  "Welcome to Server-Driven UI",
                    "style": "header",
                },
            },
            {
                Type: "button",
                Properties: map[string]interface{}{
                    "text":    "Click Me",
                    "onClick": "/api/click",
                },
            },
            {
                Type: "image",
                Properties: map[string]interface{}{
                    "src":    "https://example.com/image.jpg",
                    "alt":    "Example Image",
                    "height": 200,
                },
            },
        },
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(ui)
}

func ClickHandler(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{
        "message": "Button clicked!",
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", HomeHandler)
    mux.HandleFunc("/api/click", ClickHandler)

    handler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    }).Handler(mux)

    log.Println("Server is running on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}
