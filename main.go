package main

import (
	"fmt"
	"petproj/feature1"
	"petproj/feature2"
	simpleconnection "petproj/feature_postgres/simple_connection"
)

func main() {
	fmt.Println("test main")
	feature1.Feature1()
	feature2.Feature2()
	simpleconnection.CheckConnection()
}
