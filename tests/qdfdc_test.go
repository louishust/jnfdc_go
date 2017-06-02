package test

import (
	"jnfdc/tasks"
	"testing"
)

// func TestQDFirst(t *testing.T) {
// 	err := tasks.FetchQDfdcFirst()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func TestQDSecond(t *testing.T) {
	err := tasks.FetchQDfdcSecond()
	if err != nil {
		t.Error(err)
	}
}
