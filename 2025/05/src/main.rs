use std::ops::RangeInclusive;

use anyhow::anyhow;

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

    fn count_fresh(&self) -> usize {
        self.food_ids
            .iter()
            .filter(|id| self.fresh_ranges.iter().any(|range| range.contains(*id)))
            .count()
    }
}

fn part1(input: &str) -> anyhow::Result<String> {
    Ok(Inventory::new(input)?.count_fresh().to_string())
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
    let mut result: isize = 0;

    Ok(result.to_string())
}

#[test]
fn part2_test() -> anyhow::Result<()> {
    assert_eq!(part2(TEST)?, "42");
    Ok(())
}
