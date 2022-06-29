package core

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDatabaseSafeADDElt : test Database add function
func TestDatabaseSafeADDElt(t *testing.T) {
	db := NewDatabase()
	fmt.Println("Initial state : ", db.Len())

	db.Add(Message{
		Message: "test without tag",
	})

	tag := "tag"
	db.Add(Message{
		Message: "test with tag",
		Tag:     &tag,
	})

	db.Add(Message{
		Message: "test without tag",
	})

	fmt.Println("After adding 3 elts : ", db.Len())
}

// TestConcurrency
func TestConcurrency(t *testing.T) {
	assert := assert.New(t)
	db := NewDatabase()
	for i := 0; i < 10; i++ {
		go db.Add(Message{
			Message: "Message " + strconv.Itoa(i),
		})
	}
	fmt.Println("test Concurrency")

	time.Sleep(time.Millisecond * time.Duration(20))
	fmt.Println("Get Elts", db.Get(""))
	assert.Equal(db.Len(), 10, "We expects 10 elements")
}

func TestFilter(t *testing.T) {
	db := NewDatabase()
	db.Add(Message{
		Message: "test without tag",
	})

	tag := "tag"
	db.Add(Message{
		Message: "test with tag",
		Tag:     &tag,
	})

	db.Add(Message{
		Message: "test without tag",
	})

	fmt.Println("test no filter", db.Get(""))

	fmt.Println("test filter tag=tag", db.Get("tag"))

	fmt.Println("test filter tag=invalid", db.Get("invalid"))
}
