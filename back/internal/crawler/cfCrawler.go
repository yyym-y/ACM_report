package crawler

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type cfCrawler struct {
	url string
	c   *colly.Collector
}

func NewCF(url string) *cfCrawler {
	return &cfCrawler{
		url: url,
		c:   colly.NewCollector(),
	}
}

func (*cfCrawler) getMarkdown(html string) (str string) {
	re := regexp.MustCompile(`<span class="tex-font-style-bf">(.*?)</span>`)
	html = re.ReplaceAllString(html, `<strong>$1</strong>`)
	str = GetMarkdown(html)
	return
}

func (p *cfCrawler) run() (str string) {
	p.c.OnHTML(".problem-statement", func(problemEle *colly.HTMLElement) {
		problemEle.ForEach("div.title", func(i int, title *colly.HTMLElement) {
			if i != 0 {
				return
			}
			str = "## " + title.Text + "\n\n"
		})
	})

	p.c.OnHTML(".problem-statement", func(problemEle *colly.HTMLElement) {
		problemEle.ForEach("div", func(i int, divEle *colly.HTMLElement) {
			if i != 10 {
				return
			}
			ss, _ := divEle.DOM.Html()
			problem := p.getMarkdown(ss)
			str += problem + "\n"
		})
		str += "\n\n"
	})

	p.c.OnHTML(".input-specification", func(InputEle *colly.HTMLElement) {
		ss, _ := InputEle.DOM.Html()
		input := p.getMarkdown(ss)
		input = strings.Replace(input, "Input", "**Input**", -1)
		str += input + "\n\n"
	})

	p.c.OnHTML(".output-specification", func(outputEle *colly.HTMLElement) {
		ss, _ := outputEle.DOM.Html()
		output := p.getMarkdown(ss)
		output = strings.Replace(output, "Output", "**Output**", -1)
		str += output + "\n\n"
	})

	p.c.OnHTML("div.sample-test", func(sample *colly.HTMLElement) {
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

	p.c.OnHTML(".note", func(notes *colly.HTMLElement) {
		ss, _ := notes.DOM.Html()
		note := p.getMarkdown(ss)
		note = strings.Replace(note, "Note", "**Note**", -1)
		str += note
	})

	p.c.Visit(p.url)
	return
}
