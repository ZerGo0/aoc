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

    let output = lines.fold((0, 0), |(hoz_pos, depth), line| {
        let split = line.split(" ").collect::<Vec<&str>>();
        let dir = split[0];
        let value = split[1].parse::<usize>().unwrap();

        match dir {
            "forward" => (hoz_pos + value, depth),
            "up" => (hoz_pos, depth - value),
            "down" => (hoz_pos, depth + value),
            _ => panic!("invalid direction"),
        }
    });

    return output.0 * output.1;
}

fn part_two(test_input: Option<&str>) -> usize {
    let input = test_input.unwrap_or(include_str!("../input.txt"));

    let lines = input.lines();

    let output = lines.fold((0, 0, 0), |(hoz_pos, depth, aim), line| {
        let split = line.split(" ").collect::<Vec<&str>>();
        let dir = split[0];
        let value = split[1].parse::<usize>().unwrap();

        match dir {
            "forward" => (hoz_pos + value, depth + aim * value, aim),
            "up" => (hoz_pos, depth, aim - value),
            "down" => (hoz_pos, depth, aim + value),
            _ => panic!("invalid direction"),
        }
    });

    return output.0 * output.1;
}

#[test]
fn test_part_one() {
    let test_input = "forward 5
down 5
forward 8
up 3
down 8
forward 2";

    let out = part_one(Some(test_input));

    assert_eq!(out, 150);
    println!("test_part_one: {}", out);
}

#[test]
fn test_part_two() {
    let test_input = "forward 5
down 5
forward 8
up 3
down 8
forward 2";

    let out = part_two(Some(test_input));

    assert_eq!(out, 900);
    println!("test_part_two: {}", out)
}
