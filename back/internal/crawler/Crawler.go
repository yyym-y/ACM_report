package crawler

import (
	"fmt"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func GetMarkdown(htmlContent string) (markdown string) {
	if markdown == "" {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		if err != nil {
			return fmt.Sprintf("Error loading HTML content: %v \n", err)
		}
		html, err := doc.Html()
		if err != nil {
			fmt.Println("Error getting HTML content:", err)
			return ""
		}
		converter := md.NewConverter("", true, nil)
		markdown, err = converter.ConvertString(html)
		if err != nil {
			fmt.Println("Error converting HTML to Markdown:", err)
			return ""
		}
	}
	markdown = strings.Replace(markdown, "$$$", "$", -1)
	markdown = strings.Replace(markdown, "\\\\", "\\", -1)
	markdown = strings.Replace(markdown, "\\_", "_", -1)
	markdown = strings.Replace(markdown, "\\[", "[", -1)
	markdown = strings.Replace(markdown, "\\]", "]", -1)
	markdown = strings.Replace(markdown, "$$", "\n$$\n", -1)
	markdown = strings.Replace(markdown, "\\dots", "\\ldots", -1)
	return
}

func RUNCrawler(p_type string, url string) (str string) {
	switch p_type {
	case "codeforces":
		str = NewCF(url).run()
	case "Atcoder":
		str = NewAT(url).run()
	}
	return
}
