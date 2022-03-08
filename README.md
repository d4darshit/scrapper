# **Scrapper**


## **Introduction**


This is a  web application written in golang that scrapes the web page and extracts basic information. It uses gocolly for scraping.


----------


### **Usage**

```
cd app
go run main.go
open localhost:8080?url={input url}
```
----------

### Data Extraction

- Page Title
- HTML Version  
- Headings
- Internal Links
- External Links
- Inaccessible Links
----------


### Implementation

This uses gocolly for scrapping the web page. 
- #####  Title - It fetches the title tage and renders the output
- ##### HTML Version - DocType tag is used to get the HTML Version as per [documentation](https://www.w3.org/QA/2002/04/valid-dtd-list.html)
- ##### Internal & External Links - Checks for hostname  and compares it with the links found in webpage
- ##### Login Form - Checks for password input type and counts its occurrences (to avoid signup forms).
- ##### Inaccessible Links - links that are broken or return response other than 20



