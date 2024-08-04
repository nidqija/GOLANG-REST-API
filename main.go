package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type randomActivity struct{
	ID     string   `json:"id"`
    Title  string   `json:"title"`
	Task   string   `json:"task"`
	Completed bool  `json:"completed"`
}

var activities =[]randomActivity{
    { ID : "1" , Title : "Clean garage" , Task : "Wash and sweep the garage , your garage stinks!" , Completed: false} ,
	{ ID : "2" , Title : "Read book" , Task : "Read books that you havent completed yet" , Completed: false} ,
    { ID : "3" , Title : "Play games" , Task : "Play your favourite game for 2 hours" , Completed: false} ,
    { ID : "4" , Title : "Buy groceries" , Task : "Go to the grocery and buy something you think is interesting" , Completed: false} ,
    { ID : "5" , Title : "Wash your car" , Task : "Wash your car , your car reeks of stinky poo poo!" , Completed: false} ,

}

func getActivities(context *gin.Context){ 
   context.IndentedJSON(http.StatusOK , activities)
}

func addActivities( context *gin.Context){
	var newActivity randomActivity 

	if err := context.BindJSON(&newActivity) ; err != nil {
		return 
	}

	activities = append(activities , newActivity)

	context.IndentedJSON(http.StatusCreated , newActivity)

}

func main(){
	router := gin.Default()
	router.GET("/activity"  , getActivities)
	router.POST("/activity" , addActivities)
	router.Run("localhost:9090")
}