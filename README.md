# RESTful-API-with-Go
### RESTful API for albums

#### We have a simple API that stores albums in memory, meaning the storage is temporary.
```Each album follows the template:```
```go
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

```
The following methods are supported:
```
/albums
   -Get: Get a list of all albums, returned as JSON.
   -POST:  Add a new album from request data sent as JSON.
 /albums/:id
   -Get: Get an album by its id, returning the album data as JSON.
   -Delete: Delete an album, given its id.
 ```
 
 How to run the app:
  1. Clone the repo
   ```
   git remote add origin https://github.com/Mohebullah98/RESTful-API-with-Go.git
   ```
  2. Open a terminal and navigate into project directory, download necessary dependencies:
  ```
  cd RESTful-API-with-Go
  go mod tidy
  ```
  3. Run the code
  ```
  go run .
  ```
  4. Now you can fetch all albums by visiting ```http://localhost:8080/albums``` in browser or by using a curl command.
  ``` curl http://localhost:8080/albums ```
  You can also try getting an album by it's id or even deleting or adding a new albums with the curl command.
  ```
  curl http://localhost:8080/albums \
      --include \
      --header "Content-Type: application/json" \
      --request "POST" \
      --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
  curl http://localhost:8080/albums/2
  curl -X DELETE http://localhost:8080/albums/3
   
   ```
   
   ## Extra
   - Persistent storage was achieved by using a json api. All routes and methods are the same as original. The only difference is album data is stored in a json file.
   ```albums.json
   [
   {
      "id": "1",
      "title": "Blue Train",
      "artist": "John Coltrane",
      "price": 56.99
   },
   {
      "id": "2",
      "title": "Jeru",
      "artist": "Gerry Mulligan",
      "price": 17.99
   },
   {
      "id": "3",
      "title": "Sarah Vaughan and clifford Brown",
      "artist": "Sarah Vaughan",
      "price": 39.99
   },
   {
      "id": "4",
      "title": "The Modern Sound of Betty Carter",
      "artist": "Betty Carter",
      "price": 49.9
   },
   {
      "id": "5",
      "title": "American Idiot",
      "artist": "Green Day",
      "price": 59.99
   },
   {
      "id": "6",
      "title": "Recovery",
      "artist": "Eminem",
      "price": 69.99
   }
]
```
   - Run the json api inside the /persistent directory. All commands should be the same.
   
  
