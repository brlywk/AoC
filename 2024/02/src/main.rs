use common::prelude::*;

fn main() {
    println!("AoC 2024 - Day 2\n");

    let input = Input::new("inputs/input.txt");

    println!("Part1: {}", part1(&input));
    // println!("Part2: {}", part2(&input));
}

fn part1(input: &Input) -> usize {
    input
        .lines
        .iter()
        .filter(|line| {
            let num_line: Vec<usize> = line
                .split_whitespace()
                .map(|s| s.parse::<usize>().unwrap())
                .collect();

            safe_report(&num_line, &[special_diff, slope])
        })
        .count()
}

// helper

fn safe_report(level: &Vec<usize>, checks: &[fn(&Vec<usize>) -> bool]) -> bool {
    checks.iter().all(|check| check(&level))
}

fn special_diff(level: &Vec<usize>) -> bool {
    level
        .windows(2)
        .all(|w| w[0].abs_diff(w[1]) >= 1 && w[0].abs_diff(w[1]) <= 3)
}

fn slope(level: &Vec<usize>) -> bool {
    level.windows(2).all(|w| w[0] > w[1]) || level.windows(2).all(|w| w[0] < w[1])
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = Input::new("inputs/test.txt");
        let p1 = part1(&input);

        assert_eq!(p1, 2);
    }

    // helper tests

    #[test]
    fn test_special_diff_pass() {
        let tests: Vec<Vec<usize>> = vec![
            vec![1, 2, 3, 4, 5],
            vec![1, 3, 5, 7, 9],
            vec![1, 2, 4, 6, 7, 9],
        ];

        for test in tests {
            assert!(special_diff(&test));
        }
    }

    #[test]
    fn test_special_diff_fail() {
        let tests: Vec<Vec<usize>> = vec![vec![1, 3, 4, 4, 8], vec![1, 5, 9, 13], vec![1, 1, 1, 1]];

        for test in tests {
            assert!(!special_diff(&test));
        }
    }

    #[test]
    fn test_slope_pass() {
        let tests: Vec<Vec<usize>> = vec![vec![1, 2, 3, 4, 5], vec![5, 4, 3, 2, 1]];

        for test in tests {
            assert!(slope(&test));
        }
    }

    #[test]
    fn test_slope_fail() {
        let tests: Vec<Vec<usize>> = vec![
            vec![1, 2, 3, 2, 4],
            vec![5, 4, 3, 4, 2],
            vec![1, 1, 2, 2, 3, 3],
        ];

        for test in tests {
            assert!(!slope(&test));
        }
    }

    #[test]
    fn test_safe_report() {
        let should_pass: Vec<usize> = vec![7, 6, 4, 2, 1];
        let should_fail: Vec<usize> = vec![1, 2, 7, 8, 9];

        assert!(safe_report(&should_pass, &[special_diff, slope]));
        assert!(!safe_report(&should_fail, &[special_diff, slope]));
    }
}
