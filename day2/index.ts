import fs from 'fs';

const solve1 = (input: string) => {
  const lines = input.split("\n");
  const safeReports = lines.map(row => {
    const values = row.split(" ");
    const safe = reportIsSafe(values);
    console.log(`Report: ${values} - ${safe}`);
    return safe;
  });
  return safeReports.filter(el => el === true).length
};

const reportIsSafe = (numbers: number[]) => {
  if (!numbers || numbers.length < 1 || numbers[0] === numbers[1])
    return false;
  const direction = numbers[0] < numbers[1] ? '+' : '-';
  for (let i = 1; i < numbers.length; i++) {
    const curr = numbers[i];
    const previous = numbers[i - 1];
    if (Math.abs(previous - curr) > 3)
      return false;

    if (direction === '+' && previous >= curr)
      return false;

    if (direction === '-' && previous <= curr)
      return false;
  }
  return true;
};

// console.log(solve1(`7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9`));

console.log(solve1(fs.readFileSync("input.txt").toString()));
