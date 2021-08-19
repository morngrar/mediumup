package main

import (
	"encoding/json"
	"fmt"
)

const getEndpoint = "https://api.medium.com/v1/me"
const postEndpointTemplate = "https://api.medium.com/v1/users/%s/posts"

type Post struct {
	Title         string   `json:"title"`
	ContentFormat string   `json:"contentFormat"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags,omitempty"`
	PublishStatus string   `json:"publishStatus"`
}

func main() {
	var post Post
	data, _ := json.MarshalIndent(post, "", " ")
	fmt.Printf("%s\n", data)

}
