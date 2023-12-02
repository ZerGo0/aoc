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

    let output = lines
        .map(|line| {
            let mut words = line.split_whitespace();
            let game_id = words
                .nth(1)
                .unwrap()
                .trim_end_matches(':')
                .parse::<usize>()
                .unwrap();

            let colors = &mut [0, 0, 0];
            let mut current_count = 0;
            for (i, word) in words.enumerate() {
                if i % 2 == 0 {
                    current_count = word.parse::<usize>().unwrap();
                } else {
                    let color = match word.chars().nth(0).unwrap() {
                        'r' => 0,
                        'g' => 1,
                        'b' => 2,
                        _ => panic!("invalid color"),
                    };

                    if current_count > 12 + color {
                        return 0;
                    }

                    colors[color] += current_count;
                }
            }

            return game_id;
        })
        .sum();

    return output;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    let output = lines
        .map(|line| {
            let words = line.split_whitespace();

            let colors = &mut [0, 0, 0];
            let mut current_count = 0;
            for (i, word) in words.skip(2).enumerate() {
                if i % 2 == 0 {
                    current_count = word.parse::<usize>().unwrap();
                } else {
                    let color = match word.chars().nth(0).unwrap() {
                        'r' => 0,
                        'g' => 1,
                        'b' => 2,
                        _ => panic!("invalid color"),
                    };

                    if current_count > colors[color] {
                        colors[color] = current_count;
                    }
                }
            }

            return colors.iter().product::<usize>();
        })
        .sum();

    return output;
}

#[allow(dead_code)]
const TEST_INPUT: &str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

#[test]
fn test_part_one() {
    let out = part_one(Some(TEST_INPUT));

    assert_eq!(out, 8);
    println!("test_part_one: {}", out);
}

#[allow(dead_code)]
const TEST_INPUT2: &str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

#[test]
fn test_part_two() {
    let out = part_two(Some(TEST_INPUT2));

    assert_eq!(out, 2286);
    println!("test_part_two: {}", out)
}
