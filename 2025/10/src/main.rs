use std::collections::{HashMap, VecDeque};

#[allow(dead_code)]
const TEST: &str = include_str!("./test.txt");
const DATA: &str = include_str!("./data.txt");

fn main() {
    let p1 = part1(DATA);
    println!("Part 1: {p1}");

    // let p2 = part2(DATA);
    // println!("Part 2: {p2}");
}

//////////////////////////////////////////////////
// Part 1
//////////////////////////////////////////////////

#[derive(Debug)]
struct Machine {
    lights: Vec<bool>,
    buttons: Vec<Vec<usize>>,
    // joltages: Vec<usize>,
}

impl Machine {
    fn parse(s: &str) -> Option<Self> {
        let (lights_str, rest) = s.split_once(' ')?;
        let (buttons_str, _joltages_str) = rest.rsplit_once(' ')?;

        Some(Self {
            lights: lights_str
                .trim_matches(|c| c == '[' || c == ']')
                .chars()
                .map(|c| c == '#')
                .collect(),
            buttons: buttons_str
                .split(' ')
                .map(|button| {
                    button
                        .trim_matches(|c| c == '(' || c == ')')
                        .split(',')
                        .filter_map(|s| s.parse::<usize>().ok())
                        .collect()
                })
                .collect(),
            // joltages: joltages_str
            //     .trim_matches(|c| c == '{' || c == '}')
            //     .split(',')
            //     .filter_map(|s| s.parse::<usize>().ok())
            //     .collect(),
        })
    }

    // using every last bit of my brain power to recognize: everything is a graph problem!
    // - light states are nodes
    // - button presses are edges
    // - "depth" of the graph is number of button presses
    //
    // minimum presses should be: find the correct state in the "lowest depth" = BFS
    fn light_minimum_buttons(&self) -> usize {
        let l = self.lights.len();
        // create a bitmask of our target state, e.g.:
        // [.##.] = [false, true, true, false] = [0, 1, 1, 0] = 0110 = 6
        // makes it easier to compare states (single number)
        let target_mask: usize = self
            .lights
            .iter()
            .enumerate()
            .fold(0, |acc, (i, &b)| acc | (usize::from(b) << i));

        // save nodes as (state, "depth = button presses") to know which "depth" we
        // are currently operating in:
        // [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1)
        //                                   (0000=0, 0)
        //                /                     |                  \
        // (3)->(1000=8, 1)    ...    (0,2)->(0101=5, 1)    ...
        //                                   /
        // ...                 ...    (0,1)->(0110=6, 2)    <- Our target state!
        let mut queue = VecDeque::from([(0, 0)]);
        let mut visited = HashMap::from([(0, 0)]);

        // BFS: pop from the front
        while let Some((state, presses)) = queue.pop_front() {
            // if the current node is our target_mask: found it!
            if state == target_mask {
                return presses;
            }

            // check children of current node by pressing all the buttons!
            for button in &self.buttons {
                let mut new_state = state;
                for &light_idx in button {
                    if light_idx < l {
                        // toggle (XOR) light:
                        // each light can be "targeted" by shifting the "toggle" there:
                        // 1 << idx   in bits     number
                        // 1 << 0     0001        1
                        // 1 << 1     0010        2
                        // 1 << 2     0100        4
                        // etc.
                        //
                        // so we essentially take the current state,
                        // e.g. from state "3" = 0011, and when we press button [1,3], we do:
                        // 0011 ^= (1 << 1)   => 0001
                        // 0001 ^= (1 << 3)   => 1001
                        // new_state now is "9" = 1001
                        new_state ^= 1 << light_idx;
                    }
                }

                if !visited.contains_key(&new_state) || visited[&new_state] > presses + 1 {
                    visited.insert(new_state, presses + 1);
                    queue.push_back((new_state, presses + 1));
                }
            }
        }

        panic!("not solvable")
    }

    // pretty much the same as above, but now we don't need to fiddle with bitmasks and can just
    // keep pressing all the buttons!
    //
    // ...aaaaaand as always inputs for AoC are too large for an actual easy modification of an
    // existing solution; I could now waste a whole lot of time to rewrite this, or I just accept
    // that I find this kind of "obscure optimization problem" rather boring and move on :P
    // fn joltage_minimum_buttons(&self) -> usize {
    //     let n = self.joltages.len();
    //     let target = &self.joltages;
    //
    //     let mut queue = VecDeque::new();
    //     let mut visited = HashMap::new();
    //
    //     let start = vec![0; n];
    //     queue.push_back((start.clone(), 0));
    //     visited.insert(start, 0);
    //
    //     while let Some((state, presses)) = queue.pop_front() {
    //         if state == *target {
    //             return presses;
    //         }
    //
    //         for button in &self.buttons {
    //             let mut new_state = state.clone();
    //
    //             for &idx in button {
    //                 if idx < n {
    //                     new_state[idx] += 1;
    //                 }
    //             }
    //
    //             // important: as we can (and should!) keep pressing buttons forever, we
    //             // unfortunately still don't need to press a button more if the associated
    //             // joltage is exceeded :(
    //             if new_state
    //                 .iter()
    //                 .zip(target)
    //                 .all(|(curr, &tgt)| curr <= &tgt)
    //                 && !visited.contains_key(&new_state)
    //             {
    //                 visited.insert(new_state.clone(), presses + 1);
    //                 queue.push_back((new_state, presses + 1));
    //             }
    //         }
    //     }
    //
    //     panic!("not solvable")
    // }
}

fn parse_input(input: &str) -> Vec<Machine> {
    input.trim().lines().filter_map(Machine::parse).collect()
}

fn part1(input: &str) -> String {
    parse_input(input)
        .iter()
        .map(Machine::light_minimum_buttons)
        .sum::<usize>()
        .to_string()
}

#[test]
fn part1_test() {
    assert_eq!(part1(TEST), "7");
}

//////////////////////////////////////////////////
// Part 2
//////////////////////////////////////////////////

// fn part2(input: &str) -> String {
//     parse_input(input)
//         .iter()
//         .map(Machine::joltage_minimum_buttons)
//         .sum::<usize>()
//         .to_string()
// }
//
// #[test]
// fn part2_test() {
//     assert_eq!(part2(TEST), "33");
// }
