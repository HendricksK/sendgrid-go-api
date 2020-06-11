// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func setupRoutes() {
	fmt.Println("set up routes called")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		queryParams := r.URL.Query()
		song := strings.Join(queryParams["song"], "")

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, song)

		response := sendEmailTest(song)
		fmt.Fprint(w, response)

	})
	
	http.HandleFunc("/send-mail", func(w http.ResponseWriter, r *http.Request) {

		subject := r.PostFormValue("subject")
		fromEmail := r.PostFormValue("fromemail")
		toEmail := r.PostFormValue("toemail")
		emailText := r.PostFormValue("emailtext")
		
		fmt.Println(subject)
		fmt.Println(fromEmail)
		fmt.Println(toEmail)
		fmt.Println(emailText)

		response := sendEmail(subject, fromEmail, toEmail, emailText)
		fmt.Fprint(w, response)

	})
}

func sendEmailTest(emailText string) string {

	from := mail.NewEmail("Kurvin Development", testEmailUserFrom)
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Kurvin", testEmailUserTo)
	plainTextContent := emailText
	htmlContent := "<div>" + emailText + "</div>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendGridApiKey)
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

func sendEmail(subjectEmail string, fromEmail string, toEmail string, emailText string) string {

	from := mail.NewEmail("Kurvin Development", fromEmail)
	subject := subjectEmail
	to := mail.NewEmail("Kurvin", toEmail)
	htmlContent := "<div>" + emailText + "</div>"
	message := mail.NewSingleEmail(from, subject, to, htmlContent, htmlContent)
	client := sendgrid.NewSendClient(sendGridApiKey)
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

var sendGridApiKey string 
var testEmailUserFrom string
var testEmailUserTo string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		fmt.Println("Error loading .env file, cannot get much done")
	}

	sendGridApiKey = os.Getenv("SEND_GRID_API")
	testEmailUserFrom = os.Getenv("TEST_FROM_USER")
	testEmailUserTo = os.Getenv("TEST_TO_USER")

	fmt.Println("main")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("end main")
}
