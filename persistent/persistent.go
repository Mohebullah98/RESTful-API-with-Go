package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// This file is different from main because this allows persistent data storage.
// Albums are stored in a json file.
// They are then copied to an albums array where we can modify them for requests.
// Sending a post or delete request will update or remove the specified album from our albums array.
// We then re-write to the json file with the updated albums array.
// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums []album

func main() {
	// Read our json file and convert all albums from json file into albums array
	file, err := os.Open("albums.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&albums)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.DELETE("/albums/:id", deleteAlbumByID)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)

	// Update our json file by writing our albums array to it. Will not overwrite
	writeToJson("albums.json")

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// This function updates our json file with the contents of the album array.
func writeToJson(fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "   ")
	err = encoder.Encode(albums)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//delete an album given its ID in the path.
func deleteAlbumByID(c *gin.Context){
	id := c.Param("id")
	
	for i, a := range albums {
		if a.ID == id {
			//albums is set to a slice of all albums upto the one we want to delete and
			//all albums after the one we want to delete.
			albums = append(albums[:i], albums[i+1:]...)
			writeToJson("albums.json")
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
	    }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}