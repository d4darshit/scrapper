package util

import (
	"fmt"
	"net/url"
	"strings"
)

func GetDomainByHostName(hostname string) (domainName string) {

	splitDomian := strings.Split(hostname, ".")

	if len(splitDomian) >= 2 {
		domainName = splitDomian[len(splitDomian)-2] + "." + splitDomian[len(splitDomian)-1]
	}

	return
}

func LinkSegregator(links []string, domain string) (internalLinks, externalLinks map[string]bool) {
	internalLinks = make(map[string]bool)
	externalLinks = make(map[string]bool)

	for _, link := range links {
		u, err := url.Parse(link)
		if err != nil {
			fmt.Println("Error parsing")
			continue
		}
		if u.Hostname() == domain || strings.HasPrefix(link, "/") || strings.HasSuffix(u.Hostname(), domain) {
			internalLinks[link] = true

		} else {

			if !strings.HasPrefix(link, ConstTelLink) && !strings.HasPrefix(link, ConstJavascriptLink) && !strings.HasPrefix(link, ConstInteralRef) {

				externalLinks[link] = true

			}
		}
	}
	return
}
