package main

import (
	"os"
	"testing"
)

var (
	fileContent string
)

func TestMain(m *testing.M) {
	// an error reading the file terminates the program, that's
	// good enough for our case...
	fileContent = ReadFile(TestFile)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateRange(t *testing.T) {
	r := CreateRange(5, 10)
	e := []int{5, 6, 7, 8, 9}

	lr := len(r)
	le := len(e)

	if lr != le {
		t.Errorf("Length mismatch. Expected %v, got %v", le, lr)
	}

	for i, n := range r {
		if n != e[i] {
			t.Errorf("Element mismatch. Expected %v, got %v", e[i], n)
		}
	}
}

func TestParseContent(t *testing.T) {
	l := ParseInput(&fileContent)
	r := len(l)
	e := 26 // empty lines will be removed by ParseInput

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestGetSeedData(t *testing.T) {
	l := ParseInput(&fileContent)
	r := GetSeedData(&l)
	e := []int{79, 14, 55, 13}

	if len(r) != len(e) {
		t.Errorf("Length mismatch. Expected %v, got %v", len(e), len(r))
	}

	for i, n := range r {
		if n != e[i] {
			t.Errorf("Data mismatch. Expected %v, got %v", e[i], n)
		}
	}
}

// No the best test, but whatever...
func TestGetMappingBlocks(t *testing.T) {
	l := ParseInput(&fileContent)
	r := GetMappingBlocks(&l)
	lr := len(r)
	le := 7

	if lr != le {
		t.Errorf("Length mismatch. Expected %v, got %v", le, lr)
	}

	for _, v := range r {
		if len(v.Mapping) == 0 || v.Name == "" {
			t.Errorf("Unexpected empty mapping object %v", v)
		}
		// log.Println(v)
		// log.Println("--------------------------------")
	}
}

// Multiple tests in one is definitely the best practice in testing 😎
func TestGetDestination(t *testing.T) {
	m := Mapping{MapRange{Start: 42, Length: 2}: MapRange{Start: 123, Length: 2}}
	s1 := 42
	e1 := 123
	r1 := GetDestination(s1, &m)

	if r1 != e1 {
		t.Errorf("Expected %v, got %v", e1, r1)
	}

	s2 := 10
	e2 := 10
	r2 := GetDestination(s2, &m)

	if r2 != e2 {
		t.Errorf("Expected %v, got %v", e2, r2)
	}
}

func TestGetLocationForSeed(t *testing.T) {
	l := ParseInput(&fileContent)
	mp := GetMappingBlocks(&l)

	seedsAndExpectedResults := map[int]int{79: 82, 14: 43, 55: 86, 13: 35}

	for seed, expectedLoc := range seedsAndExpectedResults {
		loc := GetLocationForSeed(seed, &mp)

		if loc != expectedLoc {
			t.Errorf("Seed: %v, Expected: %v, got %v", seed, expectedLoc, loc)
		}
	}
}

func TestEvaluatePart1(t *testing.T) {
	lines := ParseInput(&fileContent)
	seeds := GetSeedData(&lines)
	maps := GetMappingBlocks(&lines)

	r := EvaluatePart1(&seeds, &maps)
	e := 35

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

// ----- Part 2 ----------------------------------

func TestConvertSeedsToSeedsPart2(t *testing.T) {
	lines := ParseInput(&fileContent)
	seeds := GetSeedData(&lines)

	r := ConvertSeedsToSeedsPart2(&seeds)
	e := SeedsPart2{MapRange{Start: 79, Length: 14}, MapRange{Start: 55, Length: 13}}

	lr := len(r)
	le := len(e)

	if lr != le {
		t.Errorf("Length mismatch. Expected %v, got %v", le, lr)
	}

	for i, mr := range r {
		if mr != e[i] {
			t.Errorf("Element mismatch. Expected %v, got %v", e[i], mr)

		}
	}
}


func TestEvaluatePart2(t *testing.T) {
	lines := ParseInput(&fileContent)
	initSeeds := GetSeedData(&lines)
	seeds:= ConvertSeedsToSeedsPart2(&initSeeds)
	maps := GetMappingBlocks(&lines)

	r := EvaluatePart2(&seeds, &maps)
	e := 46

	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
