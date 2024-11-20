# perftests

## Get the data files
the data files are avilable @ 
https://drive.google.com/open?id=0B6Tbo6PE8yPeVkFpd2ZXZU5yXzA

get all files under in the above link to ./data folder
so the ./data folder should contain names.txt
the data folder should be under the root of the project.


## performance data
| code                                   | env    | language           | version         | mem (mb)     | time (ms) |
|----------------------------------------|--------|--------------------|-----------------|--------------|-----------|
| filter 18 M names                      | mac os | go                 | 1.8             | 700          | 352       |
|                                        |        | java               | 1.8             | 1024         | 353       |
|                                        |        | dotnet             | 1.0.4(release)  | 600          | 762       |
|                                        |        | node.js            | 10              | 500          | 656       |
|                                        |        | python             | 2.7             | 800          | 1600      |
|                                        |        | python             | 3.6             | 1024         | 903       |
| sum(1 M array)                         |        | go                 | 1.8             |              | 70        |
|                                        |        | java               | 1.8             |              | 160       |
|                                        |        | dotnet             | 1.0.4 (release) |              | 280       |
|                                        |        | node.js            | 8               |              | 1200      |
|                                        |        | python             | 2.7             |              | 12000     |
|                                        |        | python             | 3.6             |              | 15000     |
| Query Against MongoDB (-c 100 -n 100k) |        | go                 | 1.10            | 11 cpu : 0%  | 368 s     |
|                                        |        | node.js            | 10.6            | 110 cpu: 20% | 321 s     |
|                                        |        | java (spring boot) | 8               | 925 cpu: 20% | 300 s     |
