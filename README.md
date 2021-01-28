# Moodify

Moodify is a Twilio SMS app that integrates the Spotify Web API to provide users a single playlist link in response to a keyword text. 

Research shows that the more choices someone has, the less satisfied they are with their choice. This app aims to reduce the number of choices a user has for a given playlist search by simply providing one link and bypassing the Spotify interface.

Use cases:
1. Text any keyword(s) to receive a link to a relevant Spotify playlist (containing the keyword)
2. Text "moodify" to receieve a random Spotify playlist link

![Screenshot](https://github.com/corinnejachelski/moodify_go/blob/master/moodify_example.png)

### Tech

Moodify is written in Golang. It utilizes the [go-spotify](https://github.com/rapito/go-spotify) library

### Installation

Moodify requires Golang to run. It also requires a Twilio number for SMS with a POST webhook configured in the Twilio Console and [ngrok tunnel server](https://www.twilio.com/blog/2013/10/test-your-webhooks-locally-with-ngrok.html). 

1. Clone the repo to your local machine. 
2. Set up a [Spotify Developer account](https://developer.spotify.com/dashboard/login) and register your app to get a Client ID and Client Secret in order to be able to use their API

Set environment variables in a file called secrets.sh. Add each variable as below:
```sh
export MY_VAR="data"
```
In the terminal:
```sh
source secrets.sh
go run *.go
```