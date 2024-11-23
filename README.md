# perftests


# Programs
## filter-large-data
This program reads a large number of lines from a file. It then filters the items with the ones have 'a'. The time is measured for the filtering these items of ~18 Million.

# Perf numbers
| Program | Go | Rust |
|---------|----|------|
|filter-large-data|115 ms|53ms|


## Environment
### Machine
Mac Mini M4 16GB

### Versions
Go - 1.23
Rust - 1.82


# How to run
## Get the data files
the data files are avilable @ 
https://drive.google.com/open?id=0B6Tbo6PE8yPeVkFpd2ZXZU5yXzA

get all files under in the above link to ./data folder
so the ./data folder should contain names.txt
the data folder should be under the root of the project.

## Run filter-large-data
### Go
cd filter-large-data/gofilter
go build -o target/
./target/gofilter

### Rust
cd filter-large-data/gofilter
cargo build --release
./target/release/rustfilter


# Performance Results
## Filter Large Data
#### Go        - 115.6075 ms
#### Rust      - 53.096458 ms
