package main

import (
	"github.com/rapito/go-spotify/spotify"
	"fmt"
	// "encoding/json"
	"github.com/buger/jsonparser"
	"math/rand"
	"time"
)

var moods = make(map[string][]string)

// break down functions getUserMood and getRandomMood (if input is moodify)

func checkMoodsData(userInput string) string {
	if _, ok := moods[userInput]; ok {

		lenVals := len(moods[userInput])

		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(lenVals)

		return (moods[userInput][random])
	} else {
		return callSpotifyAPI(userInput)
	}
}


func callSpotifyAPI(userInput string) string {


	// Create a new spotify object
	// pass in client ID, client secret
	
	spot := spotify.New("760873230dcd4368bccc0ed1cf4bb536", "1eb159b774e54e049e403b65a4c668e4")

	// Authorize against Spotify first
	authorized, _ := spot.Authorize()
	if authorized {

			// check if userInput in moods, return link and skip API call

		// else
		// create payload for API request with userInput as search query
		payload := "?q=" + userInput + "&type=playlist"
		// need to add in error handling if search/response not valid
		response,_ := spot.Request("GET", "search/%s", nil, payload)
		// if jsonparse "playlists" "total" = 0

		totalItemsByte, _, _, _ := jsonparser.Get(response, "playlists", "total")
		totalItems := string(totalItemsByte)

		if totalItems == string('0') {
			return "Sorry, no playlists for this search. Try something else."
		} else {
			// use jsonparser library to iterate through each object in items array 
			jsonparser.ArrayEach(response, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
				// get Spotify url for each object
				linkByte, _ ,_ ,_ := jsonparser.Get(value, "external_urls", "spotify")
				link := string(linkByte)
				// add link to map for given user input
				moods[userInput] = append(moods[userInput], link)
				}, "playlists", "items")
		}
		
		lenVals := len(moods[userInput])

		// need to Seed program to not get same random num every time
		rand.Seed(time.Now().UnixNano())
		random := rand.Intn(lenVals)
		fmt.Println(random)

		return (moods[userInput][random])
			
	} else {
		return "Sorry, we couldn't authorize Spotify"
	}
}


