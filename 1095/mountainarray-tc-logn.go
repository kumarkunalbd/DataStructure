/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

func findInMountainArray(target int, mountainArr *MountainArray) int {
	s := 0
	e := mountainArr.length() - 1
	peak := math.MinInt
	mapMountain := map[int]int{}
	for s <= e {
		mid := (s + e) / 2
		mapMountain[mid] = mountainArr.get(mid)
		mapMountain[mid+1] = mountainArr.get(mid + 1)
		if mapMountain[mid] > mapMountain[mid+1] {
			peak = mid
			e = mid - 1
		} else {
			s = mid + 1
		}
	}
	fmt.Printf("l1: peak=%v map=%v\n", peak, mapMountain)

	// search in 1st Array
	s = 0
	e = peak
	for s <= e {
		mid := (s + e) / 2
		//var midNumber int
		if _, ok := mapMountain[mid]; !ok {
			mapMountain[mid] = mountainArr.get(mid)
		}
		if mapMountain[mid] == target {
			return mid
		}

		if mapMountain[mid] > target {
			e = mid - 1
		} else {
			s = mid + 1
		}
	}

	// search in 2nd array
	s = peak + 1
	e = mountainArr.length() - 1
	for s <= e {
		mid := (s + e) / 2
		if _, ok := mapMountain[mid]; !ok {
			mapMountain[mid] = mountainArr.get(mid)
		}
		if mapMountain[mid] == target {
			return mid
		}
		if mapMountain[mid] > target {
			s = mid + 1
		} else {
			e = mid - 1
		}
		//fmt.Printf("l3: s=%v e=%v\n",s,e)
	}
	return -1
}

