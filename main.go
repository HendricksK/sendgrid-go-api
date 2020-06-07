// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"io"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"github.com/joho/godotenv"
)

func setupRoutes() {
	fmt.Println("set up routes called")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("/ called")
		response := sendEmail()
		io.WriteString(w, response)
		fmt.Println(response)
    })
}

func sendEmail() string {

	from := mail.NewEmail("Kurvin Development", "")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Kurvin", "")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	// client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
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
