package main

import (
	"io"
	"log"
	"bytes"
	"context"
	"html/template"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	tmp := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/view.html"))
	buf := new(bytes.Buffer)
	fw := io.Writer(buf)
	dat := struct {
		Title string
	}{
		Title: "ServerlessSinglePage",
	}
	if err := tmp.ExecuteTemplate(fw, "base", dat); err != nil {
		log.Fatal(err)
	} else {
		log.Print("Event received.")
	}
	res := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(buf.Bytes()),
		Headers: map[string]string{
			"Content-Type":           "text/html",
		},
	}
	return res, nil
}

func main() {
	lambda.Start(HandleRequest)
}
