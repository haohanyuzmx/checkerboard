package model

import "gorm.io/gorm"

type ChessIn struct {
	gorm.Model
	Ids    string
	X      int
	Y      int
	Step   int
	Player int
}

type PlayerMessage struct {
	Pos struct{
		X int
		Y int
	}
	Player int
	Ids string
	Back CallBack
}

type CallBack func(result int,error error)