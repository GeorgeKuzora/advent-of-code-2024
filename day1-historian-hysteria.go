package main

func findDistance(l1 []int, l2 []int)  {
	
}

func quickSort(list *[]int) {

}

func quickSortStep(list *[]int, lowId int, highId int) {
	if len(*list) == 0 {
		return
	}
	pivotId := lomutoPartition(list, lowId, highId)

	quickSortStep(list, lowId, pivotId-1)
	quickSortStep(list, pivotId+1, highId)
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
