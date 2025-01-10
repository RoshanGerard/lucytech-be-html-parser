package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/context"
	"golang.org/x/net/html"
	"log"
	"lucytech/internal/models"
	"net/http"
	"net/url"
	"strings"
)

type HTMLParserService struct {
	url         string
	htmlContext string
	htmlDoc     goquery.Document
}

func HtmlParserConstruct(url string) *HTMLParserService {

	instance := &HTMLParserService{
		url:         url,
		htmlContext: "",
	}

	_, _, err := instance.getHtmlContext(url)

	if err != nil {
		return nil
	}
	return instance
}

func (service *HTMLParserService) getHtmlContext(url string) (string, *goquery.Document, error) {
	service.url = url

	// Step 1: Fetch the fully rendered HTML content using chromedp
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var htmlContent string

	err := chromedp.Run(ctx,
		chromedp.Navigate(service.url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &htmlContent),
	)

	if err != nil {
		panic(err)
	}

	// Step 2: Parse the HTML content to html.Node
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))

	if err != nil {
		panic(err)
	}

	service.htmlContext = htmlContent
	service.htmlDoc = *doc

	return service.htmlContext, doc, nil
}

func (service *HTMLParserService) getHtmlVersion() (string, error) {
	// Note: HTML version determine not based on chromedp retrieved context, because it shed metadata required
	// to detect HTML version.
	resp, err := http.Get(service.url)

	if err != nil {
		log.Fatal("Failed to retrieve HTML document ", err)
	}

	defer resp.Body.Close()

	// Determine HTML version.
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	for n := doc.Nodes[0]; n != nil; n = n.NextSibling {
		firstNode := n.FirstChild

		if firstNode.Type == html.DoctypeNode {
			return service.identifyHtmlVersion(firstNode), nil
		}
	}

	if err != nil {
		log.Fatal("Failed to determine HTML version ", err)
	}

	return "Unknown", err
}

func (service *HTMLParserService) getHtmlTitle() (string, error) {
	doc := service.htmlDoc

	title := doc.Find("title").Text()

	return title, nil
}

func (service *HTMLParserService) identifyHtmlVersion(doc *html.Node) string {
	if doc.Data == "html" && doc.Attr == nil {
		return "HTML5"
	}

	switch doc.Data {
	case "HTML 4.01 Strict":
		return "HTML 4.01 Strict"
	case "HTML 4.01 Transitional":
		return "HTML 4.01 Transitional"
	case "HTML 4.01 Frameset":
		return "HTML 4.01 Frameset"
	case "XHTML 1.0 Strict":
		return "XHTML 1.0 Strict"
	case "XHTML 1.0 Transitional":
		return "XHTML 1.0 Transitional"
	case "XHTML 1.0 Frameset":
		return "XHTML 1.0 Frameset"
	case "XHTML 1.1":
		return "XHTML 1.1"
	default:
		return "Unknown"
	}
}

func (service *HTMLParserService) countLinks() (int, int) {
	internalLinks := 0
	externalLinks := 0

	doc := service.htmlDoc

	targetUrl := service.url
	parsedTargetUrl, err := url.Parse(targetUrl)

	if err != nil {
		panic(err)
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
				parsedHref, err := url.Parse(href)
				if err == nil {
					if parsedHref.Hostname() == parsedTargetUrl.Hostname() {
						internalLinks++
					} else {
						externalLinks++
					}
				}
			} else {
				internalLinks++
			}
		}
	})

	fmt.Printf("Internal Links: %d\n", internalLinks)
	fmt.Printf("External Links: %d\n", externalLinks)

	return internalLinks, externalLinks
}

func (service *HTMLParserService) isLoginPage() bool {
	doc := service.htmlDoc

	usernameExists := doc.Find("input[type='text'], input[type='email'], input[name='username'], input[name='user'], input[name='login']").Length() > 0
	passwordExists := doc.Find("input[type='password'], input[name='password'], input[name='pass']").Length() > 0
	buttonExists := doc.Find("button[type='submit'], input[type='submit'], button[name='login']").Length() > 0

	return usernameExists && passwordExists && buttonExists
}

func ParseHTML(url string) models.HtmlParseResponseDto {
	service := HtmlParserConstruct(url)

	htmlVersion, err := service.getHtmlVersion()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("HTML Content", htmlVersion)
	}

	title, err := service.getHtmlTitle()

	internalLinks, externalLinks := service.countLinks()

	isLoginPage := service.isLoginPage()

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("HTML Content", title)
	}

	return models.HtmlParseResponseDto{
		Title:         title,
		Version:       htmlVersion,
		InternalLinks: internalLinks,
		ExternalLinks: externalLinks,
		IsLoginPage:   isLoginPage,
	}
}
