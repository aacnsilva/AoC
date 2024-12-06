isSafe :: [Int] -> Bool
isSafe levels =
  let differences = zipWith (-) (tail levels) levels
      allIncreasing = all (> 0) differences
      allDecreasing = all (< 0) differences
      withinRange = all (\d -> abs d >= 1 && abs d <= 3) differences
   in (allIncreasing || allDecreasing) && withinRange

isSafeWithRemoval :: [Int] -> Bool
isSafeWithRemoval levels = any (isSafe . removeAtIndex levels) [0 .. length levels - 1]
  where
    removeAtIndex xs i = take i xs ++ drop (i + 1) xs

countSafeReports :: [[Int]] -> Int
countSafeReports = length . filter (\report -> isSafe report || isSafeWithRemoval report)

main :: IO ()
main = do
  input <- readFile "input.txt"
  let reports = map (map read . words) (lines input)
  print $ countSafeReports reports
