package main

import (
	"fmt"
	"os"
)

// Data Structures: in the struct we define it with 2 fields which rep the title and body: Describes how page data will be stored in memory
type Page struct {
	Title string
	Body []byte
}

//this is a method named save that takes p which points to a pointer to Page: is a persitence storage that saves pages 
//it will save the pages body as a text and use the title as the filename
//octal integer literal 0600 indicates that the file should be created with read-write permissions for the current user only
func (p *Page) save() error{
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

//now let us load the pages
func loadPage(title string) (*Page, error) {
	filename :=title + ".txt"
	body,err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// intializes the Page struct
	return &Page{Title: title, Body: body}, nil
}
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}