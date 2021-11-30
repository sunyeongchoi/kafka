package main

import (
	"fmt"
	"kafka-golang/Basic"
)


func main() {
	fmt.Println("Kafka Consumer Example")
	Basic.Consumer()

	//fmt.Println("kafka Producer Example")
	//Basic.Producer()
}