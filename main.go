package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostFoo struct {
	Name string `json:"name"`
}

type Foo struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func getFoo(ctx *gin.Context) {
	id := ctx.Param("id")
	guid, _ := uuid.Parse(id)
	for _, f := range foos {
		if f.ID == guid {
			ctx.IndentedJSON(http.StatusOK, f)
			return
		}
	}
	err := fmt.Sprintf("Specified foo with ID:%s not found", id)
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
}

func addFoo(ctx *gin.Context) {
	var requestFoo PostFoo
	if err := ctx.BindJSON(&requestFoo); err != nil {
		err := fmt.Sprintf("Failed to create foo: %v", err)
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": err})
		return
	}
	newFoo := Foo{
		ID:   uuid.New(),
		Name: requestFoo.Name,
	}
	foos = append(foos, newFoo)
	ctx.IndentedJSON(http.StatusCreated, newFoo)
}

func deleteFoo(ctx *gin.Context) {
	id := ctx.Param("id")
	guid, _ := uuid.Parse(id)
	var fooIndex = -1

	for i, f := range foos {
		if f.ID == guid {
			fooIndex = i
		}
	}

	if fooIndex == -1 {
		err := fmt.Sprintf("Cannot delete foo with ID:%s. Foo with that id not found", id)
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}
	foos = append(foos[:fooIndex], foos[fooIndex+1:]...)
	ctx.IndentedJSON(http.StatusNoContent, nil)
}

var foos = []Foo{}

func main() {
	router := gin.Default()
	router.GET("/foo/:id", getFoo)
	router.DELETE("/foo/:id", deleteFoo)
	router.POST("/foo", addFoo)

	router.Run("localhost:8080")
}
