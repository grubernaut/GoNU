package main

import (
	"reflect"
	"testing"
)

func TestAppendEnd(t *testing.T) {
	testBytes := []byte{'a', 'b', 'c', 'd'}
	res := appendEnd(testBytes)
	expected := []byte{'a', 'b', 'c', 'd', '$'}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Bad: %#v", string(res))
	}
}

func TestAppendTab(t *testing.T) {
	testBytes := []byte{'\t', 'a', 'b'}
	res := appendTab(testBytes)
	expected := []byte{'^', 'I', 'a', 'b'}

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Bad: %#v", string(res))
	}
}

func TestNumberLine(t *testing.T) {
	testBytes := []byte{'a', 'b', 'c'}
	res := numberLine(testBytes, 1)
	expected := []byte("     1  abc")

	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("Bad: %#v", string(res))
	}
}
