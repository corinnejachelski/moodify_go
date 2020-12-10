package main

import (
	"github.com/rapito/go-spotify/spotify"
	"fmt"
	"github.com/buger/jsonparser"
	"math/rand"
	"time"
	"os"
	"io/ioutil"
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
	// call Spotify API if not in data
	} else {

		response := callSpotifyAPI(userInput)
		linktoSend := sendLink(response, userInput)

		return linktoSend
	}
}


func callSpotifyAPI(userInput string) []byte {
	
	// Authorize against Spotify first
	authorized, _ := spot.Authorize()
	if authorized {

		var response []byte

		if userInput == "moodify" {
			// call Spotify API to get playlists from their categories/mood endpoint (random moods)
			response = getRandomMoodPlaylist()
		} else {
			// call API with user input to search endpoint for playlists
			payload := "?q=" + userInput + "&type=playlist"
		
			response,_ = spot.Request("GET", "search/%s", nil, payload)
		}

		return response	
	} else {
		return []byte("Sorry, we can't authorize Spotify at this time.")
	}
}

func parseJson (response []byte, userInput string) int {

	totalItemsByte, _, _, _ := jsonparser.Get(response, "playlists", "total")
	totalItems := string(totalItemsByte)

	// if search does not return any items
	if totalItems == string('0') {
		return 0
	} else {
		// use jsonparser library to iterate through each object in items array 
		jsonparser.ArrayEach(response, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			// get Spotify url for each object
			linkByte, _ ,_ ,_ := jsonparser.Get(value, "external_urls", "spotify")
			link := string(linkByte)
			// add url to map for given user input
			moods[userInput] = append(moods[userInput], link)
			}, "playlists", "items")

		return 1
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

func sendLink(response []byte, userInput string) string{

		if string(response) == "Sorry, we can't authorize Spotify at this time." {

			return string(response)
		} else {
			checkResponse := parseJson (response, userInput)

			if checkResponse == 0 {

				return "Sorry, no playlists matched your search. Try another."
			} else {
				linkToSend := returnLink(userInput)

				return linkToSend
			}
		}		
}


func getRandomMoodPlaylist() []byte {

	// Open jsonFile of playlists
    response, err := os.Open("mood_playlists.json")

    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer response.Close()

    byteValue, _ := ioutil.ReadAll(response)

    return byteValue		
}

