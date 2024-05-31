# HOTOGO-LESSON5

# RESTAPI-GO-GIN

## 1.Set up

Install gin by this command line:

```go
go get github.com/gin-gonic/gin
```

## 2.Get method analyze:

```go
func get_player_by_id(c *gin.Context) {
	id := c.Param("id")

	for _, value := range madridteam {
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "can not find the player"})
}
```

Let’s take this function and analyze:

- **gin.Context**:  a structure that contains both the http.Request and the http.Response that a normal http.Handler would use. You can think of that as the request link “http:localhost:8080/info”
- **Param, IndentedJson** are methods of Context struct. Firstly, Param helps us extract information of the request for further coding in the long run. Secondly, IndentedJson helps to return the ht-tp response. But the gin team’s WARNING:” we recommend using this only for development p-urposes since printing pretty JSON is more CPU and bandwidth consuming. Use Context.JSON() instead”
- **gin.H{””:””}** is a helper type that represents a dictionary or hash map of key-value pairs. It's sp-ecifically designed for simplifying the creation of JSON responses in Gin applications.

## [3.](http://3.POST) POST METHOD

```go
func post_a_player(c *gin.Context) {
	var new_player Footballer

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&new_player); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "can not add a player"})
		return
	}

	//add a new player to slice
	madridteam = append(madridteam, new_player)
	c.IndentedJSON(http.StatusOK, new_player)
}
```

- **BindJson:** refers to the process of unmarshalling JSON data from an HTTP request body into a Go struct within the Gin framework. Gin provides a built-in function called `c.BindJSON` that performs this task.

## 3. API ENDPOINT

**API ENDPOINT** is a specific Uniform Resource Identifier (URI) that serves as an entry point for interacting with an application's functionalities.

1. **Code**

```go
	router := gin.Default()
	router.GET("/players/", getPlayers)
	router.POST("/players/", post_a_player)
	router.GET("/players/:id", get_player_by_id)
	router.Run("localhost:8080")
```

.

- **gin.Default:** function serves as a convenient way to create a new Gin engine instance with some pre-configured settings suitable for most web applications.
2. **Grouping**

```go
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/players/", getPlayers)
		v1.POST("/players/", post_a_player)
		v1.GET("/players/:id", get_player_by_id)
	}
	router.Run("localhost:8080")
```

In addition, we  can group the endpoint by doing so. It is a convenient way to manage versioning, middleware scope,..

## References:

[Developing a RESTful API with Go and Gin - The Go Programming Language](https://go.dev/doc/tutorial/web-service-gin)
