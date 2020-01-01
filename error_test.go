package main

import (
	error "github.com/gabrielpsilva/sonicspeed/error"
	"testing"
)


func TestError_New(t *testing.T) {

	e := error.New(1400,"test", "")
	if e.Documentation != error.NoSolution{
		t.Error("error.NoSolution expected as documentation")
	}

	e2 := error.NewE(1400, &e, "error :(")
	if e2.Documentation != "error :(" {
		t.Error("\"error :(\" expected as documentation")
	}

}

