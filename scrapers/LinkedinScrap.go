package scrapers

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/matheusgb/go-jobs-crawler/structs"
)

func LinkedinScrap() {
	jobs := []structs.LinkedinJob{}
	collyCollector := colly.NewCollector()
	linkedinUrl := "https://co.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=Golang&location=Brasil&locationId=&geoId=106057199&f_TPR=r2592000&start=0"

	collyCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	})

	collyCollector.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			panic("Error on request")
		}
	})

	collyCollector.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("li", func(_ int, el *colly.HTMLElement) {
			job := structs.LinkedinJob{
				Title: el.ChildText("h3"),
			}
			jobs = append(jobs, job)
		})
	})

	collyCollector.Visit(linkedinUrl)
	for _, job := range jobs {
		fmt.Println(job.Title)
	}
}
