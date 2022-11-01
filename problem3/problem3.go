package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	// your code here
	var l float64 = float64(len(s.score))
	var jumlah float64
	for _, v := range s.score {
		jumlah = jumlah + float64(v)
	}
	res := jumlah / l
	return res
}

func (s Student) Min() (min int, name string) {
	// your code here
	min = s.score[0]
	var index_nama int
	for _, v := range s.score {
		if v < min {
			min = v
		}
	}
	for i := 0; i < len(s.score); i++ {
		if s.score[i] == min {
			index_nama = i
		}

	}
	name = s.name[index_nama]
	return min, name

}

func (s Student) Max() (max int, name string) {
	// your code here
	max = s.score[0]
	var index_nama int
	for _, v := range s.score {
		if v > max {
			max = v
		}
	}
	for i := 0; i < len(s.score); i++ {
		if s.score[i] == max {
			index_nama = i
		}

	}
	name = s.name[index_nama]
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 6; i++ {
		var name string
		fmt.Print("Input " + fmt.Sprint(i) + " Student's Name : ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score : ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Println("\n\nAverage Score Students is", a.Average())
	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is : "+nameMax+" (", scoreMax, ")")
	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is : "+nameMin+" (", scoreMin, ")")
}
