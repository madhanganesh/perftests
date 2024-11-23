const factors = value => {
  const result = [];
  for (let i = 1; i <= value; i++) {
    if (value % i === 0) {
      result.push(i);
    }
  }
  return result;
};

const input = 100000000; // 100 million 100,000,000
console.time('start');
const result = factors(input).length;
console.timeEnd('start');
console.log(result);
