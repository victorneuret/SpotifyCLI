package MySpotify

import (
	"fmt"
	"github.com/victorneuret/SpotifyCLI/Utils"
	"log"
	"strconv"
	"strings"
)

func (s SpotifyClient) NextMusic() {
	if err := s.client.Next(); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) PrevMusic() {
	if err := s.client.Previous(); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) PauseMusic() {
	if err := s.client.Pause(); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) PlayMusic() {
	if err := s.client.Play(); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) ShuffleMusic() {
	if err := s.client.Shuffle(true); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) StopShuffle() {
	if err := s.client.Shuffle(false); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) RepeatMusic() {
	if err := s.client.Repeat("context"); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) RepeatOneMusic() {
	if err := s.client.Repeat("track"); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) StopRepeat() {
	if err := s.client.Repeat("off"); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) CurrentlyPlaying() {
	current, err := s.client.PlayerCurrentlyPlaying()
	if err != nil {
		log.Fatal(err)
	}

	if !current.Playing {
		fmt.Println("No music currently playing.")
		return
	}

	progressLine := []string{strings.Repeat("-", Utils.GetTermWidth()-2)}[0]
	pos := current.Progress * len(progressLine) / current.Item.Duration
	progressLine = "[" + progressLine[:pos] + "•" + progressLine[pos+1:] + "]"

	fmt.Println("Title     ->   ", current.Item.Name)
	fmt.Println("Artists   ->   ", cleanArtistString(current))
	fmt.Println("Album     ->   ", current.Item.Album.Name)
	fmt.Println("Volume    ->   ", strconv.Itoa(currentDevice(s.client).Volume)+"%")
	fmt.Println(progressLine)

}

func (s SpotifyClient) VolumeUp() {
	currentVolume := currentDevice(s.client).Volume
	if currentVolume > 90 {
		currentVolume = 100
	} else {
		currentVolume += 10
	}
	if err := s.client.Volume(currentVolume); err != nil {
		log.Fatal(err)
	}
}

func (s SpotifyClient) VolumeDown() {
	currentVolume := currentDevice(s.client).Volume
	if currentVolume < 10 {
		currentVolume = 0
	} else {
		currentVolume -= 10
	}
	if err := s.client.Volume(currentVolume); err != nil {
		log.Fatal(err)
	}
}
