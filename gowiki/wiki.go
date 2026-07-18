package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"html/template"
	"regexp"
	"errors"
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

//global variable to store our validation expression
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//validate path and extract page title
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil //the title is the second subexpression
}

//deifne a wrapper function thattakes a fucntion of the above type, and returns a fucntion of type http.HandlerFunc
func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	// below is the code from getTitle
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
 
// the above closure returned by makeHandler is a function that takes an http.ResponseWriter and http.Request(http.HandlerFunc).
// it extracts the ttitle from the request path and avlidates it with the validPath regexp 

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page ) {
	 err := templates.ExecuteTemplate(w, tmpl+ ".html",p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Using net/http to serve wiki pages
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	//title, err := getTitle(w, r)
	//if err !=nil {
	//	return
	//}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+ title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	renderTemplate(w, "view", p)
}


func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("editHandler called!")
	//title := r.URL.Path[len("/edit/"):]
	
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}


// this save function is used to handle the submit button
func saveHandler(w http.ResponseWriter, r *http.Request, title string){
	//title := r.URL.Path[len("/save/"):]
	
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
	http.HandleFunc("/view/", makeHandler(viewHandler))//now we wrap the handler functions ith MakeHandler in main, before they are registered with the http package.
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/",makeHandler(saveHandler))
	fmt.Println("Port is running in port 8080")
	log.Fatal(http.ListenAndServe(":8080",nil))
}