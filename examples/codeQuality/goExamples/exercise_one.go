package go_unit_test_bootcamp

import (
	"sort"
	//"fmt"
)

func FindMissingDrone(droneIds []int) int {

	result := 0

	sort.Ints(droneIds)
	//fmt.Println("sorted:", droneIds)

	for i := 0; i <= len(droneIds)-1; i++ {
		if len(droneIds) == 1 {
			result = droneIds[0]
			break
		} else if droneIds[i] != droneIds[i+1] {
			result = droneIds[i]
			break
		} else {
			i++
		}
	}

	return result

}
