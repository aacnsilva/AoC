import fs from 'fs';

function processInstructions(input: string): number {
  const regex = /mul\((\d+),(\d+)\)/g;

  let match: RegExpExecArray | null;
  let sum = 0;

  while ((match = regex.exec(input)) !== null) {
    const x = parseInt(match[1], 10);
    const y = parseInt(match[2], 10);
    sum += x * y;
  }

  return sum;
}

// const input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";
const input = fs.readFileSync("input.txt").toString();
console.log(processInstructions(input));
