const isSquare = num => Math.floor(Math.sqrt(num)) === Math.sqrt(num);
const isOdd = num => num % 2 !== 0;
const getPathKey = (x, y) => `${x}:${y}`;

const checkNodeValue = (map, x, y) => {
  let possiblePaths = [];
  let val = 0;
  for (let i = -1; i <= 1; i++) {
    for (let j = -1; j <= 1; j++) {
      if (i === 0 && j === 0) continue;
      path = getPathKey(x + i, y + j);
      if (map[path]) {
        val += map[path];
      }
    }
  }
  return val;
};

const getAnswer = breakFn => {
  let x = 0;
  let y = 0;
  let bound = 0;
  let index = 1;
  let incrementVal = 1;
  let map = {};

  while (true) {
    if (isOdd(index) && isSquare(index)) {
      bound++;
    }

    let nodeValue = checkNodeValue(map, x, y);
    let path = getPathKey(x, y);

    if (breakFn(nodeValue, index, x, y)) {
      break;
    }

    map[path] = nodeValue === 0 ? 1 : nodeValue;

    if (Math.abs(x + incrementVal) <= bound) {
      x += incrementVal;
      index++;
      continue;
    }
    if (Math.abs(y + incrementVal) <= bound) {
      y += incrementVal;
      index++;
      continue;
    }
    incrementVal = -incrementVal;
  }
};

let input = 289326;

getAnswer((nodeValue, nodeIndex, x, y) => {
  if (nodeIndex === input) {
    console.log(`part one: ${Math.abs(x) + Math.abs(y)}`);
    return true;
  }
});

getAnswer((nodeValue, nodeIndex, x, y) => {
  if (nodeValue > input) {
    console.log(`part two: ${nodeValue}`);
    return true;
  }
});
