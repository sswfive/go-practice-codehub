package main

import (
	"fmt"
	_ "io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "C"
	case score < 80:
		g = "B"
	case score < 90:
		g = "A"
	}
	return g
}



func main(){
	//const filename = "abc.txt"

	//contents, err := ioutil.ReadFile(filename)
	//if err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s\n", contents)
	//}

	//if contents, err := ioutil.ReadFile(filename); err != nil{
	//	fmt.Println(err)
	//}else{
	//	fmt.Printf("%s\n", contents)
	//}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(85),
		grade(98),
		grade(100),
		//grade(101),
		//grade(-1),
		)
}