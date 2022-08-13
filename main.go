package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var authcookie string
var DefaultClient = &http.Client{Timeout: time.Second * 10}

func main() {
	ReadFile, _ := os.ReadFile("Account.txt")
	if string(ReadFile) == "" {
		fmt.Println("Couldn't find account\nEnter your account info in user:password format")
		var UserPass string
		fmt.Scanln(&UserPass)
		os.WriteFile("Account.txt", []byte(UserPass), 0)
	}
	GetAuthCookie(ReadFile)
}

func Start() {
	fmt.Println("-----Pick a option-----")
	fmt.Println("1: Request spam UserID\n2: UserID lookup")
	var Input string
	fmt.Scanln(&Input)
	switch Input {
	case "1":
		fmt.Println("Enter UserID to request spam")
		var UserID string
		fmt.Scanln(&UserID)
		for i := 0; i < 5; i++ {
			go RequestSpam(UserID)
		}
		RequestSpam(UserID)
	case "2":
		fmt.Println("Enter UserID to lookup")
		var UserID string
		fmt.Scanln(&UserID)
		UserSearch(UserID)
	default:
		fmt.Println("Pick a valid option")
		time.Sleep(3 * time.Second)
		Start()
	}
}

func RequestSpam(UserID string) {
	for i := 0; i < 25; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/requestInvite/"+UserID, strings.NewReader("{\"platform\":\"standalonewindows\"}"))
		request.Header = http.Header{
			"Cookie":           {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth=" + authcookie},
			"X-Client-Version": {"2022.2.2-1213--Release"},
			"X-Platform":       {"standalonewindows"},
			"Content-Type":     {"application/json"},
			"User-Agent":       {"Transmtn-Pipeline"},
			"Host":             {"api.vrchat.cloud"},
		}
		response, _ := DefaultClient.Do(request)
		fmt.Println(response.Status)
	}
}

func UserSearch(UserID string) {
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/users/"+UserID+"?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat", nil)
	request.Header = http.Header{
		"Content-Type":    {"application/x-www-form-urlencoded"},
		"Origin":          {"vrchat.com"},
		"Host":            {"api.vrchat.cloud"},
		"Connection":      {"Keep-Alive, TE"},
		"TE":              {"identity"},
		"User-Agent":      {"VRC.Core.BestHTTP"},
		"Cookie":          {"auth=" + authcookie + "; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth="},
		"accept":          {"*/*"},
		"Accept-Encoding": {"identity"},
	}
	response, _ := DefaultClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	Formated := strings.Split(string(body), ",")
	for i := 0; i < len(Formated); i++ {
		fmt.Println(Formated[i])
	}
}

func GetAuthCookie(Account []byte) {
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/auth/user?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat", nil)
	request.Header = http.Header{
		"Content-Type":    {"application/x-www-form-urlencoded"},
		"Origin":          {"vrchat.com"},
		"Host":            {"api.vrchat.cloud"},
		"User-Agent":      {"VRC.Core.BestHTTP"},
		"Authorization":   {"Basic " + base64.StdEncoding.EncodeToString(Account)},
		"Cookie":          {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26;"},
		"Accept-Encoding": {"identity"},
	}
	response, _ := DefaultClient.Do(request)
	fmt.Println(response.Status)
	body, _ := io.ReadAll(response.Body)
	format := strings.Split(string(body), "\"")
	log.Println("Logged in as: " + format[11])
	authcookie = response.Cookies()[0].Value
	Start()
}
