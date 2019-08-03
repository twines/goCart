package util

import (
	"fmt"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"goCart/pkg/setting"
	"html/template"
	"math"
	"strconv"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}

type Paginate struct {
	TotalNumber int
	PerPage     int
	CurrentPage int
	TotalPage   int
	Params      map[string]interface{}
	Context     *gin.Context
	Url         string
	LinkNumber  int
}

func (p *Paginate) Paginate() template.HTML {
	var paginate string
	p.computePage()
	p.parseParams()
	page := p.Context.Query("page")
	if currentPage, err := strconv.Atoi(page); err == nil {
		if currentPage >= p.TotalPage {
			p.CurrentPage = p.TotalPage
		} else {
			p.CurrentPage = currentPage
		}
	} else {
		p.CurrentPage = 1
	}
	if p.LinkNumber == 0 {
		p.LinkNumber = 8
	}

	if p.TotalPage <= 1 {
		return template.HTML(paginate)
	}
	paginate += `<ul class="pagination">`
	if p.CurrentPage == 1 {
		paginate += `<li class="paginate_button previous disabled">`
		paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, p.CurrentPage, "上一页")
		paginate += `</li>`
	} else {
		paginate += `<li class="paginate_button previous">`
		paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, p.CurrentPage-1, "上一页")
		paginate += `</li>`
	}
	if p.TotalPage <= p.LinkNumber {
		for i := 1; i <= p.TotalPage; i++ {

			if i == p.CurrentPage {
				paginate += `<li class="paginate_button active">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			} else {
				paginate += `<li class="paginate_button">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			}
		}
	} else if p.TotalPage > p.LinkNumber && p.CurrentPage < p.LinkNumber {
		if p.CurrentPage <= p.LinkNumber/2 {
			for i := 1; i < p.CurrentPage; i++ {
				if i == p.CurrentPage {
					paginate += `<li class="paginate_button active">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				} else {
					paginate += `<li class="paginate_button">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				}
			}
			for i := p.CurrentPage; i < p.CurrentPage+p.LinkNumber-p.CurrentPage; i++ {
				if i == p.CurrentPage {
					paginate += `<li class="paginate_button active">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				} else {
					paginate += `<li class="paginate_button">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				}
			}
			paginate += "<li class='paginate_button'><a>....</a></li>"
		} else {
			if p.CurrentPage > p.LinkNumber/2+1 {
				paginate += "<li class='paginate_button'><a>....</a></li>"
			}
			for i := p.CurrentPage - p.LinkNumber/2; i < p.CurrentPage; i++ {
				if i == p.CurrentPage {
					paginate += `<li class="paginate_button active">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				} else {
					paginate += `<li class="paginate_button">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				}
			}
			for i := p.CurrentPage; i < p.CurrentPage+p.LinkNumber/2; i++ {
				if i == p.CurrentPage {
					paginate += `<li class="paginate_button active">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				} else {
					paginate += `<li class="paginate_button">`
					paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
					paginate += `</li>`
				}
			}
		}

	} else if p.TotalPage > p.LinkNumber && p.CurrentPage >= p.LinkNumber {
		paginate += "<li class='paginate_button'><a>....</a></li>"
		for i := p.CurrentPage - p.LinkNumber/2; i < p.CurrentPage; i++ {
			if i == p.CurrentPage {
				paginate += `<li class="paginate_button active">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			} else {
				paginate += `<li class="paginate_button">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			}
		}
		for i := p.CurrentPage; i < p.CurrentPage+p.LinkNumber/2; i++ {
			if i > p.TotalPage {
				break
			}
			if i == p.CurrentPage {
				paginate += `<li class="paginate_button active">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			} else {
				paginate += `<li class="paginate_button">`
				paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, i, i)
				paginate += `</li>`
			}
		}

	}

	if p.CurrentPage >= p.TotalPage {
		paginate += `<li class="paginate_button next disabled">`
		paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, p.TotalPage, "下一页")
		paginate += `</li>`
	} else {
		paginate += `<li class="paginate_button next">`
		paginate += fmt.Sprintf("<a href='%vpage=%v'>%v</a>", p.Url, p.CurrentPage+1, "下一页")
		paginate += `</li>`
	}
	paginate += "</ul>"
	return template.HTML(paginate)
}
func (p *Paginate) computePage() {
	if p.PerPage <= 0 {
		p.PerPage = 15
	}
	if p.TotalNumber <= 0 {
		p.TotalPage = 1
	} else {
		p.TotalPage = int(math.Ceil(float64(p.TotalNumber) / float64(p.PerPage)))
	}
	if p.CurrentPage <= 1 {
		p.CurrentPage = 1
	}
}
func (p *Paginate) parseParams() string {
	var params string
	if len(p.Params) > 0 {
		for k, v := range p.Params {
			if params == "" {
				params += fmt.Sprintf("%v=%v", k, v)
			} else {
				params += fmt.Sprintf("&%v=%v", k, v)
			}
		}
	}
	if p.Context.Request.URL.RawQuery == "" {
		p.Url = fmt.Sprintf("%v?%v&", p.Context.FullPath(), params)
	} else {
		p.Url = fmt.Sprintf("%v?%v&", p.Context.FullPath(), params)
	}
	return params
}
