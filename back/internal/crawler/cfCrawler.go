package crawler

import (
	"fmt"
	"regexp"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func getMarkdown(htmlContent string) (markdown string) {
	if markdown == "" {
		re := regexp.MustCompile(`<span class="tex-font-style-bf">(.*?)</span>`)
		htmlContent := re.ReplaceAllString(htmlContent, `<strong>$1</strong>`)

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
		if err != nil {
			fmt.Println("Error loading HTML content:", err)
			return ""
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

func CfCrawler(c *colly.Collector, url string) (str string) {
	c.OnHTML(".problem-statement", func(problemEle *colly.HTMLElement) {
		problemEle.ForEach("div.title", func(i int, title *colly.HTMLElement) {
			if i != 0 {
				return
			}
			str = "## " + title.Text + "\n\n"
		})
	})

	c.OnHTML(".problem-statement", func(problemEle *colly.HTMLElement) {
		problemEle.ForEach("div", func(i int, divEle *colly.HTMLElement) {
			if i != 10 {
				return
			}
			ss, _ := divEle.DOM.Html()
			// fmt.Println(ss)
			problem := getMarkdown(ss)
			str += problem + "\n"
		})
		str += "\n\n"
	})

	c.OnHTML(".input-specification", func(InputEle *colly.HTMLElement) {
		ss, _ := InputEle.DOM.Html()
		input := getMarkdown(ss)
		input = strings.Replace(input, "Input", "**Input**", -1)
		str += input + "\n\n"
	})

	c.OnHTML(".output-specification", func(outputEle *colly.HTMLElement) {
		ss, _ := outputEle.DOM.Html()
		output := getMarkdown(ss)
		output = strings.Replace(output, "Output", "**Output**", -1)
		str += output + "\n\n"
	})

	c.OnHTML("div.sample-test", func(sample *colly.HTMLElement) {
		var inputs []string
		var outputs []string
		sample.ForEach(".input > pre", func(i int, preEle *colly.HTMLElement) {
			each := ""
			preEle.ForEach("div", func(i int, h *colly.HTMLElement) {
				each += h.Text + "\n"
			})
			if each == "" {
				each = preEle.Text
			}
			if len(each) >= 1 {
				each = each[:len(each)-1]
			}
			inputs = append(inputs, each)

		})
		sample.ForEach(".output > pre", func(i int, preEle *colly.HTMLElement) {
			outputs = append(outputs, preEle.Text)
		})

		for i := 0; i < len(inputs); i++ {
			str += "```\n" + inputs[i] + "\n```\n```\n" + outputs[i]
			if outputs[i][len(outputs[i])-1] != '\n' {
				str += "\n"
			}
			str += "```\n---\n"
		}
	})

	c.OnHTML(".note", func(notes *colly.HTMLElement) {
		ss, _ := notes.DOM.Html()
		note := getMarkdown(ss)
		note = strings.Replace(note, "Note", "**Note**", -1)
		str += note
	})

	c.Visit(url)
	return
}
