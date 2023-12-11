package main

import "sync"

type SafeArray struct {
	array []int
	locks []sync.Mutex
}

func NewSafeArray(size int) *SafeArray {
	safeArray := &SafeArray{
		array: make([]int, size),
		locks: make([]sync.Mutex, size),
	}
	return safeArray
}

func (sa *SafeArray) Write(index, value int) {

	sa.locks[index].Lock()
	defer sa.locks[index].Unlock()
	sa.array[index] = value
}

func (sa *SafeArray) Read(index int) int {
	sa.locks[index].Lock()
	defer sa.locks[index].Unlock()
	return sa.array[index]
}
