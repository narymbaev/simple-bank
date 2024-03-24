package main

import "fmt"

type TxString string

type ts struct {
	title TxString
	qwe string
}

func main(){
	a := ts{}

	fmt.Println("title:", a.title, " qwe:", a.qwe)

}
