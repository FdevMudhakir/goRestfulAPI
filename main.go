package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID	string		`json:"id"`
	Item	string	`json:"item"`
	Completed	bool `json:"completed"`
}


var todos = []todo{
	{ID: "1", Item:"Clean", Completed:false},
	{ID: "2", Item:"Read Book", Completed:false},
	{ID: "3", Item:"Record Video", Completed:false},	
}

//API for getting Todos from the Go.
func getTodos(context	*gin.Context) {
	//Here we are converting all the todos received as context to json
	context.IndentedJSON(http.StatusOK, todos)

}

//API for adding/posting newTodos to the existing one from the Admin/Postman.
func addTodo(context *gin.Context){
	var newTodo todo
	//The BindJSON method is used to convert from JSON format to Go data structure.
	if err := context.BindJSON(&newTodo); err != nil{
		return
	}
	todos = append(todos, newTodo)
	//The IndentedJSON method is used to convert back to JSON from Go data structure.
	context.IndentedJSON(http.StatusCreated, newTodo)
}


//API for getting a specific todo with an id
func getTodo (context *gin.Context){
	//We are extracting a specific todo from our todos with id here.
	id := context.Param("id")
	todo, err := getTodoById((id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, todo)
}

//API for checking a todo status whether completed or not.
func toggleTodoStatus(context *gin.Context){
	//We are extracting a specific todo from our todos with id here.
	id := context.Param("id")
	todo, err := getTodoById((id))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Todo not found"})
		return 
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

//API for getting a specific todo with an id
func getTodoById (id string) (*todo, error){
	//The * infront of todo that is return means return all todos with no error.
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("todos you are looking for not found")
}


func main (){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", addTodo)	
	router.Run("localhost:1990")
}