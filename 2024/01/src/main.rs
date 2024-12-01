use std::usize;

use common::prelude::*;

fn main() {
    println!("AoC 2024 - Day 1\n");

    let input = Input::new("../inputs/input.txt");

    println!("Part1: {}", part1(&input));
    println!("Part2: {}", part2(&input));
}

fn get_lists(input: &Input) -> (Vec<usize>, Vec<usize>) {
    input
        .lines
        .iter()
        .map(|line| {
            line.split_once("   ")
                .map(|(l, r)| (l.parse::<usize>().unwrap(), r.parse::<usize>().unwrap()))
                .unwrap()
        })
        .unzip()
}

fn part1(input: &Input) -> usize {
    let (mut left_list, mut right_list) = get_lists(&input);

    left_list.sort();
    right_list.sort();

    left_list
        .iter()
        .zip(right_list.iter())
        .map(|(&left, &right)| left.abs_diff(right))
        .sum()
}

fn part2(input: &Input) -> usize {
    let mut score: usize = 0;
    let (left_list, right_list) = get_lists(&input);

    left_list.iter().for_each(|&left| {
        let similarity = right_list.iter().filter(|&&right| right == left).count();
        score += left * similarity;
    });

    score
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

    #[test]
    fn test_part2() {
        let input = Input::new("inputs/test.txt");
        let p2 = part2(&input);
        assert_eq!(p2, 31)
    }
}
