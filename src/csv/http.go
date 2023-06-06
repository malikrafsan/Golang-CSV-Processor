package csv

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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

	resp := contracts.SuccessResponse(
		"File uploaded successfully",
		UploadCSVData{
			FileID: fileID,
		})
	c.IndentedJSON(http.StatusOK, resp)
}

func GetProcessedCSV(c *gin.Context) {
	fileID := c.Query("file_id")

	pathname := "dumps/" + fileID + ".csv"
	f, err := os.Open(pathname)
	if err != nil {
		log.Println(err)

		resp := contracts.FailedResponse(
			"File not found",
			http.StatusNotFound)
		c.IndentedJSON(http.StatusNotFound, resp)
		return
	}

	csvReader, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Println(err)

		resp := contracts.FailedResponse(
			"Failed to read file",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	sums, err := SumCSV(&csvReader)
	if err != nil {
		log.Println(err)

		resp := contracts.FailedResponse(
			"Failed to sum file",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	strSums, err := utils.IntSliceToStrSlice(&sums)
	if err != nil {
		log.Println(err)

		resp := contracts.FailedResponse(
			"Failed to convert sums to string",
			http.StatusInternalServerError)
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	strSumsJoined := strings.Join(strSums, ",")

	resp := contracts.SuccessResponse(
		"File processed successfully",
		GetProcessedCSVData{
			FileID: fileID,
			Sum:    strSumsJoined,
		})
	c.IndentedJSON(http.StatusOK, resp)
}
