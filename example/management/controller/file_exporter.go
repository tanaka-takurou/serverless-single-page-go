package controller

import (
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

type ContentData struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Image string `json:"image"`
	Categories []string `json:"categories"`
}

type ConstantData struct {
	ContentCount int `json:"contentCount"`
	ContentList []ContentData `json:"contentList"`
	CategoryList []string `json:"categoryList"`
	CategoryContentMap map[string][]int `json:"categoryContentMap"`
}

func FileExport() error {
	contents, _ := GetContentList()
	categories, _ := GetCategoryList()
	tmp_contents := []Content{}
	tmp_categories := []Category{}
	contentList := []ContentData{}
	for _, m := range contents {
		tmp_categories, _ = GetCategoryListRelatedContent(m.Id)
		var categoryTags []string
		for _, c := range tmp_categories {
			categoryTags = append(categoryTags, "<a href='./?category="+ GetLowerHyphenSeparatedString(c.Name) + "'>" + c.Name + "</a>\n")
		}
		contentList = append(contentList, ContentData{m.Title, m.Description, m.Imagetag, categoryTags})
	}
	categoryContentMap := map[string][]int{}
	for _, c := range categories {
		tmp_contents, _ = GetContentListRelatedCategory(c.Id)
		categoryContentMap[GetLowerHyphenSeparatedString(c.Name)] = GetContentIdList(tmp_contents)
	}
	constant := ConstantData{len(contents), contentList, GetCategoryNameList(categories), categoryContentMap}
	constant_json, err := json.Marshal(constant)
	if err != nil {
		return err
	}
	constant_json_ := []byte(constant_json)
	ioutil.WriteFile("../constant/constant.json", constant_json_, os.ModePerm)
	return nil
}

func GetContentIdList(contents []Content) ([]int) {
	var contentIds []int
	for _, m := range contents {
		contentIds = append(contentIds, m.Id)
	}
	return contentIds
}

func GetCategoryNameList(categories []Category) ([]string) {
	var categoryNames []string
	for _, c := range categories {
		categoryNames = append(categoryNames, GetLowerHyphenSeparatedString(c.Name))
	}
	return categoryNames
}

func GetLowerHyphenSeparatedString(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "_", -1)
}
