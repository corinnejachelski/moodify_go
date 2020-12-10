
package main

import (
	"bitbucket.org/ckvist/twilio/twiml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

var twilioNum string = "+19596006663"
var accountSid string =  "AC8a655345de45cf94a4dda3cc7d3a5ca3"
var authToken string = "d0b90e20a14bdd6f642a77ede3692d5a"

func sendSMS(w http.ResponseWriter, text string, recipient string) {
	resp := twiml.NewResponse()
	resp.Action(twiml.Message{
		Body: fmt.Sprintf(text),
		From: twilioNum,
		To:   recipient,
	})
	resp.Send(w)
	fmt.Println("Sending this message", text, "to", recipient)
}



func inboundSMSHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sender := r.FormValue("From")
	// make all inputs lowercase to avoid case sensitive duplications
	body := strings.ToLower(r.FormValue("Body"))

	fmt.Println("received this message", body, "from", sender)

	switch body {
	case "moodify":
		// textBody = getRandomMood()
		sendSMS(w, "Text me when you are leaving!", sender)

	default:
		responseBody := checkMoodsData(body)
		sendSMS(w, responseBody, sender)

	}

}

// func outboundSMSHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {


// 	client := twilio.NewClient(accountSid, authToken, nil)

// 	// Send a message
// 	msg, err := client.Messages.SendMessage(twilioNum, "+14104874686", "Sent via go :) ✓", nil)

// 	fmt.Fprintf(w, "Sent a message")

// 	if err == nil {
// 		fmt.Println("status >>>", msg.Status)
// 		fmt.Println("body >>>", msg.Body)
// 		fmt.Println("res >>>", r)
// 	} else {
// 		fmt.Println("error >>>", err)
// 	}
// }