package service

import (
	"fmt"
	"log"
	"net/url"
	"scrapper/util"
	"sync"

	"github.com/gocolly/colly"
)

type WebScrapper struct {
	// Heading is a map of string and integer, it stores the count of headings
	Heading map[string]int
	// Title stores the value of title
	Title string
	// LinkGroup stores internal,external and inaccessible links
	LinkGroup Link
	// HasLoginForm is true if login page is detected, else false
	HasLoginForm bool
	// HTMLVersion used to store html version of the given url
	HTMLVersion string
	// used internally to avoid concurrent map write
	hedingMutex sync.Mutex
	// Url This is set at the time of request
	Url string
	// domain is extracted from url and used to differentiate between internal and external urls
	domain string
	// passwordInputCount This field counts the number of input fields with type as password
	passwordInputCount int
}

type Link struct {
	InternalLinks         map[string]bool
	ExternalLinks         map[string]bool
	InaccessibleLinks     map[string]bool
	inaccessibleLinkMutex sync.Mutex
}

/*
New It is the entry point of the function
Input Params: url

*/
func New(url string) *WebScrapper {
	o := WebScrapper{Url: url}
	o.LinkGroup.InternalLinks = make(map[string]bool)
	o.LinkGroup.ExternalLinks = make(map[string]bool)
	o.LinkGroup.InaccessibleLinks = make(map[string]bool)
	o.Heading = make(map[string]int)
	return &o
}

// Scarpper This function takes url as an input and sets Result Global variable
func (o *WebScrapper) Scrapper() {
	const functionName = "main.main"

	fmt.Println(functionName, "started")

	u, err := url.Parse(o.Url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(functionName, u.Hostname(), "url", o.Url)

	o.domain = util.GetDomainByHostName(u.Hostname())
	// create a anew collector
	c := colly.NewCollector(
		colly.MaxDepth(0),
		colly.Async(true),
	)
	// get all the links and segrate
	c.OnHTML("body", o.GetLinks)

	// get all the headings
	for i := 1; i <= 6; i++ {
		c.OnHTML("h"+fmt.Sprint(i), o.GetHeadings)
	}

	// check if there is a login form
	c.OnHTML("html", o.SetLoginForm)
	// get page title
	c.OnHTML("html", o.GetTitle)
	// get html version
	c.OnHTML("html", o.GetHTMLVersion)
	// in case of error - mark it as in
	c.OnError(o.GetInAccessibleLinks)

	err = c.Visit(o.Url)
	if err != nil {
		fmt.Println(err)
	}

	c.Wait()
}

/*
GetLinks This will segregate all the links
*/
func (o *WebScrapper) GetLinks(e *colly.HTMLElement) {
	// get title
	links := e.ChildAttrs("a", "href")
	o.LinkGroup.InternalLinks, o.LinkGroup.ExternalLinks = util.LinkSegregator(links, o.domain)

	for link := range o.LinkGroup.InternalLinks {
		_ = e.Request.Visit(link)
	}

}

/*
GetTitle Get Page title
*/
func (o *WebScrapper) GetTitle(e *colly.HTMLElement) {

	o.Title = e.ChildText("Title")

}

/*
GetHTMLVersion  This will fetch the HTML Version of page
*/
func (o *WebScrapper) GetHTMLVersion(e *colly.HTMLElement) {

	o.HTMLVersion = util.GetDocVersion(string(e.Response.Body))

}

/*
GetInAccessibleLinks This function is used as error callback and will store inaccessible links
*/
func (o *WebScrapper) GetInAccessibleLinks(r *colly.Response, err error) {
	o.LinkGroup.inaccessibleLinkMutex.Lock()
	o.LinkGroup.InaccessibleLinks[r.Request.URL.String()] = true
	defer o.LinkGroup.inaccessibleLinkMutex.Unlock()
}

/*
GetHeadings Increments the count for a tag when it is encountered
*/
func (o *WebScrapper) GetHeadings(e *colly.HTMLElement) {
	o.hedingMutex.Lock()
	o.Heading[e.Name]++
	o.hedingMutex.Unlock()
}

/*
SetLoginForm Check for Input type=password and will set it true if there is only one password field, to avoid signup forms

*/
func (o *WebScrapper) SetLoginForm(e *colly.HTMLElement) {
	elements := e.ChildAttrs("input", "type")

	for _, val := range elements {
		if val == "password" {
			o.passwordInputCount++
		}
	}
	if o.passwordInputCount == 1 {
		o.HasLoginForm = true
	}
}
