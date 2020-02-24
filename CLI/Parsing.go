package CLI

import (
	MySpotify "github.com/victorneuret/SpotifyCLI/Spotify"
	"log"
	"os"
)

var commands = map[string]interface{}{
	"-play":        MySpotify.SpotifyClient.PlayMusic,
	"-pause":       MySpotify.SpotifyClient.PauseMusic,
	"-next":        MySpotify.SpotifyClient.NextMusic,
	"-prev":        MySpotify.SpotifyClient.PrevMusic,
	"-vup":         MySpotify.SpotifyClient.VolumeUp,
	"-vdown":       MySpotify.SpotifyClient.VolumeDown,
	"-now":         MySpotify.SpotifyClient.CurrentlyPlaying,
	"-repeat":      MySpotify.SpotifyClient.RepeatMusic,
	"-repeatone":   MySpotify.SpotifyClient.RepeatOneMusic,
	"-repeatstop":  MySpotify.SpotifyClient.StopRepeat,
	"-shuffle":     MySpotify.SpotifyClient.ShuffleMusic,
	"-shufflestop": MySpotify.SpotifyClient.StopShuffle,
}

func Parse(client *MySpotify.SpotifyClient) {
	for _, arg := range os.Args[1:] {
		filterAndSend(arg, client)
	}
}

func filterAndSend(arg string, client *MySpotify.SpotifyClient) {
	for k, v := range commands {
		if k == arg {
			v.(func(spotifyClient MySpotify.SpotifyClient))(*client)
			return
		}
	}
	errorPrintQuit(arg)
}

func errorPrintQuit(arg string) {
	log.Fatal("unknown argument " + arg)
}
