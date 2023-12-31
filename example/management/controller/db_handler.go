package controller

import (
	"strings"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetContentList() ([]Content, error) {
	db, err := sqlConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	contents := []Content{}
	db.Table("content").Find(&contents)
	return contents, nil
}

func GetCategoryList() ([]Category, error) {
	db, err := sqlConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	categories := []Category{}
	db.Table("category").Find(&categories)
	return categories, nil
}

func GetContentListRelatedCategory(categoryId int) ([]Content, error) {
	db, err := sqlConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	contents := []Content{}
	db.Table("content").Select("content.id, content.title, content.description, content.imagetag").Joins("inner join content_category_relation on content.id = content_category_relation.content_id where content_category_relation.category_id = " + strconv.Itoa(categoryId)).Scan(&contents)
	return contents, nil
}

func GetCategoryListRelatedContent(contentId int) ([]Category, error) {
	db, err := sqlConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	categories := []Category{}
	db.Table("category").Select("category.id, category.name").Joins("inner join content_category_relation on category.id = content_category_relation.category_id where content_category_relation.content_id = " + strconv.Itoa(contentId)).Scan(&categories)
	return categories, nil
}

func SaveContent(title string, description string, imagetag string, category1 int, category2 int) error {
	db, err := sqlConnect()
	if err != nil {
		return err
	}
	defer db.Close()
	imagetag_ := strings.Replace(imagetag, "\"//", "\"https://", -1)
	var content = Content{Title: title, Description: description, Imagetag: imagetag_}
	db.Table("content").Create(&content)
	if category1 > 0 {
		var c1 = ContentCategoryRelation{ContentId: content.Id, CategoryId: category1}
		db.Table("content_category_relation").Create(&c1)
	}
	if category2 > 0 {
		var c2 = ContentCategoryRelation{ContentId: content.Id, CategoryId: category2}
		db.Table("content_category_relation").Create(&c2)
	}
	return nil
}

func SaveCategory(name string) error {
	db, err := sqlConnect()
	if err != nil {
		return err
	}
	defer db.Close()
	var category = Category{Name: name}
	db.Table("category").Create(&category)
	return nil
}

func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "your_mysql_user"
	PASS := "your_mysql_pass"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "example_site"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

type Content struct {
	Id          int       `gorm:"primary_key"`
	Title       string    `json:title`
	Description string    `json:description`
	Imagetag    string    `json:imagetag`
}

type Category struct {
	Id          int       `gorm:"primary_key"`
	Name        string    `json:name`
}

type ContentCategoryRelation struct {
	Id          int       `gorm:"primary_key"`
	ContentId     int       `json:content_id`
	CategoryId  int       `json:category_id`
}
