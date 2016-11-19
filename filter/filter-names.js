const readline = require('readline');
const fs = require('fs');
let names = [];

const rl = readline.createInterface({
    input: fs.createReadStream('../data/names.txt')
});

rl.on('line', (line) => {
    names.push(line);
});

rl.on('close', () => {
    console.time('start');
    names = names.filter(name => name.indexOf('A') === -1);
    console.timeEnd('start');
    console.log(names.length);
});
