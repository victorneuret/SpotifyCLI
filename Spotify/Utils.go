package MySpotify

import (
	"github.com/zmb3/spotify"
	"log"
)

func cleanArtistString(current *spotify.CurrentlyPlaying) string {
	var artists string
	for _, artist := range current.Item.Artists {
		if len(artists) == 0 {
			artists += artist.Name
		} else {
			artists += ", " + artist.Name
		}
	}
	return artists
}

func currentDevice(client *spotify.Client) spotify.PlayerDevice {
	var device spotify.PlayerDevice
	devices, err := client.PlayerDevices()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		if d.Active {
			device = d
			break
		}
	}
	return device
}
