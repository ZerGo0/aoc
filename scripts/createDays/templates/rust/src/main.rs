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

    return output;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    return output;
}

#[allow(dead_code)]
const TEST_INPUT: &str = "";

#[test]
fn test_part_one() {
    let out = part_one(Some(TEST_INPUT));

    assert_eq!(out, 0);
    println!("test_part_one: {}", out);
}

#[test]
fn test_part_two() {
    let out = part_two(Some(TEST_INPUT));

    assert_eq!(out, 0);
    println!("test_part_two: {}", out)
}
