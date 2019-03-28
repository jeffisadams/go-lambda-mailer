package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/jeffisadams/go-lambda-mailer/src/template"
	gopher "github.com/jpoehls/gophermail"
)

// gopher "github.com/jpoehls/gophermail"

// func handleRequest(event events.SQSEvent) {
// 	// message := gopher.Message{}
// 	// message.AddTo("Jeff Adams <jeff@tenmilesquare.com>")
// 	// message.SetFrom("Jeff Adams <jeffisadams@gmail.com>")

// 	for _, record := range event.Records {
// 		fmt.Println(record.Body)
// 	}

// 	data := EmailData{
// 		Username: "YOU",
// 	}

// 	html, err := RenderTemplate("test", data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(html)

// 	// body, err := message.Bytes()
// 	// if err != nil {
// 	// 	fmt.Println("Got here and have an error")
// 	// 	return events.APIGatewayProxyResponse{}, err
// 	// }
// 	// fmt.Printf("%s", body)

// 	// return events.APIGatewayProxyResponse{
// 	// 	Body:       `{"hello":"dummy"}`,
// 	// 	StatusCode: 200,
// 	// }, nil
// }

func main() {
	// Define arbitrary map data here
	data := make(map[string]string)
	data["username"] = "ME"

	// Load your template
	var buf bytes.Buffer
	dir, _ := os.Getwd()
	fmt.Fprintf(&buf, "%s/templates/test.html", dir)

	filebuf, _ := ioutil.ReadFile(buf.String())
	templateString := string(filebuf)

	// Render the temaplate
	html, err := template.Render(templateString, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(html)

	message := gopher.Message{}

	message.AddTo("Jeff Adams <jeff@tenmilesquare.com>")
	message.SetFrom("Jeff Adams <jeffisadams@gmail.com>")

	message.Body = "You are out of luck cause your mail client can't parse html"
	message.HTMLBody = html

	msgBytes, err := message.Bytes()
	if err != nil {
		fmt.Printf("Encountered Mail build err: %s", err)
	}

	// Now prep the SES Send process
	rawMessage := ses.RawMessage{
		Data: msgBytes,
	}

	input := &ses.SendRawEmailInput{
		RawMessage: &rawMessage,
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		fmt.Println(err)
	}

	svc := ses.New(sess)

	result, err := svc.SendRawEmail(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", result)
}
