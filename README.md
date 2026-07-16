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