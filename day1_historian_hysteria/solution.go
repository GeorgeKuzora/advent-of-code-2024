package day1

import "math"


func FindDistance(l1 []int, l2 []int) int  {
	quickSort(&l1)
	quickSort(&l2)
	sum := 0
	
	for i := 0; i < len(l1); i++ {
		sum += int(math.Abs(float64(l1[i]) - float64(l2[i])))
	}
	return sum
}

func quickSort(list *[]int) {
	quickSortStep(list, 0, len(*list)-1)
}

func quickSortStep(list *[]int, lowId int, highId int) {
	if lowId < highId {
		pivotId := lomutoPartition(list, lowId, highId)
		quickSortStep(list, lowId, pivotId-1)
		quickSortStep(list, pivotId+1, highId)
	}
}

func lomutoPartition(listP *[]int, lowId int, highId int) int {
	list := *listP
	smallerElementsBoundaryId := lowId - 1
	currentId := lowId
	for currentId < highId {
		if list[currentId] <= list[highId] {
			smallerElementsBoundaryId++
			swapElements(listP, smallerElementsBoundaryId, currentId)
		}
		currentId++
	}
	newPivotId := smallerElementsBoundaryId + 1
	swapElements(listP, newPivotId, highId)
	return newPivotId 
}

func swapElements(listP *[]int, firstId int, secondId int) {
	list := *listP
	tempEl := list[firstId]
	list[firstId] = list[secondId]
	list[secondId] = tempEl
}

func hoarePartition(listP *[]int, lowId int, highId int) int {
	list := *listP
	lowPointerId := lowId
	highPointerId := highId
	pivotId := lowId
	pivotValue := list[pivotId]
	for lowPointerId < highPointerId {
		for list[lowPointerId] < pivotValue && lowPointerId <= highPointerId {
			lowPointerId++
		}
		for list[highPointerId] >= pivotValue && highPointerId >= lowPointerId {
			highPointerId++
		}
		swapElements(listP, lowPointerId, highPointerId)
	}
	return lowPointerId
}
