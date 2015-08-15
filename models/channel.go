package models

import (
	"github.com/OpenStreamProject/StreamProxy/proxy"
	"strings"
	"time"
)

type Channel struct {
	Id					string
	Url 				string
	StreamUrl 			string
	StreamUrlCreatedAt	int64
}

// Retrives contents of the Playlist file for the channel
func (c *Channel) GetPlaylistFile() []byte {
	return proxy.GetBytesForURL(c.playlistURL())
}

// Retrieves the contents of another file relative to the location
// of the playlist file
func (c *Channel) GetFile(file string, query string) []byte {
	url := c.playlistURL()

	path_components := strings.Split(url, "/")
	file_name := path_components[len(path_components) - 1]

	url = strings.Split(url, file_name)[0]

	url = url + file + "?" + query

	return proxy.GetBytesForURL(url)
}

// Gets the URL for the playlist and extracts it from the base url if needed
func (c *Channel) playlistURL() string {

	if c.StreamUrl == "" || c.checkExpired() {
		c.StreamUrl = c.extractPlaylistURL()
		c.StreamUrlCreatedAt = time.Now().Unix()
	}
	
	return c.StreamUrl
}

// Most sites have expiration on playlist urls. This forces the server to get a new playlist
// url every 30 minutes.
func (c *Channel) checkExpired() bool {
	ttl := 30 // minutes
	secs := time.Now().Unix()

	return secs - c.StreamUrlCreatedAt > int64(ttl * 60)
}

// Loads the base url and extracts the playlist url
func (c *Channel) extractPlaylistURL() string {
	base_url := strings.TrimSpace(c.Url)

	data := proxy.GetStringForURL(base_url)

	segments := strings.Split(data, "\"")

	url := ""

	for _, segment := range(segments) {
		segments := strings.Split(segment, "'")
		
		for _, segment := range(segments) {

			if(len(segment) > 4 &&
				string(segment[0:4]) == "http" &&
				strings.Contains(segment, "m3u8")) {
				url = strings.Replace(segment, "\r","",0)
				url = strings.Replace(url, "\n","",0)
				break;
			}
		}

		if url != "" {
			break;
		}
	}

	return url
}