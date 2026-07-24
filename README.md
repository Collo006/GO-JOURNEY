# GO-API

## GET ALBUMS REQUEST
## Explanation of func getAlbums()
- gin.Context: it carries request details, validates and serializes JSON and more
- Context.IndentedJSON: serializes the struct into JSON and add it to the response
- the first argument(http.StatusOK): status code you want to send to the client
- the second argument(albums): fetches the data from variable albums

## handler fucntion
- (gin.Default)Initialize a Gin router using Default
- (getAlbums) only passing the name not the function

## POST ALBUMS REQUEST
- Context.BindJSON to bind the request body to newAlbum
- Append the album struct initialized from the JSON to the albums slice
- Append a 201 status code to the response
- With Gin, you can associate a handler with an HTTP method-and-path combination. In this way, you can separately route requests sent to a single path based on the method the client is using.

## HANDLER TO RETURN A SPECIFIC ITEM
- Add logic to retrieve requested album
- Map the path to the logic

# GO WEB APPLICATIONS
## Introducing the net/http package

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "HI there, I love %s!", r.URL.Path[1:])
}

func main(){
    http.HandleFunc("/",handler)
    log.Fatal(http.ListenAndServe(":8080",nil))
}
- the main function calls to http.HandleFunc, which tells net/http to handle all requests to the root
- it then calls http.ListenAndServe, specifying that it should listen on port 8080
- ListenAndServe only returns an error(when an unexpected error occurs). in order to log that error we wrap it in log.Fatal
- w, http.ResponseWRiter it assembles the HTTP server's response: by writing to it, send data to the HTTP client
- r *http.Request is a data structure that represents the client HTTP request. r.URL.Path is the oath component of the request URL. 
- [1:] means create a sub-slice of Path from the 1st character to the end. this drops the leading "/" from the path name

## go build -o wiki
- use the above if your port is too old

## Handling non-existent pages

- The http.Redirect function adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP response.

## Error Handling
http.Error func sends a speficied HTTP repsonse code (Internal Server Error) and an error message
- we re-use the variable err:  The same err variable is reused for each operation because each error is checked and handled before moving on. It's one of the most common idioms you'll encounter in Go programs.

## Saving Pages
- The page title (provided in the URL) and the form's only field, Body, are stored in a new Page. The save() method is then called to write the data to a file, and the client is redirected to the /view/ page.

- The value returned by FormValue is of type string. We must convert that value to []byte before it will fit into the Page struct. We use []byte(body) to perform the conversion.

## Template Caching
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
    t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
- var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
- Above is a better approach beacuse we call ParseFiles once at program initialization the latter would call everytime we need it
- func template.Must is a convenience wrapper that panics when passed a non-nil error value and otherwise returns the *TEmplate unaltered
- A panic is Go's way of saying:
- "Something has gone so seriously wrong that this program cannot continue."

## panic
- panic(err)
- Prints a stack trace.
- Shows exactly where the problem occurred.
- Can be recovered using recover() (advanced topic).
## log.Fatal
- log.Fatal(err)
- Prints the error message.
- Calls os.Exit(1).
- Does not print a stack trace.
- Cannot be recovered.

## Validation
- regexp.MustCompile will parse and compile the regular expression and return a regexp. 
- regexp.MustCompile is different from Compile in that it will panic if the expression compilation fails, while Compile returns an error as a second parameter.
- if the title is valid, it will be returned along with a nil error value. If the title is invalid, the function will write a "404 Not Found" error to the HTTP connection and return an error to the handler.

## Introducing Function Literals and 
- func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandleFunc {
     // here we will extract the page title from the request,
     // and call the provided handler "fn"
}
- the returned function is called a closure beacuse it encloses values defined outside of it. variabe fn is enclosed by the closure. "fn" will be one of our save, edit or view handlers.
- the above closure returned by makeHandler is a function that takes an http.ResponseWriter and http.Request(http.HandlerFunc).
- it extracts the title from the request path and validates it with the validPath regexp .
- now we reomve the calls to getTitle from the handler fucntions since it is being done in the makeHandler

## REGEX 
- "\\[[a-zA-Z0-9]+\\]"
- Adding \ tells regex:

"No, I mean the actual character '['."

Likewise:

\]

means

"the actual ]"

# ACCESSING RELATIONAL DATABASE
### How to create MariaDB
- sudo apt update
- sudo apt install mariadb-server -y
- sudo service mariadb start
- mysql -u root

### Restart MariaDB
- sudo service mysql stop
- sudo mysqld-safe --skip-grant-tables --skip-networking &

## How to check if table is working
- run the script to the path source /path/to/create-tables.sql
- select * from album;

## Find and Import a database driver
- In your browser, visit the SQLDrivers wiki page to identify a driver you can use.
- Use the list on the page to identify the driver you’ll use. For accessing MySQL in this tutorial, you’ll use Go-MySQL-Driver.
- Note the package name for the driver – here, github.com/go-sql-driver/mysql.

###  Get a databse handle and connect
- Declare a db variable of type *sql.DB. This is your database handle.
- Making db a global variable simplifies this example. In production, you’d avoid the global variable, such as by passing the variable to functions that need it or by wrapping it in a struct.
- Use the MySQL driver’s Config – and the type’s FormatDSN -– to collect connection  properties and format them into a DSN for a connection string.

- The Config struct makes for code that’s easier to read than a connection string would be.

-Call sql.Open to initialize the db variable, passing the return value of FormatDSN.

- Check for an error from sql.Open. It could fail if, for example, your database connection specifics weren’t well-formed.

- To simplify the code, you’re calling log.Fatal to end execution and print the error to the console. In production code, you’ll want to handle errors in a more graceful way.

- Call DB.Ping to confirm that connecting to the database works. At run time, sql.Open might not immediately connect, depending on the driver. You’re using Ping here to confirm that the database/sql package can connect when it needs to.

- Check for an error from Ping, in case the connection failed.

- Print a message if Ping connects successfully.

### how to create user password and user name
- sudo mysql -u root

- sql-- 1. Create the user for local connections
CREATE USER 'username'@'localhost' IDENTIFIED BY 'password';

-- 2. Give the new user access to your database (e.g., recordings)
GRANT ALL PRIVILEGES ON recordings.* TO 'username'@'localhost';

-- 3. Save changes and exit
FLUSH PRIVILEGES;
EXIT;

### Query for multiple rows
- in the func albumsByArtist it creates an empty slice
- sends a SQL query to the db
- reads very row returned
- converts each row into an Album struct
- retunrs the completed slice
- Query() is used when you expect multiple rows
- QueryRow() is used for one expected row
- rows.Scan() transfers data from the current db row into the struct. the &alb is used to detec the adressess of the variables so it can write values into them.
- rows.Err() is used to check if there are any other errors during the iteration not during the initial query or the individual Scan() calls