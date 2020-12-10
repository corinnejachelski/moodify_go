

package main

import (
    "fmt"
    "net/http"
    "log"
	"github.com/julienschmidt/httprouter"
)


func main() {

	router := httprouter.New()

	router.POST("/receiveSMS", parseAndSendSMS)

	// router.GET("/sendSMS", outboundSMSHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

