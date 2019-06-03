package cg

import (
	"errors"
	"encoding/json"

	"go-learn/cgss/src/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterServer)AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil{
		return err
	}

	resp, err := client.Call("addplayer", string(b))
	if err != nil && resp.Code == "200"{
		return nil
	}
	return err
}


