package main

func Add10(pointer *int) {
	*pointer = *pointer + 10
}

func mutiplyBy2(pointer *[]int) {
	for i := range *pointer {
		(*pointer)[i] = (*pointer)[i] * 2
	}
}
