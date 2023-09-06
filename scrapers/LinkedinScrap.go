package scrapers

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
	"github.com/matheusgb/go-jobs-crawler/structs"
)

func LinkedinScrap(recursion ...int) {
	loadMorePages := ""
	startQuery := 0
	page := 1
	tech := ""
	country := ""

	if len(recursion) > 0 {
		page = recursion[0]
		startQuery = recursion[1]
	} else {
		fmt.Println("Enter the technology you want to search for jobs:")
		fmt.Scanln(&tech)
		fmt.Println("Enter the country you want to search for jobs:")
		fmt.Scanln(&country)
	}

	jobs := []structs.LinkedinJob{}
	collyCollector := colly.NewCollector()
	linkedinUrl := fmt.Sprintf("https://linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?keywords=%s&location=%s&geoId=106057199&trk=public_jobs_jobs-search-bar_search-submit&start=%d", tech, country, startQuery)

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
				Title:   el.ChildText("h3"),
				URL:     el.ChildAttr("a", "href"),
				Company: el.ChildText("h4"),
				Time:    el.ChildText("time"),
			}
			jobs = append(jobs, job)
		})
	})

	collyCollector.Visit(linkedinUrl)
	json, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))

	fmt.Printf("This is the end of page %d. Load more job opportunities? (y/n) ", page)
	fmt.Scanln(&loadMorePages)
	if loadMorePages == "y" || loadMorePages == "Y" {
		page++
		startQuery += 25
		LinkedinScrap(page, startQuery)
	} else {
		fmt.Println("Bye!")
	}

}
