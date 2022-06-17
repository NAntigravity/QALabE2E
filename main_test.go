package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementTotalCount(t *testing.T) {
	t.Run("Increment", func(t *testing.T) {
		totalCountBeforeTest, err := GetTotalCount()
		if err != nil {
			t.Error(err)
		}
		t.Logf("Total count before test: %d", totalCountBeforeTest)
		err = IncrementTotalCount()
		if err != nil {
			t.Error(err)
		}
		totalCountAfterTest, err := GetTotalCount()
		if err != nil {
			t.Error(err)
		}
		t.Logf("Total count after test: %d", totalCountAfterTest)
		if !assert.LessOrEqual(t, totalCountBeforeTest, totalCountAfterTest, "You are a good boy! Your code are beautiful!") {
			t.Errorf("Test count not incremented after sending request. You are a bad programmer. Try to find another job")
		}
	})
}
