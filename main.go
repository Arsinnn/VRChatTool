package main

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	sha := sha256.New()
	sha.Write([]byte("s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2s2"))
	fmt.Println(sha.Sum([]byte("s")))
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
	fmt.Println("1: Request spam UserID\n2: World spoof\n3: Quest Spoof\n4: Request spam multi")
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
		WorldSpoof()
	case "3":
		QuestSpoof()
	case "4":
		RequestSpamMulti()
	default:
		fmt.Println("Pick a valid option")
		time.Sleep(3 * time.Second)
		Start()
	}
}

func RequestSpam(UserID string) {
	client := http.Client{}
	for i := 0; i < 25; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/requestInvite/"+UserID, nil)
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

func WorldSpoof() {
	fmt.Println("Enter WorldID")
	var WorldID string
	fmt.Scanln(&WorldID)
	client := http.Client{}
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/travel/"+WorldID+":Joe~region(jp)/token?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat", nil)
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("X-MacAddress", "e8d99139c531a78d370d341c95d2b2f0fc68a330")
	request.Header.Add("X-Client-Version", "2022.2.2-1213--Release")
	request.Header.Add("X-Platform", "standalonewindows")
	request.Header.Add("X-Unity-Version", "2019.4.31f1")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", "vrchat.com")
	request.Header.Add("Host", "api.vrchat.cloud")
	request.Header.Add("TE", "identity")
	request.Header.Add("User-Agent", "VRC.Core.BestHTTP")
	request.Header.Add("Cookie", "auth="+authcookie+"; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth=")
	request.Header.Add("Accept-Encoding", "identity")
	respose, _ := client.Do(request)
	fmt.Println(respose.Status)
	Start()
}

func QuestSpoof() {
	client := http.Client{}
	request, _ := http.NewRequest("GET", "https://api.vrchat.cloud/api/1/worlds/recent?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat&n=30&order=descending&offset=0&tag=system_approved&releaseStatus=public&maxUnityVersion=2019.4.31f1&minUnityVersion=5.5.0f1&maxAssetVersion=4&minAssetVersion=0&platform=android", nil)
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("X-MacAddress", "4263a2add68d70bc167a626a27db4ea90c279aef")
	request.Header.Add("X-Client-Version", "2022.2.1p4-1205--Release")
	request.Header.Add("X-Platform", "android")
	request.Header.Add("X-Unity-Version", "2019.4.31f1")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Origin", "vrchat.com")
	request.Header.Add("Host", "api.vrchat.cloud")
	request.Header.Add("TE", "identity")
	request.Header.Add("User-Agent", "VRC.Core.BestHTTP")
	request.Header.Add("Cookie", "auth= "+authcookie+"; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth=")
	response, _ := client.Do(request)
	fmt.Println(response.Status)
	Start()
}

func RequestSpamMulti() {
	Auth, _ := os.ReadFile("Authcookies.txt")
	split := strings.Split(string(Auth), ":")
	fmt.Println("Enter UserID to request spam")
	var UserID string
	fmt.Scanln(&UserID)
	for i := 0; i < len(split); i++ {
		go RequestSpam2(UserID, split[i])
	}
	Start()
}
func RequestSpam2(UserID string, AuthCookie string) {
	client := http.Client{}
	for i := 0; i < 140; i++ {
		request, _ := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/requestInvite/"+UserID, nil)
		request.Header.Add("Cookie", "apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; auth="+AuthCookie)
		request.Header.Add("X-Client-Version", "2022.2.2-1213--Release")
		request.Header.Add("X-Platform", "standalonewindows")
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("User-Agent", "Transmtn-Pipeline")
		request.Header.Add("Host", "api.vrchat.cloud")
		response, _ := client.Do(request)
		fmt.Println(AuthCookie, response.Status)
	}
}

/*try
  {
      HttpWebRequest request = (HttpWebRequest)WebRequest.Create("https://api.vrchat.cloud/api/1/worlds/recent?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&organization=vrchat&n=30&order=descending&offset=0&tag=system_approved&releaseStatus=public&maxUnityVersion=2019.4.31f1&minUnityVersion=5.5.0f1&maxAssetVersion=4&minAssetVersion=0&platform=android");
      request.Method = "GET";
      request.Headers["X-Requested-With"] = "XMLHttpRequest";
      request.Headers["X-MacAddress"] = "4263a2add68d70bc167a626a27db4ea90c279aef";
      request.Headers["X-Client-Version"] = "2022.2.1p4-1205--Release";
      request.Headers["X-Platform"] = "android";
      request.Headers["X-Unity-Version"] = "2019.4.31f1";
      request.ContentType = "application/x-www-form-urlencoded";
      request.Headers["Origin"] = "vrchat.com";
      request.Host = "api.vrchat.cloud";
      request.Headers["TE"] = "identity";
      request.UserAgent = "VRC.Core.BestHTTP";
      request.Headers["Cookie"] = "auth=authcookie_780a62ea-430f-424f-b631-1677acc0924e; apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26; twoFactorAuth=";
      request.GetResponse();
  }
  catch (Exception e)
  {
      Console.WriteLine(e);
  }*/
