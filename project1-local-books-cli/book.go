package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 结构体字段必须大写，才能序列化
type Book struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Price    string `json:"price"`
	ImageUrl string `json:"image_url"`
}

type BookEntity interface {
	FindOneById(id int64) (Book, error)
	Create() (Book, error)
	Update(id int64) (Book, error)
	Delete(id int64)
}

type bookEntity struct {
	books []Book
	// reader *bio.Reader
	// write *bio.Writer
}

func (*bookEntity) init() {
	books := make([]Book, 0)
	file, err := os.Open("./books.json")
	defer file.Close()
	if err != nil {
		log.Fatal("读取books.json文件失败", err)
	}
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&books)
	if err != nil {
		log.Fatal("序列化books数据失败")
	}
	for _, book := range books {
		fmt.Printf("书本%v", book)
	}
}

// Create implements BookEntity.
func (*bookEntity) Create() (Book, error) {
	fmt.Println("插入一本书")
	return Book{}, nil
}

// Delete implements BookEntity.
func (*bookEntity) Delete(id int64) {
	panic("unimplemented")
}

// FindOneById implements BookEntity.
func (*bookEntity) FindOneById(id int64) (Book, error) {
	panic("unimplemented")
}

// Update implements BookEntity.
func (*bookEntity) Update(id int64) (Book, error) {
	panic("unimplemented")
}

func NewBookEntity() BookEntity {
	be := &bookEntity{}
	be.init()
	return be
}
