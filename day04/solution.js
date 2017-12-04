const fs = require('fs');
const path = require('path');

let maps = [];

const checkValidity = (phrases, getMapKey) => {
  let valid = 0;
  phraseLoop: for (let p = 0; p < phrases.length; p++) {
    let phrase = phrases[p];
    if (!phrase) continue;
    maps.push({});

    let words = phrase.split(' ');

    wordLoop: for (let w = 0; w < words.length; w++) {
      let map = maps[p];
      let mapKey = getMapKey(words[w]);

      if (map[mapKey]) {
        continue phraseLoop;
      }
      map[mapKey] = true;
      if (w === words.length - 1) {
        valid++;
      }
    }
  }
  return valid;
};

const input = fs.readFileSync(path.join(__dirname, 'input')).toString();
const phrases = input.split('\n');

const partOne = checkValidity(phrases, word => word);

const partTwo = checkValidity(phrases, word => word.split('').sort().join());

console.log(`part one: ${partOne}`);
console.log(`part two: ${partTwo}`);
