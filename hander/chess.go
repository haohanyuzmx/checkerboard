package hander

import (
	"checkerboard/model"
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

var manger Manger

func Init() {
	manger=Manger{
		AllBoard:    make(map[string]*Chess),
	}
}


type Manger struct {
	AllBoard  map[string]*Chess
	BoardLock sync.RWMutex

	AllCallBack sync.Map
}

func saveChess(chess *Chess, ids string) {
	manger.BoardLock.Lock()
	manger.AllBoard[ids] = chess
	manger.BoardLock.Unlock()
}
func getChess(ids string) *Chess {
	manger.BoardLock.RLock()
	chess, ok := manger.AllBoard[ids]
	if !ok {
		return nil
	}
	manger.BoardLock.RUnlock()
	return chess
}
func deleteChess(ids string) {
	manger.BoardLock.Lock()
	delete(manger.AllBoard, ids)
	manger.BoardLock.Unlock()
}
func SaveCallBack(ids string, back model.CallBack) {
	manger.AllCallBack.Store(ids, back)
}
func getCallBack(ids string) model.CallBack {
	load, _ := manger.AllCallBack.Load(ids)
	return load.(model.CallBack)
}
func deletCallBack() {

}

type Chess struct {
	sync.Mutex
	Board  [15][15]int
	Step   int
	Player int //黑棋是1，白棋2
}

func Start() (ids string) {
	ids = uuid.NewV4().String()
	c := Chess{
		Board:  [15][15]int{},
		Step:   0,
		Player: 0,
	}//pool
	saveChess(&c, ids)
	return ids
}

func Check(message model.PlayerMessage) (step int, err error) {
	chess := getChess(message.Ids)
	if chess != nil {
		return 0,errors.New("没有这局棋")
	}

	chess.Lock()
	defer chess.Unlock()

	if chess.Player != message.Player {
		err=errors.New("不是你的回合")
		message.Back(0, err)
		return 0,err
	}
	x := message.Pos.X
	y := message.Pos.Y

	i := chess.Board[x][y]
	if i != 0 {
		err=errors.New("已经下了")
		message.Back(0, err)
		return 0,err
	}

	chess.Board[x][y] = chess.Player
	winner := result(chess.Board, x, y)
	chess.Step++
	if winner == 0 {
		if chess.Player == 1 {
			chess.Player = 2
		} else {
			chess.Player = 1
		}
	} else {
		message.Back(winner, nil)
		deleteChess(message.Ids)
	}
	return chess.Step, nil
}

func result(board [15][15]int, xofPlay, yofPlay int) (winner int) {
	var (
		num1 = 0
		num2 = 0
		x    = 1
		y    = 1
	)
	player := board[xofPlay][yofPlay]
	for player == board[xofPlay+x+num1][yofPlay+y+num1] {
		num1++
	}
	for player == board[xofPlay-x-num2][yofPlay-y-num2] {
		num2++
	}
	if num1+num2 >= 5 {
		return player
	}
	//...
	return 0
}
