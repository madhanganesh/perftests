# perftests

## Get the data files
the data files are avilable @ 
https://drive.google.com/open?id=0B6Tbo6PE8yPeVkFpd2ZXZU5yXzA

get all the files under in the above link to ./data folder
so the ./data folder should contain names.txt
the data folder should be under the root of the project.


## performance data
| code               | env    | language | version | mem (mb) | time (ms) |
|--------------------|--------|----------|---------|----------|-----------|
| filter 18 M names  | mac os | go       | 1.7     | 700      | 350       |
|                    |        | java     | 1.8     | 1024     | 353       |
|                    |        | node.js  | 7       | 1024     | 2135      |
|                    |        | python   | 2.7     | 800      | 1600      |
|                    |        | python   | 3.6     | 1024     | 1000      |
| sum(1 M array)     |        | go       | 1.7     |          | 165       |
|                    |        | node.js  | 7.2.1   |          | 1200      |
|                    |        | python   | 2.7     |          | 12000     |
|                    |        | python   | 3.6     |          | 15000     |
