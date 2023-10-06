package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"math"
	"math/rand"

	"time"
	// "github.com/eiannone/keyboard"
	"os"
)

func main() {

	var b Ball
	var lr, rr Racket
	var f Field
	l := 45
	w := 15
	//b.createBall(float64(l)/2.0, float64(w)/2.0, rand.Float64()*0.4+0.4, 0.6*randOneMinusOne())
	b.createBall(15, 8, 0.4, 0.6)
	lr.createRacket(1, 2, 3, true)
	//rr.createRacket(l-1, w/2, 3, true)
	//lr.createRacket(1, 1, w, true)
	rr.createRacket(l, 1, w, true)
	f.createField(l, w, &b, &lr, &rr)

	f.upDatePosition()
	f.show(0)

	var nn NeuronNetwork
	nn.newNN()
	nn.executeNN()

	//test := ""

	var i uint32
	for true {
		if i != f.point {
			i = f.point
			fmt.Println("Your points:", f.point)

		}
		//fmt.Println(i)

		err := f.move()
		if err != nil {
			break
		}

		f.upDatePosition()

		//time.Sleep(time.Second * 7)
		//writeTest(&f, &lr, &b)
		//js, _ := json.Marshal(simpleAlgoritm(&f, &lr, &b))
		//test += strconv.Itoa(lr.y+1) + ":" + strconv.Itoa(int(math.Round(b.vY+b.y))) + ":" + string(js) + "\n"
		//fmt.Print(strconv.Itoa(lr.y+1) + ":" + strconv.Itoa(int(math.Round(b.vY+b.y))) + ":" + string(js) + "\n")
		nn.pushTheData(int8(lr.y+1), int8(math.Round(b.vY+b.y)))
		//fmt.Println(nn.finalNeuron.out)
		//resSimple := simpleAlgoritm(&f, &lr, &b)
		if nn.finalNeuron.out < 0 {
			f.moveRacketVertical(false, f.lr)
		} else {
			f.moveRacketVertical(true, f.lr)
		}
		// if resSimple[0] == 1 {
		// 	f.moveRacketVertical(false, f.lr)
		// } else if resSimple[1] == 1 {
		// 	f.moveRacketVertical(true, f.lr)
		// }
		f.show(0)

	}
	fmt.Println("Game over", "    score: ", f.point)
	//createTest(test)
	time.Sleep(time.Second)
	fmt.Scan(&w)

}

func simpleAlgoritm(f *Field, r *Racket, b *Ball) []float64 {

	centerRacketY := r.y + int(r.size/2)
	upY := centerRacketY - 1
	downY := centerRacketY + 1

	//lenghtCenter := math.Sqrt(math.Pow(float64(centerRacketY)-(b.vY+b.y), 2) + math.Pow(float64(r.x)-(b.vX+b.x), 2))
	lenghtUp := math.Sqrt(math.Pow(float64(upY)-(b.vY+b.y), 2) + math.Pow(float64(r.x)-(b.vX+b.x), 2))
	lenghtDown := math.Sqrt(math.Pow(float64(downY)-(b.vY+b.y), 2) + math.Pow(float64(r.x)-(b.vX+b.x), 2))

	if lenghtUp < lenghtDown {
		return []float64{1.0, 0.0}
	}
	return []float64{0.0, 1.0}

	// if lenghtUp < lenghtCenter && lenghtUp < lenghtDown {
	// 	//fmt.Println("вверх")
	// 	//f.moveRacketVertical(false, r)
	// 	return []float64{1.0, 0.0, 0.0}
	// }
	// if lenghtDown < lenghtCenter && lenghtDown < lenghtUp {
	// 	//fmt.Println("вниз")
	// 	//f.moveRacketVertical(true, r)
	// 	return []float64{0.0, 0.0, 1.0}
	// }
	// return []float64{1.0, 1.0, 0}
}

func randOneMinusOne() float64 {
	if rand.Float64()-0.5 >= 0 {
		return 1
	} else {
		return -1
	}
}

func createTest(s string) {

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	file.WriteString(s)
	defer file.Close()
}

func writeTest(f *Field, r *Racket, b *Ball) {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	js, _ := json.Marshal(simpleAlgoritm(f, r, b))
	str := strconv.Itoa(r.y+1) + ":" + strconv.Itoa(int(math.Round(b.vY+b.y))) + ":" + string(js)
	//fmt.Println(str)
	file.WriteString(str)
	file.Close()

}
