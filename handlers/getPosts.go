package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//GetPosts Get Feed
func GetPosts(w http.ResponseWriter, r *http.Request) {

	svc := GetS3()
	var posts []Post
	var post Post
	bucketRegion := os.Getenv("BUCKET_REGION")

	result, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(os.Getenv("BUCKET_NAME"))})
	if err != nil {
		fmt.Println(err)
	}
	for i, item := range result.Contents {
		if i != 0 {
			var url string = "https://s3-" + bucketRegion + ".amazonaws.com/" + os.Getenv("BUCKET_NAME") + *item.Key
			post.Name = strings.Split(*item.Key, "/")[1]
			post.Link = url
			posts = append(posts, post)
		}
	}
	postJSON, errm := json.Marshal(posts)
	if errm != nil {
		panic(errm)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(postJSON)
}
