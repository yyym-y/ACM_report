package crawler

import (
	"strings"

	"github.com/gocolly/colly"
)

type hduCrawler struct {
	url string
	c   *colly.Collector
}

func NewHDU(url string) *hduCrawler {
	return &hduCrawler{
		url: url,
		c:   colly.NewCollector(),
	}
}

func (*hduCrawler) getMarkdown(html string) (str string) {
	str = GetMarkdown(html)
	return
}

func (p *hduCrawler) run() (str string) {
	p.c.OnHTML("table", func(tableEle *colly.HTMLElement) {
		tableEle.ForEach("tr", func(i int, trEle *colly.HTMLElement) {
			if i != 8 {
				return
			}
			var title []string
			var content []string
			trEle.ForEach("h1", func(i int, h1 *colly.HTMLElement) {
				str += "## " + h1.Text + "\n\n"
			})
			trEle.ForEach("div.panel_title", func(i int, ttitle *colly.HTMLElement) {
				ss, _ := ttitle.DOM.Html()
				title = append(title, p.getMarkdown(ss))
			})
			trEle.ForEach("div.panel_content", func(i int, proSta *colly.HTMLElement) {
				ss, _ := proSta.DOM.Html()
				content = append(content, p.getMarkdown(ss)+"\n\n")
			})
			for i := 0; i < len(content); i++ {
				str += "**" + title[i] + "**\n\n"
				str += content[i]
			}
			str = strings.Replace(str, "```\n", "```", -1)
		})
	})

	p.c.Visit(p.url)
	return
}
