package main

import "fmt"

func main() {
	//11111
	SecToNano := int64(1000000000)
	StepCostTime := int64(float32( int64(300)*SecToNano) )
	fmt.Printf("cost time: %v",StepCostTime)
}
