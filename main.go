package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

var authcookie string

func main() {
	ReadFile, _ := os.ReadFile("Account.txt")
	if string(ReadFile) == "" {
		fmt.Println("Couldn't find account\nEnter your account info in user:password format")
		var UserPass string
		fmt.Scanln(&UserPass)
		os.WriteFile("Account.txt", []byte(UserPass), 0)
	}
	GetAuthCookie()
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
	client := &http.Client{}
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
		response, _ := client.Do(request)
		fmt.Println(response.Status)
	}
}

func UserSearch(UserID string) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/users/"+UserID+"?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat", nil)
	request.Header = http.Header{
		"X-Requested-With": {"XMLHttpRequest"},
		"X-Macaddress":     {"3ceb8cc8df874eeba6ff679158315ec54f1470eb"},
		"X-Client-Version": {"2022.2.2p2c-1221--Release"},
		"X-Platform":       {"standalonewindows"},
		"X-Unity-Version":  {"2019.4.31f1"},
		"Content-Type":     {"application/x-www-form-urlencoded"},
		"Origin":           {"vrchat.com"},
		"Host":             {"api.vrchat.cloud"},
		"Connection":       {"Keep-Alive, TE"},
		"TE":               {"identity"},
		"User-Agent":       {"VRC.Core.BestHTTP"},
		"Cookie":           {"auth=; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth="},
		"accept":           {"*/*"},
		"Accept-Encoding":  {"identity"},
	}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)
	Formated := strings.Split(string(body), ",")
	for i := 0; i < len(Formated); i++ {
		fmt.Println(Formated[i])
	}
}

func GetAuthCookie() {
	ReadFile, _ := os.ReadFile("Account.txt")
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/auth/user?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat", nil)
	request.Header = http.Header{
		"X-Requested-With": {"XMLHttpRequest"},
		"X-MacAddress":     {"3ceb8cc8df874eeba6ff679158315ec54f1470eb"},
		"X-Client-Version": {"2022.2.2p2c-1221--Release"},
		"X-Platform":       {"standalonewindows"},
		"X-Unity-Version":  {"2019.4.31f1"},
		"Content-Type":     {"application/x-www-form-urlencoded"},
		"Origin":           {"vrchat.com"},
		"Host":             {"api.vrchat.cloud"},
		"Connection":       {"Keep-Alive, TE"},
		"TE":               {"identity"},
		"User-Agent":       {"VRC.Core.BestHTTP"},
		"Authorization":    {"Basic " + base64.StdEncoding.EncodeToString(ReadFile)},
		"Cookie":           {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth="},
		"Accept-Encoding":  {"identity"},
	}
	response, _ := client.Do(request)
	body, _ := io.ReadAll(response.Body)
	Format2 := strings.Split(string(body), "\"")
	fmt.Println("Logged in as: " + Format2[11])
	authcookie = response.Cookies()[0].Value
	Start()
}
