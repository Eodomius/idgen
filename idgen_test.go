package idgen

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	id := Generate()
	if id == 0 {
		t.Error("Generate() returned 0")
	}
}
func TestDestructure(t *testing.T){
	id := Generate()
	snowflake := Destructure(id)
	if snowflake.timestamp != GetTimestamp(id) {
		t.Errorf("GetTimestamp() returned %v, expected %v", GetTimestamp(id), snowflake.timestamp)
	}
	if snowflake.workerID != getWorkerID(id) {
		t.Errorf("getWorkerID() returned %v, expected %v", getWorkerID(id), snowflake.workerID)
	}
	if snowflake.processID != getProcessID(id) {
		t.Errorf("getProcessID() returned %v, expected %v", getProcessID(id), snowflake.processID)
	}
	if snowflake.increment != getIncrement(id) {
		t.Errorf("getIncrement() returned %v, expected %v", getIncrement(id), snowflake.increment)
	}
}

func TestPerformances(t *testing.T){
	// Musure time for Generate()
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		Generate()
	}
	elapsed := time.Since(start)
	if elapsed.Nanoseconds() / 1000000 > 7 {
		t.Errorf("Generate() took %v", elapsed)
	}
	fmt.Println("\n- Generate() took", elapsed, "\n- Nanoseconds per id:", elapsed.Nanoseconds() / 1000000, "\n\u200b")
}