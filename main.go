package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
	//	var post Post
	//	data, _ := json.MarshalIndent(post, "", " ")
	//	fmt.Printf("%s\n", data)

	// command-line arguments
	var tags string

	flag.StringVar(&tags, "t", "", "A comma-separated list of tags")

	flag.Usage = func() {
		fmt.Println("Usage: mediaup [OPTION]... TITLE FILE")
		fmt.Printf("\nAvailable options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(1)
	}

	//fmt.Printf("%v\n", flag.NArg())

	args := flag.Args()
	post := Post{
		Title:         args[0],
		ContentFormat: "markdown",
		PublishStatus: "draft",
	}

	//filePath := args[1]

	data, _ := json.MarshalIndent(post, "", " ")
	fmt.Printf("%s\n", data)

}
