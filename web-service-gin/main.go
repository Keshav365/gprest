package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"web-services-gin/models"

	"github.com/gin-gonic/gin"
	//"./taskManager.db"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	r := gin.Default()
	router := r.Group("/tasks")
	{
		router.POST("/", addTask)
		router.GET("/:id", getTask)
		router.GET("/", getAllTask)
		router.PUT("/:id", updateTask)
		router.DELETE("/:id", deleteTask)
	}

	go func() {
		if err := r.Run(":8099"); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	// Handle termination signals to clean up
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

}

func getTask(c *gin.Context) {
	iid := c.Param("id")
	var i int
	if _, err := fmt.Sscan(iid, &i); err == nil {
		fmt.Printf("i=%d, type: %T\n", i, i)
	}
	fmt.Println(iid)
	tasks, err := models.GetTask(int(i))
	checkErr(err)
	if len(tasks) < 1 {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": tasks})
	}
}

func getAllTask(c *gin.Context) {
	tasks, err := models.GetAllTasks()
	checkErr(err)

	if tasks == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(200, gin.H{"data": tasks})
	}
}

func addTask(c *gin.Context) {

	rawData, err1 := c.GetRawData()
	if err1 != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	fmt.Println("Raw JSON Payload:", string(rawData))

	var task models.Task
	if err := json.Unmarshal(rawData, &task); err != nil {
		c.JSON(400, gin.H{"error": "Failed to bind JSON payload"})
		return
	}

	if task.Title == "" || task.Description == "" || task.Due_Date == "" {
		c.JSON(400, gin.H{"error": "Missing required fields (title, description, due_date)"})
		return
	}

	if task.Status == "" {
		task.Status = "Pending"
	}

	err := models.AddTask(task.Title, task.Description, task.Due_Date, task.Status)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to add task"})
		return
	}

	c.JSON(201, gin.H{
		"message": "Task added successfully",
		"data":    task,
	})
}

func updateTask(c *gin.Context) {
	// Extract task ID from URL parameter
	iid := c.Param("id")
	var taskId int
	if _, err := fmt.Sscan(iid, &taskId); err != nil || taskId == 0 {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	// Parse the JSON payload for the updated task details
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON payload"})
		return
	}

	// Validate required fields in the payload
	if task.Title == "" || task.Description == "" || task.Due_Date == "" || task.Status == "" {
		c.JSON(400, gin.H{"error": "Missing required fields (title, description, due_date, status)"})
		return
	}

	// Call the models.UpdateTask function to update the task in the database
	err := models.UpdateTask(taskId, task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update task", "details": err.Error()})
		return
	}

	// Return the updated task with a success message
	c.JSON(200, gin.H{"message": "Task updated successfully", "data": task})
}

func deleteTask(c *gin.Context) {
	iid := c.Param("id")

	var i int
	if _, err := fmt.Sscan(iid, &i); err != nil || i == 0 {
		c.JSON(400, gin.H{"error": "Invalid task ID"})
		return
	}

	err := models.DeleteTask(i)

	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete task", "details": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deleted successfully", "task_id": iid})
}
