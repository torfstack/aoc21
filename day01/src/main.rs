use std::fs;

fn main() {
    let lines = read_input();
    let mut current = 99_999;
    let mut count = 0;
    lines[..lines.len()-1]
        .iter()
        .map(|line| line.parse::<u32>().unwrap())
        .for_each(|n| {
            if n > current {
                count = count+1;
            }
            current = n
        });

    println!("{}", count)
}

fn read_input() -> Vec<String> {
    let read = fs::read_to_string("input").unwrap();
    println!("{}", read);
    return read
        .split("\n")
        .map(|l| l.to_string())
        .collect()
}
