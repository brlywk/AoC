package main

import "core:fmt"
import "core:slice"
import "core:strconv"
import "core:strings"
import "core:testing"

TEST :: #load("./test.txt")
DATA :: #load("./data.txt")

main :: proc() {
	p1 := part1(DATA)
	fmt.println("Part 1:", p1)

	p2 := part2(DATA)
	fmt.println("Part 2:", p2)
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

Rule :: [2]int
Update :: [dynamic]int

rule_parse :: proc(input: string) -> Rule {
	// as always:
	// assume AoC input is well formatted and massively abuse the temp_allocator :P
	tokens := strings.split(input, "|", context.temp_allocator)

	a, _ := strconv.parse_int(tokens[0])
	b, _ := strconv.parse_int(tokens[1])

	return {a, b}
}

input_parse :: proc(input: []u8) -> ([dynamic]Rule, [dynamic]Update) {
	rules := make([dynamic]Rule, context.temp_allocator)
	updates := make([dynamic]Update, context.temp_allocator)

	// prepare input data
	input_split := strings.split(string(input), "\n\n", context.temp_allocator)
	section_rules := input_split[0]
	section_updates := input_split[1]

	// parse rules
	for line in strings.split_lines_iterator(&section_rules) {
		rule := rule_parse(line)
		append(&rules, rule)
	}

	// parse updates
	for line in strings.split_lines_iterator(&section_updates) {
		update_line := make(Update, context.temp_allocator)

		line_copy := line
		for num_str in strings.split_iterator(&line_copy, ",") {
			num, _ := strconv.parse_int(num_str)
			append(&update_line, num)
		}

		append(&updates, update_line)
	}

	return rules, updates
}

validate_rules :: proc(input: []u8, sort := false) -> int {
	rules, updates := input_parse(input)
	result := 0

	// HACK: Assign the rules array to the context user_ptr so we can access it
	// in our custom sorting proc for updates
	context.user_ptr = &rules

	outer: for update in updates {
		for rule in rules {
			idx_x := slice.linear_search(update[:], rule.x) or_continue
			idx_y := slice.linear_search(update[:], rule.y) or_continue

			if idx_x > idx_y {
				if sort {
					u_sorted := update[:]

					slice.sort_by(u_sorted, proc(a, b: int) -> bool {
						rules := (^[dynamic]Rule)(context.user_ptr)

						for rule in rules {
							if rule.x == a && rule.y == b do return true
							if rule.x == b && rule.y == a do return false
						}

						return false
					})

					result += u_sorted[len(update) / 2]
				} else {
					continue outer
				}
			}
		}

		// if we arrive here, all applicable rules for this update are
		// valid and we can take the middle number
		if !sort {
			result += update[len(update) / 2]
		}
	}

	return result
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)
	result := validate_rules(input)

	return fmt.aprint(result)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "143"
	actual := part1(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

part2 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)
	result := validate_rules(input, true)

	return fmt.aprint(result)
}

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "123"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

