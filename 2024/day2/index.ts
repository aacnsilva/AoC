import fs from 'fs';

function isSafe(levels: number[]): boolean {
  const differences = levels.slice(1).map((level, index) => level - levels[index]);
  const allIncreasing = differences.every(diff => diff > 0);
  const allDecreasing = differences.every(diff => diff < 0);
  const withinRange = differences.every(diff => Math.abs(diff) >= 1 && Math.abs(diff) <= 3);
  return (allIncreasing || allDecreasing) && withinRange;
}

function countSafeReports(reports: number[][]): number {
  return reports.filter(isSafe).length;
}

function solve1(input: string): void {
  const reports = input.trim().split("\n").map(line => line.split(" ").map(Number));
  console.log(countSafeReports(reports));
}

// const input = `7 6 4 2 1
// 1 2 7 8 9
// 9 7 6 2 1
// 1 3 2 4 5
// 8 6 4 4 1
// 1 3 6 7 9`;
const input = fs.readFileSync("input.txt").toString();
solve1(input);
