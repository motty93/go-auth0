package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	Env "github.com/joho/godotenv"
)

type Result struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

var (
	url          string
	clientID     string
	clientSecret string
	audience     string
	userName     string
	password     string
	payloadStr   string
	payload      *strings.Reader
	result       Result
)

func init() {
	// envの読み込み処理
	err := Env.Load(".env")
	if err != nil {
		fmt.Printf(".env read error: %v", err)
	}

	url = os.Getenv("AUTH0_ENDPOINT")
	clientID = os.Getenv("AUTH0_CLIENT_ID")
	clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	audience = os.Getenv("AUTH0_AUDIENCE")
	userName = os.Getenv("USER_NAME")
	password = os.Getenv("PASSWORD")
	// payloadStr = "grant_type=client_credentials&client_id=" + clientID + "&client_secret=" + clientSecret + "&audience=" + audience
	payloadStr = "grant_type=password&username=" + userName + "&password=" + password + "&client_id=" + clientID + "&client_secret=" + clientSecret + "&audience=" + audience
}

func main() {
	payload = strings.NewReader(payloadStr)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("user: ", userName)
	fmt.Println("password: ", password)
	fmt.Println("access_token: ", result.AccessToken)

	clipboard.WriteAll(result.AccessToken)
	fmt.Println("access_tokenがクリップボードにコピーされました！")
}
