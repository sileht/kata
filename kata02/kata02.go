package kata02

func BinaryChopRecursive(searchValue int, array []int, initialIndex int) int {
	count := len(array)

	if count == 0 {
		return -1
	}

	middleIndex := count / 2

	if searchValue == array[middleIndex] {
		return initialIndex + middleIndex
	} else if searchValue > array[middleIndex] {
		return BinaryChopRecursive(searchValue, array[middleIndex+1:], initialIndex+middleIndex+1)
	} else {
		return BinaryChopRecursive(searchValue, array[:middleIndex], initialIndex)
	}
}

func ChopV1(searchValue int, array []int) int {
	// Recursive version
	// If the array is really big, we can reach goroutine MaxStack and the
	// program will stop
	return BinaryChopRecursive(searchValue, array, 0)
}

func ChopV2(searchValue int, array []int) int {
	// While loop version
	// less readable, but doesn't have the recurision limit
	remainingArray := array
	count := len(remainingArray)
	position := 0

	for count > 0 {
		middleIndex := count / 2
		if searchValue == remainingArray[middleIndex] {
			return position + middleIndex
		} else if searchValue > remainingArray[middleIndex] {
			remainingArray = remainingArray[middleIndex+1:]
			position = position + middleIndex + 1
		} else {
			remainingArray = remainingArray[:middleIndex]
		}
		count = len(remainingArray)
	}
	return -1
}

func ChopV3(searchValue int, array []int) int {
	// For loop version
	// Just a variant of the while loop with a ugly for syntax :)
	remainingArray := array
	count := len(remainingArray)
	position := 0

	for middleIndex := count / 2; count > 0; middleIndex = count / 2 {
		if searchValue == remainingArray[middleIndex] {
			return position + middleIndex
		} else if searchValue > remainingArray[middleIndex] {
			remainingArray = remainingArray[middleIndex+1:]
			position = position + middleIndex + 1
		} else {
			remainingArray = remainingArray[:middleIndex]
		}
		count = len(remainingArray)
	}

	return -1
}

func ChopV4(searchValue int, array []int) int {
	// goto version
	// well, goto...
	remainingArray := array
	position := 0

restart:
	count := len(remainingArray)
	if count == 0 {
		return -1
	}

	middleIndex := count / 2
	if searchValue > remainingArray[middleIndex] {
		remainingArray = remainingArray[middleIndex+1:]
		position = position + middleIndex + 1
		goto restart
	} else if searchValue < remainingArray[middleIndex] {
		remainingArray = remainingArray[:middleIndex]
		goto restart
	}
	return position + middleIndex
}

func BinaryChopGoroutine(result chan int, searchValue int, array []int, initialIndex int) {
	count := len(array)

	if count == 0 {
		result <- -1
		return
	}

	middleIndex := count / 2

	if searchValue == array[middleIndex] {
		result <- initialIndex + middleIndex
	} else if searchValue > array[middleIndex] {
		go BinaryChopGoroutine(result, searchValue, array[middleIndex+1:], initialIndex+middleIndex+1)
	} else {
		go BinaryChopGoroutine(result, searchValue, array[:middleIndex], initialIndex)
	}
}

func ChopV5(searchValue int, array []int) int {
	// Goroutine version, just for fun
	// Same as recursive version but with goroutine
	// We will never have two goroutine in //
	result := make(chan int)
	go BinaryChopGoroutine(result, searchValue, array, 0)
	return <-result
}
