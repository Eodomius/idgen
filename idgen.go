package idgen

import (
	"time"
)

const epoch = 1640991600000 // Jan 1st, 2022
const workerID = uint64(0)
const processID = uint64(1)
var increment = uint64(0)

type Snowflake struct {
	timestamp uint64
	workerID  uint64
	processID uint64
	increment uint64
}

func Generate() uint64 {
	timestamp := uint64(time.Now().UnixNano() / int64(time.Millisecond) - epoch)
	if increment >= 4096 {
		increment = 0
	}
	increment++
	return ((timestamp << 22) | (workerID << 17) | (processID << 12) | increment)
}

func Destructure(id uint64) Snowflake {
	timestamp := (id >> 22) + epoch
	workerID := (id >> 17)
	processID := (id >> 12)
	increment := id
	return Snowflake{timestamp, workerID, processID, increment}
}

func GetTimestamp(id uint64) uint64 {
	return (id >> 22) + epoch
}
func GetWorkerID(id uint64) uint64 {
	return (id >> 17)
}
func GetProcessID(id uint64) uint64 {
	return (id >> 12)
}
func GetIncrement(id uint64) uint64 {
	return id
}
