## Introduction

LucyTech's HTML Parser Server Practical Test Backend

This application is designed to efficiently crawl through the provided URL and extract essential information from the web page. The server is capable of parsing HTML content and provides the following information:

- **Internal Links**: Extracts and lists all internal links present on the page.
- **External Links**: Extracts and lists all external links leading to different domains.
- **HTML Version**: Identifies and displays the HTML version used by the web page.
- **Page Title**: Retrieves and shows the title of the web page.
- **Login Page**: Detect login page.

With these capabilities, our HTML Parser Server aims to simplify web scraping and data extraction tasks, making it a valuable tool for developers and analysts.

## How to Run
Run the main.go using ``go`` cli command.

```shell
cd ~/<project>
go run ./cmd/main.go
```

## Dependencies

This project uses the following dependencies:

- **Go** version 1.23.4

- **Direct Dependencies**:
  - [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) v1.10.1: A little like jQuery, but for Go.
  - [github.com/chromedp/chromedp](https://github.com/chromedp/chromedp) v0.11.2: Package chromedp is a high-level Chrome DevTools Protocol.
  - [github.com/gorilla/mux](https://github.com/gorilla/mux) v1.8.1: A powerful URL router and dispatcher for matching incoming requests to their respective handler.
  - [golang.org/x/net](https://golang.org/x/net) v0.33.0: Supplementary Go networking libraries.

- **Indirect Dependencies**:
  - [github.com/andybalholm/cascadia](https://github.com/andybalholm/cascadia) v1.3.3: CSS selector library.
  - [github.com/chromedp/cdproto](https://github.com/chromedp/cdproto) v0.0.0-20241022234722-4d5d5faf59fb: Package cdproto contains the generated DevTools Protocol domains.
  - [github.com/chromedp/sysutil](https://github.com/chromedp/sysutil) v1.1.0: System utilities for Chromedp.
  - [github.com/gobwas/httphead](https://github.com/gobwas/httphead) v0.1.0: Fast HTTP headers parsing library.
  - [github.com/gobwas/pool](https://github.com/gobwas/pool) v0.2.1: Resource pools for efficient HTTP processing.
  - [github.com/gobwas/ws](https://github.com/gobwas/ws) v1.4.0: A fast WebSocket library.
  - [github.com/josharian/intern](https://github.com/josharian/intern) v1.0.0: Efficient string interning.
  - [github.com/mailru/easyjson](https://github.com/mailru/easyjson) v0.7.7: Fast JSON serialization for golang.
  - [golang.org/x/sys](https://github.com/golang/sys) v0.28.0: Low-level interaction with the operating system.

Make sure to run `go mod tidy` to keep your `go.mod` file updated with only the necessary dependencies.

## Author

- **Developer**: Roshan G. Bolonna
- **Email**: srgbolonna@gmail.com
- **GitHub**: [GitHub Profile](https://github.com/RoshanGerard)
- **Website**: [LinkedIn](https://www.linkedin.com/in/srgbolonna)

If you have any questions or feedback, feel free to reach out.
