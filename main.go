package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

type MeResponse struct {
	Data struct {
		ID string
	}
}

type Config struct {
	Token string `json:"integrationToken"`
}

func main() {
	var token string

	configDirPath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}
	configDirPath += "/mediumup"

	// making sure config dir exists
	err = os.MkdirAll(configDirPath, 0777)
	if err != nil {
		panic(err)
	}

	configPath := configDirPath + "config.json"

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("Unable to read file: %v", err)
		var input string
		fmt.Println("Enter your integration token: ")
		fmt.Scanln(&input)
		cfg := Config{Token: input}
		data, err := json.Marshal(cfg)
		if err != nil {
			panic(err)
		}
		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			panic(err)
		}

		token = input
	} else {
		var cfg Config
		err := json.Unmarshal(configFile, &cfg)
		if err != nil {
			panic(err)
		}

		token = cfg.Token
	}

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

	if tags != "" {
		post.Tags = strings.Split(tags, ",")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Getting home dir resulted in error: %v\n", err)
	}

	path := strings.Replace(args[1], "~", homeDir, -1)
	fileContents, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	post.Content = string(fileContents)

	//	data, _ := json.MarshalIndent(post, "", " ")
	//	fmt.Printf("%s\n", data)

	// generating user request
	dummy := []byte{}
	meReq, err := http.NewRequest("GET", getEndpoint, bytes.NewBuffer(dummy))
	if err != nil {
		panic(err)
	}

	meReq.Header.Set("Authorization", "Bearer "+token)
	meReq.Header.Set("Accept", "application/json")
	meReq.Header.Set("Accept-Charset", "utf-8")
	meReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	meResponse, err := client.Do(meReq)
	if err != nil {
		panic(err)
	}
	defer meResponse.Body.Close()

	body, err := ioutil.ReadAll(meResponse.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("'Me' response:", meResponse.Status)

	var meBuf MeResponse
	err = json.Unmarshal(body, &meBuf)
	if err != nil {
		panic(err)
	}

	postURL := fmt.Sprintf(postEndpointTemplate, meBuf.Data.ID)
	postBytes, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	postReq, err := http.NewRequest("POST", postURL, bytes.NewBuffer(postBytes))
	if err != nil {
		panic(err)
	}

	postReq.Header.Set("Authorization", "Bearer "+token)
	postReq.Header.Set("Accept", "application/json")
	postReq.Header.Set("Accept-Charset", "utf-8")
	postReq.Header.Set("Content-Type", "application/json")

	postResponse, err := client.Do(postReq)
	if err != nil {
		panic(err)
	}
	defer postResponse.Body.Close()

	fmt.Println("post response:", postResponse.Status)
}
