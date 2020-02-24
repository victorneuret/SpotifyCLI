package main

import (
	"github.com/victorneuret/SpotifyCLI/CLI"
	"github.com/victorneuret/SpotifyCLI/Config"
	MySpotify "github.com/victorneuret/SpotifyCLI/Spotify"
	"log"
	"net/http"
)

func main() {
	Config.LoadConfiguration()

	http.HandleFunc("/callback", MySpotify.CompleteAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	client := MySpotify.InitSpotify()
	//client.PauseMusic()
	//client.PrevMusic()
	//client.StopRepeat()
	//client.CurrentlyPlaying()
	CLI.Parse(client)
	return
}
