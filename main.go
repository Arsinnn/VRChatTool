package main

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	Start()
}

var authcookie string

func Start() {
	Reader := bufio.NewReader(os.Stdin)
	if _, err := os.Stat("Authcookie.txt"); errors.Is(err, os.ErrNotExist) {
		os.Create("Authcookie.txt")
	}
	if _, err := os.Stat("Authcookies.txt"); errors.Is(err, os.ErrNotExist) {
		os.Create("Authcookies.txt")
	}
	if _, err := os.Stat("Avatars.txt"); errors.Is(err, os.ErrNotExist) {
		os.Create("Avatars.txt")
	}
	joe, _ := os.ReadFile("Authcookie.txt")
	if !(strings.Contains(string(joe), "authcookie_")) {
		fmt.Println("authcookie not found enter your authcookie")
		Input, _ := Reader.ReadString('\n')
		os.WriteFile("Authcookie.txt", []byte(Input), 0)
	}
	joe2, _ := os.ReadFile("Authcookie.txt")
	authcookie = string(joe2)
	fmt.Println("-----Pick a option-----")
	fmt.Println("1: Request spam UserID")
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
	default:
		fmt.Println("Pick a valid option")
		time.Sleep(3 * time.Second)
		Start()
	}
}

func RequestSpam(UserID string) {
	client := http.Client{}
	for i := 0; i < 25; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/requestInvite/"+UserID, strings.NewReader("{\"platform\":\"standalonewindows\"}"))
		request.Header.Add("Cookie", "apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth="+authcookie)
		request.Header.Add("X-Client-Version", "2022.2.2-1213--Release")
		request.Header.Add("X-Platform", "standalonewindows")
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("User-Agent", "Transmtn-Pipeline")
		request.Header.Add("Host", "api.vrchat.cloud")
		response, _ := client.Do(request)
		fmt.Println(response.Status)
	}
}
