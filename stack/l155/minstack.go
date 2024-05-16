package l155

import "fmt"

type MinStack struct {
	Numbers []int
	MapI    map[int]map[int]bool
	MinHeap []int
}

func Constructor() MinStack {
	return MinStack{
		Numbers: make([]int, 0),
		MapI:    make(map[int]map[int]bool, 0),
		MinHeap: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.Numbers = append(this.Numbers, val)
	this.MinHeap = append(this.MinHeap, val)
	if this.MapI[val] == nil {
		this.MapI[val] = make(map[int]bool, 0)
	}
	this.MapI[val][len(this.MinHeap)-1] = true
	i := len(this.MinHeap) - 1
	for i > 0 {
		pi := (i - 1) / 2
		if this.MinHeap[i] < this.MinHeap[pi] {
			imap := this.MapI[this.MinHeap[i]]
			pimap := this.MapI[this.MinHeap[pi]]
			delete(imap, i)
			delete(pimap, pi)
			pimap[i] = true
			imap[pi] = true
			this.MinHeap[i], this.MinHeap[pi] = this.MinHeap[pi], this.MinHeap[i]
		} else {
			break
		}
		i = pi
	}
	fmt.Printf("Pushed: stack=%v heap=%v map=%v\n", this.Numbers, this.MinHeap, this.MapI)
}

func (this *MinStack) Pop() {
	valD := this.Numbers[len(this.Numbers)-1]
	this.Numbers = this.Numbers[:len(this.Numbers)-1]
	dMap := this.MapI[valD]
	var iD int
	for k := range dMap {
		iD = k
		break
	}
	//fmt.Printf("PL1: dMap=%v valD=%v\n",dMap,valD)
	delete(dMap, iD)
	if len(dMap) == 0 {
		delete(this.MapI, valD)
	}
	if iD == len(this.MinHeap)-1 {
		this.MinHeap = this.MinHeap[:len(this.MinHeap)-1]
		//fmt.Printf("Popped top: stack=%v heap=%v map=%v num=%v\n",this.Numbers,this.MinHeap,this.MapI,valD)
		return
	}

	uMap := this.MapI[this.MinHeap[len(this.MinHeap)-1]]
	//fmt.Printf("PL2: dMap=%v uMap=%v valD=%v\n",dMap,uMap,valD)
	delete(uMap, len(this.MinHeap)-1)
	uMap[iD] = true
	this.MinHeap[iD], this.MinHeap[len(this.MinHeap)-1] = this.MinHeap[len(this.MinHeap)-1], this.MinHeap[iD]
	this.MinHeap = this.MinHeap[:len(this.MinHeap)-1]
	for {
		lc := 2*iD + 1
		minI := iD
		if lc < len(this.MinHeap) && this.MinHeap[lc] < this.MinHeap[minI] {
			minI = lc
		}
		rc := 2*iD + 2
		if rc < len(this.MinHeap) && this.MinHeap[rc] < this.MinHeap[minI] {
			minI = rc
		}
		if minI == iD {
			break
		}
		replM := this.MapI[this.MinHeap[minI]]
		delete(uMap, iD)
		uMap[minI] = true
		delete(replM, minI)
		replM[iD] = true
		this.MinHeap[iD], this.MinHeap[minI] = this.MinHeap[minI], this.MinHeap[iD]
		iD = minI
		uMap = this.MapI[this.MinHeap[iD]]
	}
	//fmt.Printf("Popped: stack=%v heap=%v map=%v num=%v\n",this.Numbers,this.MinHeap,this.MapI,valD)
}

func (this *MinStack) Top() int {
	val := this.Numbers[len(this.Numbers)-1]
	return val
}

func (this *MinStack) GetMin() int {
	return this.MinHeap[0]
}
