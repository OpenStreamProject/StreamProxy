package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/OpenStreamProject/StreamProxy/models"
)

func channelRootHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel := models.GetChannelWithId(vars["channel_id"])

	w.Write(channel.GetPlaylistFile())
}