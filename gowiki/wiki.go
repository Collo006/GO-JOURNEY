package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"html/template"
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

//Template Caching
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))


func renderTemplate(w http.ResponseWriter, tmpl string, p *Page ) {
	 err := templates.ExecuteTemplate(w, tmpl+ ".html",p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Using net/http to serve wiki pages
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+ title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}


func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("editHandler called!")
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err !=nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/",saveHandler)
	fmt.Println("Port is running in port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}