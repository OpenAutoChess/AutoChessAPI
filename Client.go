package main

import (
	"encoding/json"
	"errors"
	socketio "github.com/googollee/go-socket.io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Rank struct {
	display string
	value int
}


type User struct {
	Id int
	Avatar string
	UserName string
	Mmr int
	Rank Rank
}

type Client struct {
	AuthToken string
	User User
}

func NewClient(s socketio.Conn) (Client, error) {
	cookies := parseCookies(s.RemoteHeader().Get("Cookie"))
	authToken := cookies["AuthToken"]

	user, err := getUserData(authToken)
	if err != nil {
		return Client{}, err
	}

	return Client{authToken, user}, nil
}


func parseCookies(cookiesRaw string) map[string]string {

	cookies := strings.Split(cookiesRaw, ";")
	cookiesMap := make(map[string]string)

	for i := 0; i < len(cookies); i+=2 {
		cookie := strings.Split(cookies[i], "=")
		cookiesMap[strings.TrimSpace(cookie[0])] = strings.TrimSpace(cookie[1])
	}

	return cookiesMap
}

func getUserData(authToken string) (User, error) {
	if len(authToken) == 0 {
		return User{}, errors.New("empty AuthToken")
	}

	client := http.Client{}

	request, err := http.NewRequest("GET", "http://auth.autochess.kz/api/users/profile/", nil)
	request.Header.Set("Authorization", authToken)
	if err != nil {
		return User{}, err
	}
	response, err := client.Do(request)
	if err != nil {
		return User{}, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return User{}, err
	}

	user := make(map[string]interface{})
	err = json.Unmarshal(body, &user)
	if err != nil {
		return User{}, err
	}

	return NewUser(user), nil
}

func NewUser(user map[string]interface{}) User {
	return User{int(user["id"].(float64)),user["avatar"].(string), user["user_name"].(string), int(user["points"].(float64)), Rank{user["rank"].(map[string]interface{})["display"].(string), int(user["rank"].(map[string]interface{})["value"].(float64))}}
}