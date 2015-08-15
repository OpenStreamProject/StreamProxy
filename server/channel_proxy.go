// RequestURI()

package server

import (
	"net/http"
	"github.com/OpenStreamProject/StreamProxy/models"
	"strings"
	"github.com/gorilla/mux"
)

func channelProxyHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	file := vars["file"]

	channel := models.GetChannelWithId(vars["channel_id"])

	url_components := strings.Split(r.URL.RequestURI(), "?")

	query := ""

	if(len(url_components) > 1) {
		query = url_components[1]
	}

	w.Write(channel.GetFile(file, query))
}