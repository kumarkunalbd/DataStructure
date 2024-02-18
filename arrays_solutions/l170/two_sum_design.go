package l170

type TwoSum struct {
	MapFreq map[int]int
}

func Constructor() TwoSum {
	return TwoSum{MapFreq: make(map[int]int, 0)}
}

func (this *TwoSum) Add(number int) {
	if freq, ok := this.MapFreq[number]; ok {
		this.MapFreq[number] = freq + 1
	} else {
		this.MapFreq[number] = 1
	}
}

func (this *TwoSum) Find(value int) bool {
	for num, freq := range this.MapFreq {
		diffN := value - num
		if _, ok := this.MapFreq[diffN]; ok {
			if diffN == num && freq < 2 {
				continue
			}
			return true
		}
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
