package server

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	ResponseError string = `{"response":"error", "error":"%s"}`
	ResponseOk    string = `{"response":"ok"}`

	ResponseGameCreated string = `{"response":"ok", "gameId":%d}`
	ResponseListGame    string = `{"response":"ok", "games":%s}`

	RequestPlay string = `{"request":"play", "param":%s}`

	RequestListGame   string = "listGame"
	RequestCreateGame string = "createGame"
	RequestJoinGame   string = "joinGame"
)

func getErrorResponse(err string) []byte {
	log.Println(err)
	r := fmt.Sprintf(ResponseError, err)
	return []byte(r)
}

func getOkResponse() []byte {
	return []byte(ResponseOk)
}

func getGameCreatedResponse(gameId int) []byte {
	r := fmt.Sprintf(ResponseGameCreated, gameId)
	return []byte(r)
}

func getListGameResponse(games []GameDescription) []byte {
	g, err := json.Marshal(games)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return []byte(fmt.Sprintf(ResponseListGame, g))
}

func getPlayRequest(bot *Bot, squares []Square) []byte {
	r, err := json.Marshal(PlayParam{bot, squares})
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return []byte(fmt.Sprintf(RequestPlay, r))
}

type PlayParam struct {
	Bot     *Bot     `json:"bot"`
	Squares []Square `json:"squares"`
}

type GameDescription struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Bots []string `json:"bots"`
}

type StandardRequest struct {
	Request string                 `json:"request"`
	Param   map[string]interface{} `json:"param"`
}
