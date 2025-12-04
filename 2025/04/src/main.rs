#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() {
    let p1 = part1(DATA);
    println!("Part 1: {p1}");

    let p2 = part2(DATA);
    println!("Part 2: {p2}");
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

const PAPER_ROLL: char = '@';

fn reachable(grid: &[Vec<char>], row: usize, col: usize) -> bool {
    let rows = grid.len();
    let cols = grid[0].len();

    let mut neighbour_rolls: usize = 0;
    let directions = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];

    for (r, c) in directions {
        #[allow(clippy::cast_possible_wrap, clippy::cast_possible_truncation)]
        let target_row = usize::try_from(row as i32 + r).ok();
        #[allow(clippy::cast_possible_wrap, clippy::cast_possible_truncation)]
        let target_col = usize::try_from(col as i32 + c).ok();

        if let Some(rr) = target_row
            && rr < rows
            && let Some(cc) = target_col
            && cc < cols
        {
            neighbour_rolls += usize::from(grid[rr][cc] == PAPER_ROLL);
        }
    }

    neighbour_rolls < 4
}

fn part1(input: &str) -> String {
    let grid: Vec<Vec<char>> = input
        .trim()
        .lines()
        .map(|line| line.chars().collect())
        .collect();

    let mut reachable_rolls: usize = 0;

    for row in 0..grid.len() {
        for col in 0..grid[0].len() {
            let is_roll = grid[row][col] == PAPER_ROLL;
            reachable_rolls += usize::from(is_roll && reachable(&grid, row, col));
        }
    }

    reachable_rolls.to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "13");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

fn part2(input: &str) -> String {
    let mut result: isize = 0;

    result.to_string()
}

#[test]
fn part2_test() {
    assert_eq!(part2(TEST), "42");
}
