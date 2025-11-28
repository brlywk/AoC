package main

import "core:fmt"
import "core:log"
import "core:strconv"
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

Disk_Map :: struct {
	sparse_data:  [dynamic]int,
	compact_data: [dynamic]int,
}

input_parse :: proc(input: []u8) -> Disk_Map {
	sparse_data := make([dynamic]int, context.temp_allocator)
	compact_data := make([dynamic]int, context.temp_allocator)

	block_id := 0
	for num, i in input {
		n := num - '0'

		// to prevent invisible characters, new lines etc from messing everything up
		// and preventing us from having nice things!
		if n < 0 || n > 9 do continue

		if i % 2 == 0 {
			// even indices: blocks
			//
			// append the current block id n-times into sparse_data
			for b in 0 ..< n do append(&sparse_data, block_id)
			block_id += 1
		} else {
			// odd indices: free spaces
			//
			// NOTE: we could resize sparse_data by the number of free indices and fill it;
			// but let's keep it simple for now, as performance should be okay-ish anyway,
			// and that way we have an easier time knowing the right index to save
			for f in 0 ..< n {
				append(&sparse_data, -1) // -1 to indicate free space
			}
		}
	}

	return {sparse_data = sparse_data, compact_data = compact_data}
}

compact :: proc(disk_map: ^Disk_Map) {
	left := 0
	right := len(disk_map.sparse_data) - 1

	for left <= right {
		// append the number if it's a block
		if disk_map.sparse_data[left] != -1 {
			append(&disk_map.compact_data, disk_map.sparse_data[left])
			left += 1
		} else {
			// if right points at -1, move it one to the left
			for right >= left && disk_map.sparse_data[right] == -1 {
				right -= 1
			}

			// if we haven't met yet, copy the last num into a free field
			if right >= left {
				append(&disk_map.compact_data, disk_map.sparse_data[right])
				right -= 1
				left += 1
			}
		}
	}
}

checksum :: proc(disk_map: Disk_Map) -> int {
	sum := 0

	for n, i in disk_map.compact_data {
		sum += n * i
	}

	return sum
}

part1 :: proc(input: []u8) -> string {
	defer free_all(context.temp_allocator)

	disk_map := input_parse(input)
	compact(&disk_map)
	result := checksum(disk_map)

	return fmt.aprint(result)
}

@(test)
part1_test :: proc(t: ^testing.T) {
	expected := "1928"
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

@(test)
part2_test :: proc(t: ^testing.T) {
	expected := "2858"
	actual := part2(TEST)
	defer delete(actual)

	testing.expect_value(t, actual, expected)
}

