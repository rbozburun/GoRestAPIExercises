package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean room", Completed: false},
	{ID: "2", Item: "Read book", Completed: false},
	{ID: "3", Item: "Record video", Completed: false},
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func getTodoByID(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found!")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")

	todo, err := getTodoByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func addTodo(context *gin.Context) {
	var newTodo todo

	// Create todo obj with given json value. If there is an error, err not be nil and return.
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found!"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()                      // Router is our server.
	router.GET("/todos", getTodos)               // -> List todos
	router.GET("/todos/:id", getTodo)            // Get specificed todo with id
	router.PATCH("/todos/:id", toggleTodoStatus) // Update the specificed ID todo
	router.POST("/todos", addTodo)               // -> Create a new todo
	router.Run("localhost:9090")

}
