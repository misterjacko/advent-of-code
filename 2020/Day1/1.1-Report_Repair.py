from datetime import datetime

def readFile(filename):
    with open(filename, 'r') as f:
      lines = f.readlines()
    return [int(x) for x in lines]

def part1(A):
  for i in A:
    remainder = 2020 - i
    if remainder in A:
      return (i*remainder)
  return ("No Matches")

#this is brute force but the data set is small so :shrug:
def part2(A):
  for i in A:
    for j in A:
      remainder = 2020 - j - i
      if remainder in A:
        if j in A:
          return (i*j * remainder)
  return ("No Matches")

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 514579
  assert part2(test_input) == 241861950

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  transactions = readFile("day1data.txt")
  print(f"Part 1: {part1(transactions)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(transactions)}")
  print (datetime.now()-starttime)