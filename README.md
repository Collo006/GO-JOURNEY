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