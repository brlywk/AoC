use anyhow::anyhow;
use std::ops::RangeInclusive;

#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() -> anyhow::Result<()> {
    let p1 = part1(DATA)?;
    println!("Part 1: {p1}");

    let p2 = part2(DATA)?;
    println!("Part 2: {p2}");

    Ok(())
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

#[derive(Debug)]
struct Inventory {
    fresh_ranges: Vec<RangeInclusive<usize>>,
    food_ids: Vec<usize>,
}

impl Inventory {
    fn new(input: &str) -> anyhow::Result<Self> {
        let (ranges, ids) = input
            .split_once("\n\n")
            .ok_or(anyhow!("Blocks not found"))?;

        let fresh_ranges = ranges
            .lines()
            .filter_map(|line| {
                let (start, end) = line.split_once('-')?;

                let s = start.parse::<usize>().ok()?;
                let e = end.parse::<usize>().ok()?;

                Some(s..=e)
            })
            .collect::<Vec<RangeInclusive<usize>>>();

        let food_ids = ids
            .lines()
            .filter_map(|line| line.parse::<usize>().ok())
            .collect::<Vec<usize>>();

        Ok(Self {
            fresh_ranges,
            food_ids,
        })
    }

    fn fresh_ingredients(&self) -> usize {
        self.food_ids
            .iter()
            .filter(|id| self.fresh_ranges.iter().any(|range| range.contains(*id)))
            .count()
    }

    fn count_fresh_ids(&self) -> usize {
        // This doesn't work b/c it expands all ranges; and with the large numbers AoC uses
        // this takes A LOT of memory :(
        // self.fresh_ranges
        //     .iter()
        //     .flat_map(Clone::clone)
        //     .collect::<HashSet<usize>>()
        // .len()

        // So: merge overlapping ranges and count elements in each

        // take only the first and last number as a tuple and sort by start number
        // so when we iterate over these ranges, we can check if subsequent ranges
        // fall (i.e. start) within the "current range"
        let mut sorted_ranges: Vec<_> = self
            .fresh_ranges
            .iter()
            .map(|range| (*range.start(), *range.end()))
            .collect();
        sorted_ranges.sort_by_key(|r| r.0);

        let mut merged_ranges = vec![];
        let (mut current_start, mut current_end) = sorted_ranges[0];

        // iterate over all ranges and check if these ranges can be merged together:
        // take the lowest start and highest end of overlapping/adjacent ranges
        for &(start, end) in &sorted_ranges[1..] {
            if start <= current_end + 1 {
                current_end = current_end.max(end);
            } else {
                merged_ranges.push(current_start..=current_end);
                current_start = start;
                current_end = end;
            }
        }
        merged_ranges.push(current_start..=current_end);

        // now we can just count the numbers without counting the numbers :P
        merged_ranges
            .iter()
            .map(|range| range.end() - range.start() + 1)
            .sum()
    }
}

fn part1(input: &str) -> anyhow::Result<String> {
    Ok(Inventory::new(input)?.fresh_ingredients().to_string())
}

#[test]
fn part1_test() -> anyhow::Result<()> {
    assert_eq!(part1(TEST)?, "3");
    Ok(())
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

fn part2(input: &str) -> anyhow::Result<String> {
    Ok(Inventory::new(input)?.count_fresh_ids().to_string())
}

#[test]
fn part2_test() -> anyhow::Result<()> {
    assert_eq!(part2(TEST)?, "14");
    Ok(())
}
