package server

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	ResponseError       string = `{"response":"error", "error":"%s"}`
	ResponseOk          string = `{"response":"ok"}`
	ResponseGameCreated string = `{"response":"ok", "gameId":%d}`
	RequestPlay         string = `{"request":"play", "param":%s}`

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

func getPlayRequest(bot *Bot, squares []Square) ([]byte, error) {
	r, err := json.Marshal(PlayParam{bot, squares})
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return []byte(fmt.Sprintf(RequestPlay, r)), nil
}

type PlayParam struct {
	Bot     *Bot     `json:"bot"`
	Squares []Square `json:"squares"`
}

type StandardRequest struct {
	Request string                 `json:"request"`
	Param   map[string]interface{} `json:"param"`
}
