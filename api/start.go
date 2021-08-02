package api

import (
	"checkerboard/hander"
	"checkerboard/model"
	"checkerboard/mq"
)

func Start() (ids string) {
	return hander.Start()
}

func Play(ids string,x,y,player int,back model.CallBack)  {
	mq.ProducePlayerMessage(model.PlayerMessage{
		Pos: struct {
			X int
			Y int
		}{
			x,y,

		},
		Player: player,
		Ids:    ids,
		Back:   back,
	})
}