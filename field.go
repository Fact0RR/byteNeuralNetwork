package main

import (
	"errors"
	"fmt"
	"math"
)

type Field struct {
	length, width int
	fieldSlice    [][]int
	lr, rr        *Racket
	ball          *Ball
	point         uint32
}

func (f *Field) createField(length, width int, ball *Ball, lr, rr *Racket) {

	f.lr = lr
	f.rr = rr
	f.ball = ball
	f.length = length
	f.width = width
	f.fieldSlice = make([][]int, width+2)
	for i := range f.fieldSlice {
		f.fieldSlice[i] = make([]int, length+2)
	}
	for i := 0; i < len(f.fieldSlice); i++ {
		for j := 0; j < len(f.fieldSlice[i]); j++ {
			if i == 0 || i == len(f.fieldSlice)-1 {
				f.fieldSlice[i][j] = 1
			}
		}
	}
	f.fieldSlice[0][1] = 4
	f.fieldSlice[len(f.fieldSlice)-1][1] = 4
	f.fieldSlice[0][len(f.fieldSlice[0])-2] = 4
	f.fieldSlice[len(f.fieldSlice)-1][len(f.fieldSlice[0])-2] = 4
}

func (f *Field) upDatePosition() {

	f.clearRacketsAndBall()
	for i := 0; i < f.lr.size; i++ {
		if f.lr.rotation {
			f.fieldSlice[f.lr.y+i][f.lr.x] = 2
		} else {
			f.fieldSlice[f.lr.y][f.lr.x+i] = 2
		}
	}
	for i := 0; i < f.rr.size; i++ {
		if f.rr.rotation {
			f.fieldSlice[f.rr.y+i][f.rr.x] = 2
		} else {
			f.fieldSlice[f.rr.y][f.rr.x+i] = 2
		}
	}

	f.fieldSlice[int(math.Round(f.ball.y))][int(math.Round(f.ball.x))] = 3

}

func (f *Field) clearRacketsAndBall() {

	for i := 1; i < len(f.fieldSlice)-1; i++ {
		for j := 1; j < len(f.fieldSlice[i])-1; j++ {
			if f.fieldSlice[i][j] != 1 {
				f.fieldSlice[i][j] = 0
			}
		}
	}
}

func (f *Field) show(nnscore uint32) {
	//slice := make([]byte,len(f.fieldSlice)*len(f.fieldSlice[0]))

	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Your points:", f.point, "NNscore: ", nnscore)

	var str string

	for i := 0; i < len(f.fieldSlice); i++ {
		for j := 0; j < len(f.fieldSlice[i]); j++ {
			if f.fieldSlice[i][j] == 1 {
				//fmt.Print("- ")
				str += "- "
			} else if f.fieldSlice[i][j] == 2 {
				//fmt.Print("| ")
				str += "| "
			} else if f.fieldSlice[i][j] == 3 {
				//fmt.Print("* ")
				str += "* "
			} else if f.fieldSlice[i][j] == 4 {
				//fmt.Print("+ ")
				str += "+ "
			} else {
				//fmt.Print(". ")
				str += "  "
			}
			//fmt.Print(f.fieldSlice[i][j], " ")
		}
		//fmt.Println()
		str += "\n"
	}
	fmt.Println(str)
}

func (f *Field) move() error {
	f.upDatePosition()
	f.ball.move()

	i := int(math.Round(f.ball.y + f.ball.vY))
	j := int(math.Round(f.ball.x + f.ball.vX))
	if j < 0 {
		return errors.New("выход за границу")
	}

	//fmt.Println(f.fieldSlice[i][j], " j= ", j, "i = ", i, " bx= ", f.ball.x, " by= ", f.ball.y, " vx =", f.ball.vX, " vy =", f.ball.vY)
	if f.fieldSlice[i][j] == 4 {
		if f.ball.vX < 0 {
			f.point++
		}
		f.ball.vY = -f.ball.vY
		f.ball.vX = -f.ball.vX
		f.ball.move()
	}
	if f.fieldSlice[i][j] == 2 {
		if f.ball.vX < 0 {
			f.point++
		}
		f.ball.vX = -f.ball.vX
		f.ball.move()
	}
	if f.fieldSlice[i][j] == 1 {
		f.ball.vY = -f.ball.vY
		f.ball.move()
	}
	return nil
}

func (f *Field) moveRacketVertical(x bool, r *Racket) {
	//x==true = движение вниз
	//x==false = движение вверх
	if r.y > 1 && r.y+r.size-1 < f.width {
		//fmt.Println("Двигаться можно")
		if x {
			r.y = r.y + 1
		}
		if !x {
			r.y = r.y - 1
		}
	} else if r.y == 1 {
		//fmt.Println("Двигаться можно только вниз")
		if x {
			r.y = r.y + 1
		}
	} else if r.y+r.size-1 == f.width {
		//fmt.Println("Двигаться можно только вверх")
		if !x {
			r.y = r.y - 1
		}
	} else {
		//fmt.Println("Двигаться нельзя")
	}
}
