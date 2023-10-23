package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// album represents(代表) data about a record album.
type album struct {
	ID     string  `json:"id"` // 指定结构体序列化为JSON是字段的名称
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// gin.Context 是Gin中最重要的部分。它携带请求详细信息、验证和序列化 JSON 等等。
func getAlbums(c *gin.Context) {
	// 调用 Context.IndentedJSON 将结构序列化为 JSON 并将其添加到响应中
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Println(err)
		return
	}

	albums = append(albums, newAlbum)
	// c.JSON(http.StatusCreated, newAlbum) // 更加紧凑
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbumByID(c *gin.Context) {
	id := c.Params.ByName("id")
	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func handleQueryArg(c *gin.Context) {
	username := c.Query("username")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
	})
}
func handlePostArg(c *gin.Context) {
	username := c.PostForm("username")
	password := c.DefaultPostForm("password", "123456") // 默认值"123456"
	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}
func handleGetFormArg(c *gin.Context) {
	username := c.Query("username")
	password := c.DefaultQuery("password", "123456") // 默认值"123456"
	c.IndentedJSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
	})
}
func handleFormArgToStruct(c *gin.Context) {
	var album album

	if err := c.BindJSON(&album); err != nil {
		log.Println(err)
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.GET("/queryArg", handleQueryArg)

	router.POST("/postFormArg", handlePostArg)
	router.GET("/GetFormArg", handleGetFormArg)
	router.POST("/formArgToStruct", handleFormArgToStruct)

	router.GET("/websocket", func(ctx *gin.Context) {
		handleWebsocket(ctx.Writer, ctx.Request)
	})
	router.Run("localhost:8080")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleWebsocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// defer ws.Close()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println("远程地址:", ws.RemoteAddr(), ",received: ", string(msg))

		if err = ws.WriteMessage(msgType, msg); err != nil {
			log.Println(err)
			return

		}
	}

}
