package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//GetPosts Get Feed
func GetPosts(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	svc := GetS3()
	var post Post
	bucketRegion := os.Getenv("BUCKET_REGION")

	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(os.Getenv("BUCKET_NAME"))})
	if err != nil {
		fmt.Println(err)
	}
	post.Name = strings.Split(*result.Contents[id].Key, "/")[1]
	post.Link = "https://s3-" + bucketRegion + ".amazonaws.com/" + os.Getenv("BUCKET_NAME") + *result.Contents[id].Key
	postJSON, errm := json.Marshal(post)
	if errm != nil {
		panic(errm)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(postJSON)
}
