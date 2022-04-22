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
	snowflake := Deconstruct(id)
	if snowflake.timestamp != GetTimestamp(id) {
		t.Errorf("GetTimestamp() returned %v, expected %v", GetTimestamp(id), snowflake.timestamp)
	}
	if snowflake.workerID != GetWorkerID(id) {
		t.Errorf("GetWorkerID() returned %v, expected %v", GetWorkerID(id), snowflake.workerID)
	}
	if snowflake.processID != GetProcessID(id) {
		t.Errorf("GetProcessID() returned %v, expected %v", GetProcessID(id), snowflake.processID)
	}
	if snowflake.increment != GetIncrement(id) {
		t.Errorf("GetIncrement() returned %v, expected %v", GetIncrement(id), snowflake.increment)
	}
}
