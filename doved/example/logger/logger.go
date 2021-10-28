package main

import "log"

func init() {

}

func main() {
	log.Println("1234")
	//log.Fatalln("1234")
	log.Panicln("1234")
}
