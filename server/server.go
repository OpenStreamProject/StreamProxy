package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"fmt"
)

var port string

func Start() {
	// Check if we have a port env variable to follow
	if os.Getenv("PORT") == "" {
		port = "8000"
	} else {
		port = os.Getenv("PORT")
	}

	// Print a startup message
	fmt.Println("Starting Stream Proxy on port " + port)


	// Route
	r := mux.NewRouter()
    r.HandleFunc("/channels/{channel_id}/playlist.m3u8", channelRootHandler)
    r.HandleFunc("/channels/{channel_id}/{file}", channelProxyHandler)


	// Start the server
    http.Handle("/", r)
    http.ListenAndServe(":"+port, nil)
}