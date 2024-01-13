package main

import (
	"bytes"
	"os"
	"testing"
	"tetris-optimizer/pkg"
)

func FindBoardMiniSize(t *testing.T) {
	content, err := os.ReadFile("../input/good-example00.txt")
	if err != nil {
		t.Errorf("error reading file")
	}
	div := []byte{10}
	newContent := bytes.Split(content, div)
	tet1 := pkg.FindTetrominoes(newContent)
	tet2 := pkg.FindBoardMinSize(tet1)

	if tet2 == 0 {
		t.Errorf("issue with minimum board size %v, and its value %v", newContent, tet2)
	}
}

func TestFindTetrominoes(t *testing.T) {
	content, err := os.ReadFile("../input/good-example00.txt")
	if err != nil {
		t.Errorf("error reading file")
	}
	div := []byte{10}
	newContent := bytes.Split(content, div)
	tet := pkg.FindTetrominoes(newContent)

	if tet == nil {
		t.Errorf("issue with %v, and %v is nil", newContent, tet)
	}
}
