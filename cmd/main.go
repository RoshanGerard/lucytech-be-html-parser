package main

import (
	"fmt"
	"lucytech/startup"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//TIP Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined or highlighted text
	// to see how GoLand suggests fixing it.

	fmt.Println("REST Demo...!")

	startup.Server()

	//service.Test()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

//func main() {
//	// Step 1: Fetch the fully rendered HTML content using chromedp
//	ctx, cancel := chromedp.NewContext(context.Background())
//	defer cancel()
//
//	var htmlContent string
//	err := chromedp.Run(ctx,
//		chromedp.Navigate(`http://localhost:5173/test`),
//		chromedp.WaitVisible(`body`, chromedp.ByQuery),
//		chromedp.OuterHTML(`html`, &htmlContent),
//	)
//	if err != nil {
//		panic(err)
//	}
//
//	// Step 2: Parse the HTML content to html.Node
//	doc, err := html.Parse(strings.NewReader(htmlContent))
//	if err != nil {
//		panic(err)
//	}
//
//	// Step 3: Count internal and external links
//	targetUrl := "http://localhost:5173/test"
//	parsedTargetUrl, err := url.Parse(targetUrl)
//	if err != nil {
//		panic(err)
//	}
//
//	internalLinks := 0
//	externalLinks := 0
//
//	var countLinks func(*html.Node)
//	countLinks = func(n *html.Node) {
//		if n.Type == html.ElementNode && n.Data == "a" {
//			for _, attr := range n.Attr {
//				if attr.Key == "href" {
//					href := attr.Val
//					if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
//						parsedHref, err := url.Parse(href)
//						if err == nil {
//							if parsedHref.Hostname() == parsedTargetUrl.Hostname() {
//								internalLinks++
//							} else {
//								externalLinks++
//							}
//						}
//					} else {
//						internalLinks++
//					}
//				}
//			}
//		}
//		for c := n.FirstChild; c != nil; c = c.NextSibling {
//			countLinks(c)
//		}
//	}
//	countLinks(doc)
//
//	fmt.Printf("Internal Links: %d\n", internalLinks)
//	fmt.Printf("External Links: %d\n", externalLinks)
//}
