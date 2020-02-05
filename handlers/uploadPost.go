package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//UploadPost Upload to S3
func UploadPost(w http.ResponseWriter, r *http.Request) {

	maxSize := int64(5 * 1024000) //5 MB
	var resp Response

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		fmt.Println("Image too large.")
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Error : Image Too Large"
	}
	file, fileHeader, err := r.FormFile("post")
	if err != nil {
		fmt.Println("Could not get uploaded file")
		w.WriteHeader(http.StatusBadRequest)
		resp.Message = "Error : Cant get File"
	}
	defer file.Close()
	fmt.Println(fileHeader.Header)

	s, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("BUCKET_REGION"))},
	)
	uploader := s3manager.NewUploader(s)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("onehearttour/posts"),
		Key:    aws.String(fileHeader.Filename),
		Body:   file,
	})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		resp.Message = "Error : AWS Error"
	} else {
		w.WriteHeader(http.StatusOK)
		resp.Message = "Success"
	}
	respJSON, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}
