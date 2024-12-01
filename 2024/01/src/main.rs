use std::usize;

use common::prelude::*;

fn main() {
    println!("AoC 2024 - Day 1\n");

    let input = Input::new("../inputs/input.txt");

    println!("Part1: {}", part1(&input));
}

fn get_lists(input: &Input) -> (Vec<&str>, Vec<&str>) {
    input
        .lines
        .iter()
        .map(|line| line.split_once("   ").unwrap())
        .unzip()
}

fn part1(input: &Input) -> usize {
    let (mut left_list, mut right_list) = get_lists(&input);

    left_list.sort();
    right_list.sort();

    left_list
        .iter()
        .zip(right_list.iter())
        .map(|(&left, &right)| {
            let l_int = left.parse::<usize>().unwrap();
            let r_int = right.parse::<usize>().unwrap();

            l_int.abs_diff(r_int)
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = Input::new("inputs/test.txt");
        let p1 = part1(&input);
        assert_eq!(p1, 11);
    }
}
