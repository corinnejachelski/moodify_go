
package main

import (
	"bitbucket.org/ckvist/twilio/twiml"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"os"
)

var twilioNum string = os.Getenv("TWILIO_NUM")
var accountSid string =  os.Getenv("TWILIO_SID")
var authToken string = os.Getenv("TWILIO_AUTH_TOKEN")

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


func parseAndSendSMS(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	sender := r.FormValue("From")
	// make all inputs lowercase to avoid case sensitive duplications
	body := strings.ToLower(r.FormValue("Body"))

	fmt.Println("received this message", body, "from", sender)

	responseBody := checkMoodsData(body)
	sendSMS(w, responseBody, sender)

	

}
