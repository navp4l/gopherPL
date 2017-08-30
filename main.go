package main

import (
	"fmt"
)

//func main() {
//
//	reactors := make([]float64, 5, 100 )
//	reactors[0] = 100.0
//	reactors[1] = 110.0
//	reactors[2] = 120.0
//	reactors[3] = 130.0
//	reactors[4] = 140.0
//
//	totalCapacity := reactors[0] + reactors [1] + reactors[2] + reactors [3] + reactors[4]
//
//	gridLoad := 75.
//
//	utilization := gridLoad / totalCapacity
//
//	fmt.Printf("%-20s%.0f\n", "Total capacity is:", totalCapacity)
//
//	fmt.Printf("%-20s%.0f\n", "Current grid load is:", gridLoad)
//
//	fmt.Printf("%-20s%.2f%%\n", "Current Utilization is:", utilization * 100)
//3
//}

func main() {

	reactors := make([]float64, 5, 100 )
	reactors[0] = 100.0
	reactors[1] = 110.0
	reactors[2] = 120.0
	reactors[3] = 130.0
	reactors[4] = 140.0

	activeReactors := []int {0,1}

	gridLoad := 175.0

	fmt.Println("1) Generate Power Plan Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please enter a valid option")

	var option string

	fmt.Scanln(&option)

	switch {

	case option == "1":
		for index, capacity := range reactors {
			fmt.Printf("Plant %d capacity: %.0f\n", index, capacity)
		}
	case option == "2":
		activeCapacity := 0.0
		for plantId := range activeReactors {
			activeCapacity += reactors[plantId]
		}

		fmt.Printf("%-20s%.0f\n", "Total capacity is:", activeCapacity)

		fmt.Printf("%-20s%.0f\n", "Current grid load is:", gridLoad)

		fmt.Printf("%-20s%.2f%%\n", "Current Utilization is:", gridLoad / activeCapacity * 100)

	default:
		fmt.Println("Please enter a valid option!")

	}

}