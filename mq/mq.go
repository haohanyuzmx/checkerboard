package mq

import (
	"checkerboard/db"
	"checkerboard/hander"
	"checkerboard/model"
)

var (
	messageOfPlay chan model.PlayerMessage
	messageOfResult chan model.ChessIn
)


func Init() {
	messageOfPlay=make(chan model.PlayerMessage,1024)
	messageOfResult=make(chan model.ChessIn,512)
}


func ProducePlayerMessage(message model.PlayerMessage)  {
	messageOfPlay<-message
}

func ConsumePlayerMessage() {
	for message := range messageOfPlay {
		step, err := hander.Check(message)
		if err != nil {
			continue
		}

		ProduceResultMessage(model.ChessIn{
			Ids:    message.Ids,
			X:      message.Pos.X,
			Y:      message.Pos.Y,
			Step:   step,
			Player: message.Player,
		})
	}
}

func ProduceResultMessage(in model.ChessIn)  {
	messageOfResult<-in
}

func ConsumeResultMessage()  {
	for result := range messageOfResult {
		db.Creat(result)
	}
}