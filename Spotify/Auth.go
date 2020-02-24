package MySpotify

import (
	"fmt"
	"github.com/victorneuret/SpotifyCLI/Config"
	"github.com/victorneuret/SpotifyCLI/Utils"
	"github.com/zmb3/spotify"
	"log"
	"net/http"
)

var (
	auth  spotify.Authenticator
	ch    = make(chan *spotify.Client)
	state = "gopher"
)

func InitSpotify() *SpotifyClient {
	//if client := getAuth(); client != nil {
	//	return client
	//}

	auth = spotify.NewAuthenticator(Config.GetConfig().RedirectURI, spotify.ScopeUserReadCurrentlyPlaying, spotify.ScopeUserReadPlaybackState, spotify.ScopeUserModifyPlaybackState)
	auth.SetAuthInfo(Config.GetConfig().ClientID, Config.GetConfig().ClientSecret)

	authUrl := auth.AuthURL(state)
	//fmt.Println("Please log in to Spotify by visiting the following page in your browser:", authUrl+"&redirect_uri="+url.QueryEscape(Config.GetConfig().RedirectURI))
	Utils.OpenBrowser(authUrl)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	_, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("You are logged in as:", user.ID)
	return &SpotifyClient{client: client}
}

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := auth.NewClient(tok)
	//Utils.WriteFile("token", tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- &client
}

//func getAuth() *SpotifyClient {
//
//	client := auth.NewClient(token)
//	user, err := client.CurrentUser()
//	if err != nil {
//		log.Println(err)
//		return nil
//	}
//	fmt.Println("You are logged in as:", user.ID)
//	return &SpotifyClient{client: &client}
//}
