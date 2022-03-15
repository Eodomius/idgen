package idgen

import "time"

const epoch = 1420070400000 // Jan 1st, 2022

func generate() uint64 {
	timestamp := uint64(time.Now().Unix()) - epoch
	workerID := uint64(0)
	processID := uint64(1)
	increment := uint64(1)
	return ((timestamp << 22) | (workerID << 17) | (processID << 12) | increment)
}