package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// ----- Globals Vars ----------------------------

const (
	TestFile  = "input_test.txt"
	InputFile = "input.txt"
)

// ----- Structs & Methods -----------------------

type Seeds []int

type MapRange struct {
	Start  int
	Length int
}

type Mapping map[MapRange]MapRange

type MappingBlock struct {
	Name string
	Mapping
}

// ----- Main -----------------------------------

func main() {
	log.Println("Advent of Code - Day 5")

	content := ReadFile(InputFile)
	lines := ParseInput(&content)

	seeds := GetSeedData(&lines)
	blocks := GetMappingBlocks(&lines)

	part1 := EvaluatePart1(&seeds, &blocks)

	log.Printf("Part 1: %v", part1)
}

// ----- Helper ---------------------------------

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Unable to read %v. Terminating.", fileName)
	}

	return string(content)
}

func ParseInput(content *string) []string {
	var result []string
	lines := strings.Split(strings.ReplaceAll(*content, "\r\n", "\n"), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			result = append(result, line)
		}
	}

	return result
}

// Creates a slice filled with numbers from start to end
//
// Note:
// The range is half-open, i.e. start is included, but end is not included
func CreateRange(start int, end int) []int {
	result := make([]int, end-start)

	for i := start; i < end; i++ {
		result[i-start] = i
	}

	return result
}

// ----- Part 1 ----------------------------------

// Extract the seed information from the first line of the input
func GetSeedData(lines *[]string) Seeds {
	var seeds []int
	l := *lines

	seedsLine := l[0]
	seedsLineStr, _ := strings.CutPrefix(seedsLine, "seeds: ")
	seedsStr := strings.Split(seedsLineStr, " ")

	for _, s := range seedsStr {
		snum, _ := strconv.Atoi(s)
		seeds = append(seeds, snum)
	}

	return seeds
}

// Create the range for source and destination from the current line
func CreateLineRangeMapping(line *string) Mapping {
	results := map[MapRange]MapRange{}
	l := *line

	lSlice := strings.Split(l, " ")
	var nums []int
	for _, lstr := range lSlice {
		num, _ := strconv.Atoi(lstr)
		nums = append(nums, num)
	}

	destStart := nums[0]
	sourceStart := nums[1]
	length := nums[2]

	startRange := MapRange{
		Start:  sourceStart,
		Length: length,
	}
	destRange := MapRange{
		Start:  destStart,
		Length: length,
	}

	results[startRange] = destRange

	return results
}

// Go through all lines and create blocks of mappings;
// A block has the name of the block and the mapping of sources to destinations specified
func GetMappingBlocks(lines *[]string) [7]MappingBlock {
	var results [7]MappingBlock
	l := *lines

	mapLines := l[1:]

	currentBlock := MappingBlock{Name: "", Mapping: map[MapRange]MapRange{}}
	blockCounter := 0

	for _, line := range mapLines {
		// it's the name of the block, save the block if necessary
		if strings.Index(line, ":") > -1 {
			newBlockFoundName := line[:len(line)-5]
			// log.Printf("New Block found:\t%v\n", newBlockFoundName)

			if currentBlock.Name == "" {
				currentBlock.Name = newBlockFoundName
			}

			if newBlockFoundName != currentBlock.Name {
				// log.Printf("Starting new block. Saving: %v", currentBlock)
				results[blockCounter] = currentBlock
				blockCounter++
				currentBlock.Name = newBlockFoundName
				currentBlock.Mapping = map[MapRange]MapRange{}
			}
		} else {
			// it's a mapping of numbers to numbers we need to handle
			// log.Println("Creating Line Range Mapping")
			tmpMap := CreateLineRangeMapping(&line)

			// log.Println("Iterating over Line Range Mapping")
			for k, v := range tmpMap {
				// log.Printf("TmpMap\tK: %v\tV: %v\n", k, v)
				currentBlock.Mapping[k] = v
			}
		}
	}

	// add the last block
	results[len(results)-1] = currentBlock

	return results
}

// Get the destination for a mapping given a source;
// If the source is not mapped, destination is the same as the source
func GetDestination(source int, mapping *Mapping) int {
	var result int

	isMapped := false
	var foundDestMap MapRange
	var sourcePos int

	for src, dest := range *mapping {
		if source >= src.Start && source < src.Start+src.Length {
			isMapped = true
			foundDestMap = dest
			sourcePos = int(math.Abs(float64(src.Start - source)))
		}
	}

	if isMapped {
		result = sourcePos + foundDestMap.Start
	} else {
		result = source
	}

	return result
}

// Get the location for a single seed
func GetLocationForSeed(seed int, mappingBlocks *[7]MappingBlock) int {
	result := seed

	for _, block := range *mappingBlocks {
		result = GetDestination(result, &block.Mapping)
	}

	return result
}

// Evaluate the mapping data and return the lowest location number
func EvaluatePart1(seeds *Seeds, mappingBlocks *[7]MappingBlock) int {
	lowestLocation := math.MaxInt

	for _, seed := range *seeds {
		location := GetLocationForSeed(seed, mappingBlocks)
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}
