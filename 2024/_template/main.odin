package main

import "core:fmt"
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

part1 :: proc(input: []u8) -> string {
	return "42"
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "42"
	actual := part1(TEST)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

part2 :: proc(input: []u8) -> string {
	return "42"
}

// @(test)
// part2_test :: proc(t: ^testing.T) {
// 	expected := "42"
// 	actual := part2(TEST)
//
// 	testing.expect_value(t, actual, expected)
// }

