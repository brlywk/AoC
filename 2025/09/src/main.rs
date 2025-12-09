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

struct RedTile {
    row: u32,
    col: u32,
}

impl RedTile {
    fn parse(s: &str) -> Option<Self> {
        let (c, r) = s.split_once(',')?;

        Some(Self {
            row: r.parse::<u32>().ok()?,
            col: c.parse::<u32>().ok()?,
        })
    }

    fn area(&self, other: &Self) -> u64 {
        let w = u64::from(self.col).abs_diff(u64::from(other.col)) + 1;
        let h = u64::from(self.row).abs_diff(u64::from(other.row)) + 1;

        w * h
    }
}

fn parse_input(input: &str) -> Vec<RedTile> {
    input.trim().lines().filter_map(RedTile::parse).collect()
}

fn part1(input: &str) -> String {
    let mut max_area: u64 = 0;

    let red_tiles = parse_input(input);

    for i in 0..red_tiles.len() {
        for j in (i + 1)..red_tiles.len() {
            let a = red_tiles[i].area(&red_tiles[j]);
            max_area = max_area.max(a);
        }
    }

    max_area.to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "50");
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
