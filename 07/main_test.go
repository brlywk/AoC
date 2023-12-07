package main

import (
	aochelper "brlywk/AoC/helper"
	"log"
	"os"
	"reflect"
	"testing"
)

const TestFile = "input_test.txt"

var testData *aochelper.InputData
var testLines []string

func testSetup() error {
	var err error
	testData, err = aochelper.NewInputData(TestFile, true)
	if err != nil {
		log.Fatalf("Error creating testData: %v", err)
	}

	testLines = testData.GetLines()

	return nil
}

func TestMain(m *testing.M) {
	if err := testSetup(); err != nil {
		log.Fatalf("Test setup encountered an error: %v", err)
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

// ---- Helper ----------------------------------

func TestGetCardValue(t *testing.T) {
	test := "A"
	expected := 14
	actual := GetCardValue(&test, &Cards)

	if actual != expected {
		t.Errorf("Expected %v\tActual %v", expected, actual)
	}
}

func TestStringToMap(t *testing.T) {
	test := "QQQJA"
	expected := map[string]int{"Q": 3, "J": 1, "A": 1}
	actual := StringToMap(&test)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected  %v\tActual %v", expected, actual)
	}
}

func TestGetHandType(t *testing.T) {
	hands := []string{"AAAAA", "AAAA2", "AAA22", "AAA42", "AA442", "AA423", "23456"}
	expected := []HandType{FiveOfAKind, FourOfAKind, FullHouse, ThreeOfAKind, TwoPair, OnePair, HighCard}
	var actual []HandType

	for _, hand := range hands {
		hm := StringToMap(&hand)
		ht := GetHandType(&hm)
		actual = append(actual, ht)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v\tActual %v", expected, actual)
	}
}

func TestHandsBetterOrNotWhoKnowsIdontThatsWhyIWriteThis(t *testing.T) {
	h1 := Hand{Raw: "KK677", Type: TwoPair}
	h2 := Hand{Raw: "KTJJT", Type: TwoPair}
	h3 := Hand{Raw: "T55J5", Type: ThreeOfAKind}
	h4 := Hand{Raw: "QQQJA", Type: ThreeOfAKind}
	h5 := Hand{Raw: "AAAAA", Type: FiveOfAKind}

	a1 := h1.IsBetterThan(&h2)
	e1 := true
	a2 := h1.IsBetterThan(&h3)
	e2 := false

	a3 := h3.IsBetterThan(&h4)
	e3 := false

	a4 := h5.IsBetterThan(&h4)
	e4 := true

	if !e1 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e1, a1, h1, h2)
	}

	if e2 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e2, a2, h1, h3)
	}

	if e3 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e3, a3, h3, h4)
	}

	if !e4 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e4, a4, h5, h4)
	}
}

func TestCompareHands(t *testing.T) {
	h1 := Hand{Raw: "KK677", Type: TwoPair}
	h2 := Hand{Raw: "KTJJT", Type: TwoPair}
	h3 := Hand{Raw: "T55J5", Type: ThreeOfAKind}
	h4 := Hand{Raw: "QQQJA", Type: ThreeOfAKind}
	h5 := Hand{Raw: "AAAAA", Type: FiveOfAKind}

	a1 := h1.Compare(&h2)
	e1 := 1
	a2 := h1.Compare(&h3)
	e2 := -1

	a3 := h3.Compare(&h4)
	e3 := -1

	a4 := h5.Compare(&h4)
	e4 := 1

	if e1 != 1 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e1, a1, h1, h2)
	}

	if e2 != -1 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e2, a2, h1, h3)
	}

	if e3 != -1 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e3, a3, h3, h4)
	}

	if e4 != 1 {
		t.Errorf("Expected: %v\tActual: %v\tH1: %v\tH2: %v", e4, a4, h5, h4)
	}
}

func TestSortHands(t *testing.T) {
	h1 := Hand{Raw: "KK677", Type: TwoPair}
	h2 := Hand{Raw: "KTJJT", Type: TwoPair}
	h3 := Hand{Raw: "T55J5", Type: ThreeOfAKind}
	h4 := Hand{Raw: "QQQJA", Type: ThreeOfAKind}
	h5 := Hand{Raw: "AAAAA", Type: FiveOfAKind}

	actual := []Hand{h3, h1, h5, h4, h2}
	SortHands(&actual)
	expected := []Hand{h5, h4, h3, h1, h2}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v\tActual: %v", expected, actual)
	}
}

// ---- Part 1 ----------------------------------

func TestEvaluatePart1(t *testing.T) {
	hands := ParseInput(&testLines)
	hands = PreparePar1(&hands)

	actual := EvaluatePart1(&hands)
	expected := 6440

	if expected != actual {
		t.Errorf("Expected: %v\tActual: %v", expected, actual)
	}
}
