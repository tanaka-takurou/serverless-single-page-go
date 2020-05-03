package main

import (
	"io"
	"log"
	"math"
	"bytes"
	"context"
	"strconv"
	"io/ioutil"
	"html/template"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type PageData struct {
	Page int
	PageList []int
	Title string
	ContentList []ContentData
}

type ContentData struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Image template.HTML `json:"image"`
	Categories []template.HTML `json:"categories"`
}

type ConstantData struct {
	ContentCount int `json:"contentCount"`
	ContentList []ContentData `json:"contentList"`
	CategoryList []string `json:"categoryList"`
	CategoryContentMap map[string][]int `json:"categoryContentMap"`
}

type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	baseTitle := "Example Site "
	tmp := template.New("tmp")
	var dat PageData
	q := request.QueryStringParameters
	page := q["page"]
	category := q["category"]
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
		"add": func(a, b int) int { return a + b },
		"sub": func(a, b int) int { return a - b },
		"mul": func(a, b int) int { return a * b },
		"div": func(a, b int) int { return a / b },
	}
	buf := new(bytes.Buffer)
	fw := io.Writer(buf)
	jsonString, _ := ioutil.ReadFile("constant/constant.json")
	constant := new(ConstantData)
	json.Unmarshal(jsonString, constant)
	maxContentPerPage := 10
	maxPage := int(math.Ceil(float64(constant.ContentCount)/float64(maxContentPerPage)))
	if contains(constant.CategoryList, category) {
		tmp = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/view.html", "templates/header.html", "templates/footer.html", "templates/pager.html"))
		dat.Title = baseTitle + category
		dat.Page = 1
		dat.PageList = []int{}
		dat.ContentList = getCategoryContent(constant.CategoryContentMap[category], constant.ContentList)
	} else {
		pageNumber := 1
		if len(page) > 0 {
			pageNumber, _ = strconv.Atoi(page)
		}
		if pageNumber > 1 && pageNumber <= maxPage {
			tmp = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/view.html", "templates/header.html", "templates/footer.html", "templates/pager.html"))
			dat.Title = baseTitle + page
			dat.Page = pageNumber
			dat.PageList = getPageList(pageNumber, maxPage)
			dat.ContentList = getContentRange(pageNumber, maxContentPerPage, constant.ContentCount, constant.ContentList)
		} else {
			tmp = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/view.html", "templates/header.html", "templates/footer.html", "templates/pager.html"))
			dat.Title = baseTitle
			dat.Page = 1
			dat.PageList = getPageList(1, maxPage)
			dat.ContentList = getContentRange(1, maxContentPerPage, constant.ContentCount, constant.ContentList)
		}
	}
	if e := tmp.ExecuteTemplate(fw, "base", dat); e != nil {
		log.Fatal(e)
	} else {
		log.Print("Event received.")
	}
	res := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(buf.Bytes()),
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}
	return res, nil
}

func contains(s []string, t string) bool {
	for _, v := range s {
		if t == v {
			return true
		}
	}
	return false
}

func getPageList(current int, max int) []int {
	var res []int
	pagerWidth := 2
	i := 1
	for {
		if i > max {
			break
		}
		if i != 1 && i != max {
			if i < current - pagerWidth {
				res = append(res, 0)
				i = current - pagerWidth
				continue
			}
			if i > current + pagerWidth {
				res = append(res, 0)
				i = max
				continue
			}
		}
		res = append(res, i)
		i++
	}
	return res
}

func getContentRange(page int, perPage int, maxContent int, data []ContentData) []ContentData {
	var res []ContentData
	mn := (page - 1) * perPage
	mx := int(math.Min(float64(page * perPage), float64(maxContent)))
	for i := mn; i < mx; i++ {
		res = append(res, data[i])
	}
	return res
}

func getCategoryContent(idList []int, data []ContentData) []ContentData {
	var res []ContentData
	for _, i := range idList {
		res = append(res, data[i-1])
	}
	return res
}

func main() {
	lambda.Start(HandleRequest)
}
