use std::fs::File;
use std::io::{BufRead, BufReader};
use std::time::Instant;

fn main() {
    let file_name = "../../data/names.txt";

    let file =
        File::open(file_name).expect("Unable to read data file. Please check README.msd file");
    let reader = BufReader::new(file);

    let names: Vec<String> = reader
        .lines()
        .filter_map(|line| match line {
            Ok(line_content) => Some(line_content),
            Err(_) => None,
        })
        .collect();

    let original_count = names.len();

    let start = Instant::now();

    let names: Vec<_> = names.iter().filter(|name| name.contains("a")).collect();

    let elapsed = start.elapsed();

    println!("original count is {}", original_count);
    println!("filtered count is {}", names.len());
    println!("{:?}", elapsed);
}
