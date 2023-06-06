package csv

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/malikrafsan/Golang-CSV-Processor/src/contracts"
	"github.com/malikrafsan/Golang-CSV-Processor/src/utils"
)

func UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)

		resp := contracts.FailedResponse(
			"Failed to upload file",
			http.StatusBadRequest)

		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	extension := filepath.Ext(file.Filename)
	if extension != ".csv" {
		resp := contracts.FailedResponse(
			"Invalid file type",
			http.StatusBadRequest)
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	fileID, err := utils.IdFromTime()
	if err != nil {
		log.Println(err)
		resp := contracts.FailedResponse(
			"Failed to generate file ID",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	dst := "dumps/" + fileID + extension
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		log.Println(err)
		resp := contracts.FailedResponse(
			"Failed to save uploaded file",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	err = ProcessFile(c.Request.Context(), file, fileID)
	if err != nil {
		log.Println(err)
		resp := contracts.FailedResponse(
			"Failed to process file",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := contracts.SuccessResponse(
		"File uploaded successfully",
		UploadCSVData{
			FileID: fileID,
		})
	c.IndentedJSON(http.StatusOK, resp)
}

func GetProcessedCSV(c *gin.Context) {
	ctx := c.Request.Context()

	fileID := c.Query("file_id")

	data, err := GetProcessedFile(ctx, fileID)
	if err != nil {
		log.Println(err)
		resp := contracts.FailedResponse(
			"Failed to get processed file",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	resp := contracts.SuccessResponse(
		"File processed successfully",
		data)
	c.IndentedJSON(http.StatusOK, resp)
}
