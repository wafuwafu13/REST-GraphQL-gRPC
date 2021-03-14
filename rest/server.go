package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	post struct {
		ID      int    `json:"id"`
		UserID  int    `json:"userId"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}
)

var users = map[string]user{
	"1": {
		ID:   1,
		Name: "foo",
	},
	"2": {
		ID:   2,
		Name: "bar",
	},
}

var posts = map[string]post{
	"1": {
		ID:      1,
		UserID:  1,
		Title:   "title1",
		Content: "content1",
	},
	"2": {
		ID:      2,
		UserID:  1,
		Title:   "title2",
		Content: "content2",
	},
	"3": {
		ID:      3,
		UserID:  2,
		Title:   "title2",
		Content: "content2",
	},
}

func main() {
	e := echo.New()
	e.GET("/users", getUsers)
	e.GET("/posts", getPosts)
	e.GET("/post/:userId", getPost)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, users)
}

func getPosts(c echo.Context) error {
	return c.JSON(http.StatusOK, posts)
}

func getPost(c echo.Context) error {
	p := c.Param("userId")
	userId, _ := strconv.Atoi(p)
	userPost := map[string]post{}
	for _, value := range posts {
		if value.UserID == userId {
			userPost[strconv.Itoa(value.ID)] = value
		}
	}
	return c.JSON(http.StatusOK, userPost)
}
