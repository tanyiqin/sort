package main

import "fmt"

var buffer = []int{65,23,897,12,43,23,78,44,213,867,21312,5,3,2,6,8,3,4,56,1111,545646,87,234,5467,867,23,46,54,6752,32342,356,345,434,234,234,23341,23,126,54,6756,}

type sortType int
const (
	bubble sortType = iota
	selection
	insert
	merge
	quick
	heap
)

func isSort(b []int) bool{
	if len(b) == 0 {
		return true
	}
	k := b[0]
	for _, v := range b {
		if k > v {
			return false
		}
		k = v
	}
	return true
}

// 冒泡排序
func bubbleSort(b []int) {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b) - i - 1; j++ {
			if b[j] > b[j+1] {
				b[j], b[j+1] = b[j+1], b[j]
			}
		}
	}
}

// 选择排序
func selectSort(b []int) {
	for i := 0; i < len(b)-1; i++ {
		min := i
		for j := i+1; j < len(b); j++ {
			if b[j] < b[min] {
				min = j
			}
		}
		b[min], b[i] = b[i], b[min]
	}
}

// 插入排序
func insertSort(b []int) {
	for i := 1; i < len(b); i++ {
		j := i
		temp := b[j]
		for j > 0 && b[j-1] > temp {
			b[j] = b[j-1]
			j--
		}
		b[j] = temp
	}
}

// 归并排序
func mergeSort(b []int) {
	if len(b) < 2 {
		return
	}
	middle := len(b)/2
	mergeSort(b[0:middle])
	mergeSort(b[middle:])
	mergef(b, 0, len(b), middle)
}

func mergef(b []int, left, right, mid int) {
	temp := make([]int, right-left+1)
	index := 0
	i,j:=left,mid
	for i<mid && j<right {
		if b[i] < b[j] {
			temp[index] = b[i]
			i++
		} else {
			temp[index] = b[j]
			j++
		}
		index++
	}
	for i<mid {
		temp[index] = b[i]
		i++
		index++
	}
	for j<right{
		temp[index] = b[j]
		j++
		index++
	}
	for x:=0;x<index;x++{
		b[left]=temp[x]
		left++
	}
}

// 快速排序
func quickSort(b []int) {
	if len(b) < 2 {
		return
	}
	pivot := b[0]
	left, right := 0, len(b)-1
	for left < right {
		for left < right && b[right] >= pivot {
			right--
		}
		if b[right] < pivot {
			b[left] = b[right]
		}
		for left < right && b[left] <= pivot {
			left++
		}
		if b[left] > pivot {
			b[right] = b[left]
		}
	}
	b[left] = pivot
	quickSort(b[:left])
	quickSort(b[left+1:])
}

// 堆排序
func buildHeap(b []int) {
	for i := len(b)/2 - 1; i >= 0; i-- {
		heapAdjust(b, i, len(b))
	}
}

func heapAdjust(b []int, i, len int) {
	largest := i
	left := 2*i+1
	right := 2*i+2
	if left < len && b[left] > b[largest] {
		largest = left
	}
	if right < len && b[right] > b[largest] {
		largest = right
	}
	if i != largest {
		swap(b, i, largest)
		heapAdjust(b, largest, len)
	}
}

func swap(b []int, i,j int) {
	b[i], b[j] = b[j], b[i]
}

func heapSort(b []int) {
	buildHeap(b)
	for i := len(b)-1; i >= 0; i-- {
		swap(b, 0, i)
		heapAdjust(b, 0, i)
	}
}

func sort(Type sortType, buff []int) bool{
	b := make([]int, len(buff))
	copy(b, buff)
	switch Type {
	case bubble:
		bubbleSort(b)
	case selection:
		selectSort(b)
	case insert:
		insertSort(b)
	case merge:
		mergeSort(b)
	case quick:
		quickSort(b)
	case heap:
		heapSort(b)
	}
	return isSort(b)
}

func main() {
	for i := 0; i < 6; i++ {
		fmt.Println(sort(sortType(i), buffer))
	}
}