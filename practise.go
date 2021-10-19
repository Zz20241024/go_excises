package main

import (
	"fmt"
	"math/rand"
)

//从小到大
type SF struct {
	SArray []int
}

//冒泡
func (sf *SF) Bubble() []int {
	length := len(sf.SArray)
	for o := 0; o < length; o++ {
		for i := 0; i < length-o-1; i++ {
			if sf.SArray[i] > sf.SArray[i+1] {
				sf.SArray[i], sf.SArray[i+1] = sf.SArray[i+1], sf.SArray[i]
			}
		}

	}
	return sf.SArray
}

//选择
func (sf *SF) Choice() []int {
	length := len(sf.SArray)
	for o := 0; o < length; o++ {
		mimum := o
		for i := o; i < length; i++ {
			if sf.SArray[mimum] > sf.SArray[i] {
				mimum = i
			}
		}
		sf.SArray[o], sf.SArray[mimum] = sf.SArray[mimum], sf.SArray[o]

	}
	return sf.SArray
}

//插入
func (sf *SF) Insert() []int {
	length := len(sf.SArray)
	if length == 1 {
		return sf.SArray
	}
	for o := 1; o < length; o++ {
		for i := o - 1; i <= o; i-- {
			if sf.SArray[i] > sf.SArray[i+1] {
				sf.SArray[i], sf.SArray[i+1] = sf.SArray[i+1], sf.SArray[i]
			} else {
				break
			}
		}

	}
	return sf.SArray
}

func (sf *SF) pop(index int, sl []int) (int, []int) {
	a := sl[index]
	return a, append(sl[:index], sl[index+1:]...)
}

//快排
func (sf *SF) QuickSort(SArray []int) []int {
	length := len(SArray)
	if length < 2 {
		return SArray
	}
	middle := int(length / 2)
	value, SList := sf.pop(middle, SArray)
	leftList := []int{}
	rightList := []int{}
	for _, i := range SList {
		if i < value {
			leftList = append(leftList, i)
		} else {
			rightList = append(rightList, i)
		}
	}
	leftListIncursion := sf.QuickSort(leftList)
	rightListIncursion := sf.QuickSort(rightList)
	return append(append(leftListIncursion, value), rightListIncursion...)
}

//合并
func (sf *SF) Merge(SArray []int) []int {
	length := len(SArray)
	if length < 2 {
		return SArray
	}
	middle := int(length / 2)
	leftListIncursion := sf.Merge(SArray[:middle])
	rightListIncursion := sf.Merge(SArray[middle:])
	return sf.mergeSort(leftListIncursion, rightListIncursion)
}

func (sf *SF) mergeSort(lArray []int, rArray []int) []int {
	llen := len(lArray)
	ls := 0
	rs := 0
	rlen := len(rArray)
	new := []int{}
	for {
		if ls == llen {
			new = append(new, rArray[rs:]...)
			break
		}
		if rs == rlen {
			new = append(new, lArray[ls:]...)
			break
		}
		if lArray[ls] < rArray[rs] {
			new = append(new, lArray[ls])
			ls++
		} else {
			new = append(new, rArray[rs])
			rs++
		}

	}
	return new

}

//heap 堆排序
func (sf *SF) Heap(SArray []int) []int {
	length := len(SArray)
	var v int
	new := []int{}
	for i := 0; i < length; i++ {
		SArray = sf.heapSort(SArray)
		v, SArray = sf.pop(0, SArray)
		new = append(new, v)
	}
	return new
}

//最大的在index=0
func (sf *SF) heapSort(SArray []int) []int {
	length := len(SArray)
	if length < 2 {
		return SArray
	}
	for i := int((length - 2) / 2); i >= 0; i-- {
		left := 2*i + 1
		right := 2*i + 2
		if right >= length {
			right = left
		}
		if SArray[left] > SArray[i] {
			SArray[i], SArray[left] = SArray[left], SArray[i]
		}
		if SArray[right] > SArray[i] {
			SArray[i], SArray[right] = SArray[right], SArray[i]
		}

	}
	return SArray
}

func main() {
	var l [10]int
	array := make([]int, 10)
	for i, _ := range l {

		array[i] = rand.Intn(1000)
	}
	fmt.Println("before: ", array)
	s := SF{array}
	fmt.Println("after: ", s.Heap(array))
}
