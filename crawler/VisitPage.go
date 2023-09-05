package crawler

import (
	"net/http"

	"golang.org/x/net/html"
)

func VisitPage() {
	linkedinUrl := "https://www.linkedin.com/jobs/search/?currentJobId=3712528457&keywords=golang&location=Brasil"

	response, err := http.Get(linkedinUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		panic("Error on request")
	}

	document, err := html.Parse(response.Body)
	if err != nil {
		panic(err)
	}

	LinkAccessed(document)
}
