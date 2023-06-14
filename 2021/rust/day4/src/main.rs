#![feature(iter_array_chunks)]

fn main() {
    let part = std::env::args().nth(1).expect("no part provided");

    let out = match part.as_str() {
        "1" => part_one(None),
        "2" => part_two(None),
        _ => panic!("invalid part provided"),
    };

    println!("Part {}: {}", part, out);
}

fn part_one(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    let moves = lines
        .clone()
        .into_iter()
        .nth(0)
        .unwrap()
        .split(",")
        .map(|x| x.parse::<usize>().unwrap())
        .collect::<Vec<usize>>();

    let mut boards = lines
        .skip(2)
        .filter(|line| line.len() > 0)
        .map(|line| {
            line.split_whitespace()
                .map(|x| (x.parse::<usize>().unwrap(), false))
                .collect::<Vec<(usize, bool)>>()
        })
        .array_chunks::<5>()
        .collect::<Vec<_>>();

    for m in moves {
        for (idx, b) in boards.clone().iter().enumerate() {
            for (i, row) in b.iter().enumerate() {
                for (j, col) in row.iter().enumerate() {
                    if col.0 == m {
                        boards[idx][i][j].1 = true;

                        if boards[idx][i].iter().all(|x| x.1)
                            || boards[idx].iter().map(|x| x[j]).all(|x| x.1)
                        {
                            let sum = boards[idx]
                                .iter()
                                .flatten()
                                .filter(|x| !x.1)
                                .map(|x| x.0)
                                .sum::<usize>();

                            return sum * m;
                        }
                    }
                }
            }
        }
    }

    return 0;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    let moves = lines
        .clone()
        .into_iter()
        .nth(0)
        .unwrap()
        .split(",")
        .map(|x| x.parse::<usize>().unwrap())
        .collect::<Vec<usize>>();

    let mut boards = lines
        .skip(2)
        .filter(|line| line.len() > 0)
        .map(|line| {
            line.split_whitespace()
                .map(|x| (x.parse::<usize>().unwrap(), false))
                .collect::<Vec<(usize, bool)>>()
        })
        .array_chunks::<5>()
        .collect::<Vec<_>>();

    let mut last_won_board_move: Option<([Vec<(usize, bool)>; 5], usize)> = None;
    let mut won_boards = vec![false; boards.len()];

    for m in moves {
        for (idx, b) in boards.clone().iter().enumerate() {
            if won_boards[idx] {
                continue;
            }

            for (i, row) in b.iter().enumerate() {
                for (j, col) in row.iter().enumerate() {
                    if col.0 == m {
                        boards[idx][i][j].1 = true;

                        if boards[idx][i].iter().all(|x| x.1)
                            || boards[idx].iter().map(|x| x[j]).all(|x| x.1)
                        {
                            won_boards[idx] = true;
                            last_won_board_move = Some((boards[idx].clone(), m));
                        }
                    }
                }
            }
        }
    }

    if last_won_board_move.is_none() {
        return 0;
    }

    let last_won_board_move = last_won_board_move.unwrap();

    let sum = last_won_board_move
        .0
        .iter()
        .flatten()
        .filter(|x| !x.1)
        .map(|x| x.0)
        .sum::<usize>();

    return sum * last_won_board_move.1;
}

#[allow(dead_code)]
const TEST_INPUT: &str = "7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7";

#[test]
fn test_part_one() {
    let out = part_one(Some(TEST_INPUT));

    assert_eq!(out, 4512);
    println!("test_part_one: {}", out);
}

#[test]
fn test_part_two() {
    let out = part_two(Some(TEST_INPUT));

    assert_eq!(out, 1924);
    println!("test_part_two: {}", out)
}
