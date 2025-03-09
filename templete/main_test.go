package main

import (
	"io"
	"log"
	"os"
	"testing"
)

func TestMakeAnswer(t *testing.T) {
	got, err := MakeAnswer()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open("expected.txt")
	if err != nil {
		log.Fatal(err)
	}
	byteData, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	expected := string(byteData)

	if got != expected {
		t.Errorf("\nAnswer:\n%s\nExpected:\n%s", got, expected)
	}
}
