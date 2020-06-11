// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
	"io"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func setupRoutes() {
	fmt.Println("set up routes called")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		queryParams := r.URL.Query()
		song := queryParams["song"]

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, song)

		response := sendEmail("tired of being alone")
		fmt.Fprint(w, response)

	})
	
	http.HandleFunc("/send-mail", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		
		fmt.Fprint(decoder)
	})
}

func sendEmail(song string) string {

	from := mail.NewEmail("Kurvin Development", testemailuserfrom)
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Kurvin", testemailuserto)
	plainTextContent := song
	htmlContent := "<div>" + song + "</div>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridapikey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}

	return response.Body
	
}

var sendgridapikey string 
var testemailuserfrom string
var testemailuserto string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file, cannot get much done")
	}

	sendgridapikey = os.Getenv("SEND_GRID_API")
	testemailuserfrom = os.Getenv("TEST_FROM_USER")
	testemailuserto = os.Getenv("TEST_TO_USER")

	fmt.Println("main")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("end main")
}
