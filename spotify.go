package main

import (
	"github.com/rapito/go-spotify/spotify"
	"fmt"
	"github.com/buger/jsonparser"
	"math/rand"
	"time"
	"os"
)

// Create a new spotify object
// pass in client ID, client secret
var spotifyClient = os.Getenv("SPOTIFY_CLIENT_ID")
var spotifySecret = os.Getenv("SPOTIFY_SECRET")
var spot = spotify.New(spotifyClient, spotifySecret)

var moods = make(map[string][]string)


func checkMoodsData(userInput string) string {

	// check if userInput is already key in moods, return link and skip API call
	if _, ok := moods[userInput]; ok {

		linkToSend := returnLink(userInput)

		return linkToSend

	} else {
		// call Spotify API if not in data
		return callSpotifyAPI(userInput)
	}
}


func callSpotifyAPI(userInput string) string {
	
	// Authorize against Spotify first
	authorized, _ := spot.Authorize()
	if authorized {

		// if userInput == "moodify" {
		// 	// call Spotify API to get playlists from their categories/mood endpoint (random moods)
		// 	response := getRandomMoodPlaylist()

		// } else {
			// call API with user input to search endpoint for playlists
		payload := "?q=" + userInput + "&type=playlist"
		
		response,_ := spot.Request("GET", "search/%s", nil, payload)
	

		totalItemsByte, _, _, _ := jsonparser.Get(response, "playlists", "total")
		totalItems := string(totalItemsByte)

		// if search does not return any items
		if totalItems == string('0') {
			return "Sorry, no playlists for this search. Try something else."
		} else {
			// use jsonparser library to iterate through each object in items array 
			jsonparser.ArrayEach(response, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				// get Spotify url for each object
				linkByte, _ ,_ ,_ := jsonparser.Get(value, "external_urls", "spotify")
				link := string(linkByte)
				// add url to map for given user input
				moods[userInput] = append(moods[userInput], link)
				}, "playlists", "items")
		}
		
		linkToSend := returnLink(userInput)

		return linkToSend
			
	} else {
		return "Sorry, we couldn't authorize Spotify at this time."
	}
}


func returnLink(userInput string) string {

	lenVals := len(moods[userInput])

	// need to seed in file so that random val is not the same every time
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(lenVals)

	// return a single link relevant to user's initial search
	return (moods[userInput][random])

}


func getRandomMoodPlaylist() string {

	authorized, authErr := spot.Authorize()
	fmt.Println(authErr)

	if authorized {

		// payload := "mood/playslists"
			
		response, err := spot.Request("GET", "browse/%s", nil, "categories", "mood", "playlists")
		fmt.Println(err)
		fmt.Println(response)
			
		return "this is a test"
	} else {
		return "Not authorized"
	}
}

