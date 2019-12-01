const fs = require('fs');
const path = require('path');

const getAnswer = (str, offset) => {
  let digits = str.toString().split('');
  let sum = 0;

  for (let i = 0; i < digits.length; i++) {
    let current = 1 * digits[i];
    let next = 1 * digits[(i + offset) % digits.length];
    if (current === next) sum += current;
  }
  return sum;
};

const input = fs
  .readFileSync(path.join(__dirname, 'input'))
  .toString()
  .trim();

const partOne = getAnswer(input, 1);
const partTwo = getAnswer(input, input.length / 2);

console.log(`part one: ${partOne}`);
console.log(`part two: ${partTwo}`);
