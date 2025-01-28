package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// UIComponent defines the structure of a UI component
type UIComponent struct {
	Type       string                 `json:"type"`       // Component type (e.g., "button", "text", "image")
	Properties map[string]interface{} `json:"properties"` // Component properties (e.g., "text": "Click me")
	Children   []UIComponent          `json:"children"`   // Nested components
}

// HomeHandler generates a UI layout dynamically
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Define the UI structure
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
					"onClick": "/api/click", // API to handle button click
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

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ui); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// ClickHandler handles button click action
func ClickHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Button clicked!",
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Set up routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/api/click", ClickHandler)

	// Start the server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
