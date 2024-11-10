package main

import (
	"brlywk/AoC/helper"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const inputFile = "input.txt"
// const inputFile = "input_test.txt"

// ----- Main ------------------------------------

func main() {
	fmt.Println("Advent of Code 2023 - Day 19")
	data, _ := aochelper.NewInputData(inputFile, false)
	lines := data.GetLines()
	workflows, parts := parseInput(&lines)

	accepted, _ := evaluateParts(&parts, &workflows)

	part1 := evaluatePart1(&accepted)
	fmt.Printf("Part 1:\t%v\n", part1)
}

// ----- Structs ---------------------------------

type rule struct {
	category string
	operator string
	value    int
	sendTo   string
}

func (r rule) String() string {
	return fmt.Sprintf("{ %v %v %v -> %v }", r.category, r.operator, r.value, r.sendTo)
}

// Special workflows (name):
// in	starting workflow
// A	accept part
// R	reject part
type workflow struct {
	name     string
	rules    []rule
	fallback string
}

func (w workflow) String() string {
	return fmt.Sprintf("{ %v\tRules: %v\tFallback: %v }\n", w.name, w.rules, w.fallback)
}

type part struct {
	x int // eXtremely cool looking
	m int // Musical
	a int // Aerodynamic
	s int // Shiny
}

func (p part) String() string {
	return fmt.Sprintf("{ x = %v\tm = %v\ta = %v\ts = %v }\n", p.x, p.m, p.a, p.s)
}

func (p *part) GetRating() int {
	return p.x + p.m + p.a + p.s
}

// Evaluates a rule on this part and specifies whether the rule applies or not
func (p *part) PassesRule(r *rule) bool {
	pVal := 42

	switch r.category {
	case "x":
		pVal = p.x
	case "m":
		pVal = p.m
	case "a":
		pVal = p.a
	case "s":
		pVal = p.s
	default:
		aochelper.PrintAndQuit("Super fatal error:\nRule %v tried to access non-existing category %v", r, r.category)
	}

	return op(r.operator, pVal, r.value)
}

// ----- Helper ----------------------------------

func op(operator string, a int, b int) bool {
	switch operator {
	case "<":
		return a < b
	case ">":
		return a > b
	default:
		aochelper.PrintAndQuit("Hyper fatal error:\nUnknown operator '%v'", operator)
	}
	return false
}

func parseInput(lines *[]string) ([]workflow, []part) {
	workflows := []workflow{}
	parts := []part{}

	for _, line := range *lines {
		if line == "" {
			continue
		}

		// parts start with '{'
		if line[0] == '{' {
			ratings := strings.Split(line, ",")
			x, _ := strconv.Atoi(string([]byte(ratings[0])[3:]))
			m, _ := strconv.Atoi(string([]byte(ratings[1])[2:]))
			a, _ := strconv.Atoi(string([]byte(ratings[2])[2:]))
			s, _ := strconv.Atoi(string([]byte(ratings[3])[2 : len(ratings[3])-1]))

			newPart := part{x, m, a, s}

			parts = append(parts, newPart)
		} else {
			curlyIdx := strings.Index(line, "{")
			name := string([]byte(line)[:curlyIdx])
			rules := string([]byte(line)[curlyIdx+1 : len(line)-1])
			rulesSplit := strings.Split(rules, ",")
			// fallback/default rule is always the last one
			fallback := rulesSplit[len(rulesSplit)-1]

			newWf := workflow{
				name:     name,
				rules:    []rule{},
				fallback: fallback,
			}

			re := regexp.MustCompile(`(\w{1})([<>]{1})(\d+):(\w+)`)
			for i := 0; i < len(rulesSplit)-1; i++ {
				matches := re.FindAllStringSubmatch(rulesSplit[i], -1)

				for _, m := range matches {
					// first element is always the whole string matched, capture groups
					// come after that
					subMatches := m[1:]

					v, _ := strconv.Atoi(subMatches[2])

					newRule := rule{
						category: subMatches[0],
						operator: subMatches[1],
						value:    v,
						sendTo:   subMatches[3],
					}

					newWf.rules = append(newWf.rules, newRule)
				}
			}

			workflows = append(workflows, newWf)
		}
	}

	return workflows, parts
}

func getWorkflowByName(name string, workflows *[]workflow) workflow {
	for _, w := range *workflows {
		if w.name == name {
			return w
		}
	}

	return workflow{}
}

func evaluateParts(parts *[]part, workflows *[]workflow) (accaptedParts []part, rejectedParts []part) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	start := getWorkflowByName("in", workflows)

	for i, pt := range *parts {
		wg.Add(1)

		go func(w int, p part, currentWorkflow workflow) {
			defer wg.Done()
			// defer fmt.Printf("\nWorker %v finished.\n", w)

			// fmt.Printf("\nWorker %v started:\n%v)\n", w, p)
			// fmt.Printf("%v:\tStarting with: %v\n", w, currentWorkflow)

			done := false

			for !done {
				// fmt.Printf("\n\t%v:\tWorkflow\n\t%v", w, currentWorkflow)
				useFallback := true

				// check rules and break rule eval as soon as a passing rule is found
				for _, r := range currentWorkflow.rules {
					// fmt.Printf("\t\t%v:\tRule: %v\n\t\t\tPart: %v\n", w, r, p)
					if p.PassesRule(&r) {
						if r.sendTo == "A" || r.sendTo == "R" {
							mutex.Lock()
							if r.sendTo == "A" {
								accaptedParts = append(accaptedParts, p)
							} else {
								rejectedParts = append(rejectedParts, p)
							}
							mutex.Unlock()

							// fmt.Printf("\t\t\t%v:\tReached termination: %v", w, r.sendTo)
							done = true
							useFallback = false
							break
						}

						// fmt.Printf("\t\t\t%v:\tRule accepted! Moving to %v\n", w, r.sendTo)
						currentWorkflow = getWorkflowByName(r.sendTo, workflows)
						useFallback = false
						break
					}
				}

				// no passing rule found, use fallback
				if useFallback {
					// we have reached the end and need to terminate this worker
					if currentWorkflow.fallback == "A" || currentWorkflow.fallback == "R" {
						mutex.Lock()
						if currentWorkflow.fallback == "A" {
							accaptedParts = append(accaptedParts, p)
						} else {
							rejectedParts = append(rejectedParts, p)
						}
						mutex.Unlock()

						// fmt.Printf("\t\t\t%v:\tReached termination fallback: %v", w, currentWorkflow.fallback)
						done = true
					}

					// fmt.Printf("\t\t\t%v:\tNo rule applied, using fallback. Sending to %v\n", w, currentWorkflow.fallback)
					currentWorkflow = getWorkflowByName(currentWorkflow.fallback, workflows)
					useFallback = false
				}
			}

		}(i, pt, start)
	}

	wg.Wait()

	return accaptedParts, rejectedParts
}

func evaluatePart1(acceptedParts *[]part) int {
	defer aochelper.Measure("Part 1")()
	sum := 0

	for _, p := range *acceptedParts {
		sum += p.GetRating()
	}

	return sum
}
