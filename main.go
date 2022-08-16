package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var authcookie string

var UserID string
var DefaultClient = &http.Client{Timeout: 10 * time.Second}

var Proxys []string
var UserPass []string
var authcookies []string

func main() {
	FileCheck()
}

func FileCheck() {
	Files := []string{"Account.txt", "Proxys.txt", "Accounts.txt"}
	for i := 0; i < len(Files); i++ {
		ReadFile, _ := os.ReadFile(Files[i])
		if string(ReadFile) == "" {
			if i == 0 {
				fmt.Println("Couldn't find account\nEnter your account info in user:password format")
				var UserPass string
				fmt.Scanln(&UserPass)
				os.WriteFile(Files[i], []byte(UserPass), 0)
			} else {
				os.Create(Files[i])
			}
		}
	}
	Scanne("Proxys.txt", &Proxys)
	Scanne("Accounts.txt", &UserPass)
	Details, _ := os.ReadFile(Files[0])
	GetAuthCookie(Details)
}
func Scanne(F string, S *[]string) {
	File, _ := os.Open(F)
	Scanner := bufio.NewScanner(File)
	for Scanner.Scan() {
		*S = append(*S, Scanner.Text())
	}
}
func Start() {
	fmt.Println("-----Pick a option-----")
	fmt.Println("1: Request spam UserID\n2: UserID lookup\n3: Invite spam UserID\n4: Friend spam UserID using proxies and VRChat accounts")
	var Input string
	fmt.Scanln(&Input)
	switch Input {
	case "1":
		fmt.Println("Enter UserID to request spam")
		fmt.Scanln(&UserID)
		for i := 0; i < 9; i++ {
			go RequestSpam(UserID)
		}
		RequestSpam(UserID)
	case "2":
		fmt.Println("Enter UserID to lookup")
		fmt.Scanln(&UserID)
		UserSearch(UserID)
	case "3":
		fmt.Println("Enter UserID to Invite spam")
		fmt.Scanln(&UserID)
		for i := 0; i < 9; i++ {
			go InviteSpam(UserID)
		}
		InviteSpam(UserID)
	case "4":
		fmt.Scanln(&UserID)
		GetAuthCookies()
		for i := 0; i < len(authcookies); i++ {
			Proxy := strings.Split(Proxys[i], ":")
			FriendRequest(authcookies[i], ProxyC(Proxy))
		}
	default:
		fmt.Println("Pick a valid option")
		time.Sleep(3 * time.Second)
		Start()
	}
}
func ProxyC(Proxy []string) *http.Client {
	if len(Proxy) == 4 {
		return &http.Client{Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   Proxy[0] + ":" + Proxy[1],
				User:   url.UserPassword(Proxy[2], Proxy[3]),
			}),
		}}
	}
	return &http.Client{Transport: &http.Transport{
		Proxy: http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   Proxy[0] + ":" + Proxy[1],
		}),
	}}
}
func RequestSpam(UserID string) {
	for i := 0; i < 150/10; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/requestInvite/"+UserID, strings.NewReader("{\"platform\":\"standalonewindows\"}"))
		request.Header = http.Header{
			"Cookie":       {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth=" + authcookie},
			"Content-Type": {"application/json"},
			"User-Agent":   {"Transmtn-Pipeline"},
			"Host":         {"api.vrchat.cloud"},
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
		"TE":              {"identity"},
		"User-Agent":      {"VRC.Core.BestHTTP"},
		"Cookie":          {"auth=" + authcookie + "; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth="},
		"Accept":          {"*/*"},
		"Accept-Encoding": {"identity"},
	}
	response, _ := DefaultClient.Do(request)
	body, _ := io.ReadAll(response.Body)
	Formated := strings.Split(string(body), ",")
	for i := 0; i < len(Formated); i++ {
		fmt.Println(Formated[i])
	}
	Start()
}
func InviteSpam(UserID string) {
	for i := 0; i < 150/10; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/invite/"+UserID, strings.NewReader("{\"worldId\":\"wrld_532efaae-2f40-4f1d-a060-1d75719f5c2d:Joe\",\"instanceId\":\"wrld_532efaae-2f40-4f1d-a060-1d75719f5c2d:Joe\",\"worldName\":\"Offline\"}"))
		request.Header = http.Header{"Cookie": {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth=" + authcookie},
			"Content-Type": {"application/json"},
			"User-Agent":   {"Transmtn-Pipeline"},
			"Host":         {"api.vrchat.cloud"},
		}
		response, _ := DefaultClient.Do(request)
		fmt.Println(response.Status)
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
	body, _ := io.ReadAll(response.Body)
	format := strings.Split(string(body), "\"")
	log.Printf("| %v | Logged in as: %v", response.Status, format[11])
	authcookie = response.Cookies()[0].Value
	Start()
}
func GetAuthCookies() {
	if authcookies == nil {
		for i := 0; i < len(UserPass); i++ {
			Proxy := strings.Split(Proxys[i], ":")
			fmt.Println(UserPass[i])
			if len(Proxy) == 4 {
				AddAuthCookie([]byte(UserPass[i]), ProxyC(Proxy))
			} else {
				AddAuthCookie([]byte(UserPass[i]), ProxyC(Proxy))
			}
		}
	}
}
func AddAuthCookie(Account []byte, Client *http.Client) {
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
	response, _ := Client.Do(request)
	body, _ := io.ReadAll(response.Body)
	format := strings.Split(string(body), "\"")
	if response.StatusCode == 200 {
		log.Printf("[FriendSpam] | %v | Logged in as: %v", response.Status, format[11])
		authcookies = append(authcookies, response.Cookies()[0].Value)
	} else {
		log.Println(response.Status, string(body))
	}
}

func FriendRequest(authcookie string, Client *http.Client) {
	request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/user/"+UserID+"/friendRequest", nil)
	request.Header = http.Header{
		"Cookie":     {"apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth=" + authcookie},
		"User-Agent": {"VRC.Core.BestHTTP"},
	}
	response, _ := Client.Do(request)
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.Status, string(body))
}
