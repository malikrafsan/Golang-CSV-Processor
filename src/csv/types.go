package csv

type UploadCSVData struct {
	FileID string `json:"file_id"`
}

type GetProcessedCSVData struct {
	FileID string `json:"file_id"`
	Sum    string `json:"sum"`
}
