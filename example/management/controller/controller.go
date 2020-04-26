package controller

import (
	"log"
	"strconv"
	"os/exec"
	"net/http"
	"html/template"
)

func HttpHandler(w http.ResponseWriter, request *http.Request){
	categoryList, err := GetCategoryList()
	if err != nil {
		panic(err.Error())
	}
	q := request.URL.Query()
	if q != nil && q["action"] != nil {
		if q["action"][0] == "deploy" {
			log.Printf("# deploy")
			err := FileExport()
			if err != nil {
				panic(err.Error())
			}
			err = exec.Command("sh", "../scripts/deploy.sh").Run()
			if err != nil {
				panic(err.Error())
			}
		} else if q["action"][0] == "savecontent" {
			log.Printf("# savecontent")
			request.ParseForm()
			category1, e := strconv.Atoi(request.Form["category1"][0])
			if e != nil {
				category1 = 0
			}
			category2, e := strconv.Atoi(request.Form["category2"][0])
			if e != nil {
				category2 = 0
			}
			err := SaveContent(request.Form["title"][0], request.Form["description"][0], request.Form["imagetag"][0], category1, category2)
			if err != nil {
				panic(err.Error())
			}
		} else if q["action"][0] == "savecategory" {
			log.Printf("# savecategory")
			request.ParseForm()
			if isNewCategory(categoryList, request.Form["categoryname"][0]) {
				err := SaveCategory(request.Form["categoryname"][0])
				if err != nil {
					panic(err.Error())
				}
				categoryList, err = GetCategoryList()
				if err != nil {
					panic(err.Error())
				}
			}
		}
	}
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	templates := template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/index.html", "templates/view.html", "templates/header.html"))
	dat := struct {
		Title string
		CategoryList []Category
	}{
		Title: "Management Page",
		CategoryList: categoryList,
	}
	err = templates.ExecuteTemplate(w, "base", dat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func isNewCategory(categoryList []Category, categoryName string) bool {
	for _, c := range categoryList {
		if c.Name == categoryName {
			return false
		}
	}
	return true
}
