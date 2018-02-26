package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tiaguinho/gosoap"
)

// url
var url string = "http://61.191.26.181:8888/SmsPort.asmx?WSDL"

type SoapResponse struct {
	SoapResult SoapResult
}

type SoapResult struct {
	result string
}

var r SoapResponse

// get Response instance
func getResponse() *http.Response {

	response, err := http.Get(url)

	// get error
	if err != nil {

		log.Println("get error", err.Error())
	}

	return response
}

// check status
func checkStatus(status string) {

	// status not 200
	if status != "200 OK" {
		log.Println("status error:", status)
	} else {
		// log.Println("status ok")
	}

}

// check post data
func checkPost() {
	soap, err := gosoap.SoapClient(url)
	if err != nil {
		log.Println("error in soap client init: ", err)
	}

	params := gosoap.Params{
		"Epid": "60063",
	}

	err = soap.Call("SendSmsEx", params)

	if err != nil {
		log.Println("error in soap request", err)
	}

	soap.Unmarshal(&r)

	// log.Println(r.SoapResult.result)
}

// loop and check
func checkWorker() {

	for {
		// response instance
		response := getResponse()

		// checkSatus
		checkStatus(response.Status)

		// check Post data
		checkPost()

		time.Sleep(30 * time.Second)
	}

}

func main() {
	checkWorker()
}
