package Problem0475

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	if len(houses) == 0 {
		return 0
	}

	res := 0

	sort.Ints(houses)
	sort.Ints(heaters)

	iHeater := sort.SearchInts(heaters, houses[0])
	iHouse := 0

	for iHouse < len(houses) {
		for iHeater < len(heaters) && houses[iHouse] > heaters[iHeater] {
			iHeater++
		}

		// if iHeater == len(heaters) {
		// 	return max(res, houses[len(houses)-1]-heaters[iHeater-1])
		// }

		left := 1<<31 - 1
		if 0 <= iHeater-1 {
			left = houses[iHouse] - heaters[iHeater-1]
		}
		right := heaters[iHeater] - houses[iHouse]
		res = max(res, min(left, right))

		for iHouse < len(houses) &&
			((0 <= iHeater-1 && houses[iHouse] <= heaters[iHeater-1]+res) ||
				(0 == iHeater)) {
			iHouse++
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
