package crawler

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type atCrawler struct {
	url string
	c   *colly.Collector
}

func NewAT(url string) *atCrawler {
	return &atCrawler{
		url: url,
		c:   colly.NewCollector(),
	}
}

func (*atCrawler) getMarkdown(html string) (str string) {
	re := regexp.MustCompile(`<var>(.*?)</var>`)
	html = re.ReplaceAllString(html, `$$$1$$`)
	str = GetMarkdown(html)
	return
}

func (p *atCrawler) run() (str string) {
	p.c.OnHTML("div.col-sm-12", func(problemEle *colly.HTMLElement) {
		problemEle.ForEach("span.h2", func(i int, title *colly.HTMLElement) {
			ss := strings.TrimSpace(title.Text)
			ss = strings.Split(ss, "\n")[0]
			str = "## " + ss + "\n\n"
		})
	})
	p.c.OnHTML("span.lang-en", func(enEle *colly.HTMLElement) {
		enEle.ForEach("div.part", func(i int, problemEle *colly.HTMLElement) {
			problemEle.ForEach("section", func(i int, divEle *colly.HTMLElement) {
				ss, _ := divEle.DOM.Html()
				ss = p.getMarkdown(ss)
				if strings.HasPrefix(ss, "### Input") {
					ss = strings.Replace(ss, "```", "---", -1)
				}
				ss = strings.Replace(ss, "\n```", "```", -1)
				str += ss
			})
			str += "\n\n"
		})
	})
	p.c.Visit(p.url)
	return
}
