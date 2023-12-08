package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime/debug"

	"brlywk/AoC/helper"
)

// ---- Declarations ----------------------------

const inputFile = "input.txt"

// ---- Structs & Methods -----------------------

// Name = (Left, Right)
type Node struct {
	Name  string
	Left  *Node
	Right *Node
}

func (n Node) String() string {
	var ln, rn string

	if n.Left == nil {
		ln = "<NIL>"
	} else {
		ln = n.Left.Name
	}

	if n.Right == nil {
		rn = "<NIL>"
	} else {
		rn = n.Right.Name
	}

	return fmt.Sprintf("\n{Name: %v, Left: %v, Right: %v}\n", n.Name, ln, rn)
}

// ---- Main ------------------------------------

func main() {
	defer aochelper.Measure("Time")()
	fmt.Println("Advent of Code 2023 - Day 8")
	data, err := aochelper.NewInputData(inputFile, true)
	if err != nil {
		fmt.Printf("Unable to read input file: %v", err)
		os.Exit(1)
	}
	lines := data.GetLines()
	directions, nodes := ParseInput(&lines)

	part1 := FindStepCount(&directions, &nodes)
	fmt.Printf("Part 1:\t%v\n", part1)
}

// ---- Helper ----------------------------------

func printAndQuit(msg string, m ...any) {
	fmt.Printf(msg, m...)
	os.Exit(1)
}

func ParseInput(lines *[]string) (string, map[string]*Node) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
			debug.PrintStack()
		}
	}()

	nodes := make(map[string]*Node)

	// Pattern is the first line
	pattern := (*lines)[0]

	re := regexp.MustCompile(`([A-Z]{3,3})`)
	matchedNames := make(map[string][]string)

	// It's probably easiert to do this in two steps to make self-referential nodes
	// easier to handle
	for i := 1; i < len(*lines); i++ {
		matches := re.FindAllString((*lines)[i], -1)

		nodeName := matches[0]
		leftName := matches[1]
		rightName := matches[2]

		nodes[nodeName] = &Node{Name: nodeName}
		matchedNames[nodeName] = []string{leftName, rightName}
	}

	// Add the actual child nodes to each node
	for _, node := range nodes {
        node.Left = nodes[matchedNames[node.Name][0]]
        node.Right = nodes[matchedNames[node.Name][1]]
	}

	// fmt.Printf("Nodes: %v\n\n", nodes)

	return pattern, nodes
}

// ---- Part 1 ----------------------------------

func FindStepCount(pattern *string, nodes *map[string]*Node) int {
	steps := 0
	currentNode := (*nodes)["AAA"]

	for currentNode.Name != "ZZZ" {
		for _, dir := range *pattern {
			switch dir {
			case 'L':
				currentNode = currentNode.Left
			case 'R':
				currentNode = currentNode.Right
			}
			steps++
		}
	}

	return steps
}
