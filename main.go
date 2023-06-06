package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/malikrafsan/Golang-CSV-Processor/src/csv"
	"github.com/malikrafsan/Golang-CSV-Processor/src/utils"
)

func ping(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "pong")
}

func main() {
	router := gin.Default()
	router.GET("/ping", ping)

	runId, err := utils.IdFromTime()
	if err != nil {
		panic(err)
	}

	logFile, err := os.Create("logs/" + runId + ".log")
	if err != nil {
		panic(err)
	}

	log.SetOutput(logFile)

	router.POST("/csv", csv.UploadCSV)
	router.GET("/csv", csv.GetProcessedCSV)

	router.Run("localhost:8080")
}
