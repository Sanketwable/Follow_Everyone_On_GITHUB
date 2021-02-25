package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Token is a bearer token
// You can get it from GITHUB>>Setting>>developer Setting>>PersonalAccessToken>>GenerateNewToken>>COPYandPASTE here.
var Token string = "ee3068ce4601cec21da6cc07827ea74c4b06ee5"
//Change only token
//
//
// Code Starts
var bearerToken string = "Bearer " + Token

// Response is a struct to accept the user response from the GET call
type Response struct {
	Name string `json:"login"`
	ID   int    `json:"id"`
}

func main() {
	from := 628740
	for {
		from := 628740
		res := getuser(from)
		for _, j := range res {
			follow(j.Name)
			from = j.ID
			fmt.Println("followed : ", j.Name)
		}
		fmt.Println("from  = ", from)
	}
}

// follow func is to follow a particular user which takes username as argument
func follow(name string) {
	baseurl := "https://api.github.com/user/following/"
	finalurl := baseurl + name
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, finalurl, nil)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	req.Header.Add("Authorization", bearerToken)
	req.Header.Add("Content-Length", "0")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	fmt.Println(res.StatusCode)
	if res.StatusCode != 204 {
		fmt.Println("need to sleep at this point")
		time.Sleep(time.Minute)
		follow(name)
	}
	defer res.Body.Close()
}

// getuser is a func to get user info (i.e user id and username)
func getuser(from int) []Response {
	//url is a base url for GET api
	url := "https://api.github.com/users?since="
	add := strconv.Itoa(from) + ">"
	finalURL := url + add

	r, err := http.Get(finalURL)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	r.Header.Add("Bearer", Token)

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("error occured ", err)
	}
	fmt.Println(r.StatusCode)
	res := []Response{}
	json.Unmarshal(response, &res)
	return res

}
