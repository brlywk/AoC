package main

import "core:fmt"
import "core:log"
import "core:strconv"
import "core:strings"
import "core:testing"

TEST :: #load("./test.txt")
DATA :: #load("./data.txt")

main :: proc() {
	p1 := part1(DATA)
	fmt.println("Part 1:", p1)

	// p2 := part2(DATA)
	// fmt.println("Part 2:", p2)
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

Equation :: struct {
	expected: int,
	parts:    [dynamic]int,
}

// As always, we abuse the temp_allocator!
// And we assume super reliable input. Who needs error handling anyway?!
equation_new :: proc(input: string) -> Equation {
	parts := make([dynamic]int, context.temp_allocator)

	// split expected result from rest
	split := strings.split(input, ":", context.temp_allocator)

	expected, _ := strconv.parse_int(split[0])

	// trim rest to get rid of left whitespace
	s := strings.trim_left(split[1], " ")

	// split rest on whitespaces to get parts
	for num_str in strings.split_iterator(&s, " ") {
		num, _ := strconv.parse_int(num_str)
		append(&parts, num)
	}

	return {expected = expected, parts = parts}
}

equation_solvable :: proc(e: Equation, concat_enabled := false) -> bool {
	if len(e.parts) == 0 do return false
	if len(e.parts) == 1 do return e.expected == e.parts[0]

	// n numbers have n-1 operations...
	n := len(e.parts) - 1
	// ... and a total of 2^(n-1) permutations
	permutations := 1 << uint(n)

	for perm in 0 ..< permutations {
		result := e.parts[0]

		for i in 0 ..< n {
			// each permutation is a number, from 0..2^(n-1), so:
			// 0 -> 0000
			// 1 -> 0001
			// 2 -> 0010
			// 3 -> 0011
			// ..
			// 15 -> 1111
			// etc.
			//
			// we can use the lsb as the operation:
			// 0 = +
			// 1 = *
			if (perm >> uint(i)) & 1 == 0 {
				result += e.parts[i + 1]
			} else {
				result *= e.parts[i + 1]
			}
		}

		if result == e.expected {
			return true
		}
	}


	return false
}

input_parse :: proc(input: []u8) -> [dynamic]Equation {
	equations := make([dynamic]Equation, context.temp_allocator)

	it := string(input)
	for line in strings.split_lines_iterator(&it) {
		e := equation_new(line)
		append(&equations, e)
	}

	return equations
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	result := 0
	equations := input_parse(input)

	for e in equations do if equation_solvable(e) {
		result += e.expected
	}

	return fmt.aprint(result)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "3749"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	result := 42

	return fmt.aprint(result)
}

// @(test)
// part2_test :: proc(t: ^testing.T) {
// 	expected := "42"
// 	actual := part2(TEST)
// 	defer delete(actual)
//
// 	testing.expect_value(t, actual, expected)
// }

