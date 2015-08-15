package models

import (
	"github.com/OpenStreamProject/StreamProxy/proxy"
	"strings"
)

var channelCache map[string]*Channel

// Searches the cache for the requested channel
// If the requested channel cannot be found the cache is reloaded
func GetChannelWithId(id string) *Channel {

	if(channelCache == nil) {
		channelCache = map[string]*Channel{}
	}

	if(channelCache[id] == nil) {
		reloadChannels()
	}

	return channelCache[id]
}

// Reloads & Loads the channel cache. Pulls the CSV from the github repository
func reloadChannels() {
	csv := proxy.GetStringForURL("https://openstreamproject.github.io/StreamDatabase/channels.csv")

	lines := strings.Split(csv, "\n")

	for _, line := range lines {
		if len(line) != 0 && string(line[0]) != "#" {
			addOrUpdateChannelWithData(strings.Split(line, ","))
		}
	}
}

// An access method that will update or create a new channel object with data
func addOrUpdateChannelWithData(data []string) {
	if(len(data)>3) {
		id := data[0]
		url := data[3]

		channel := channelCache[id]

		if channel == nil {
			channel = &(Channel{})
		} else if channel.Url != url {
			channel.StreamUrl = ""
			channel.StreamUrlCreatedAt = 0
		}
	
		channel.Url = url
		channel.Id = id


		channelCache[string(id)] = channel;
	}
}
