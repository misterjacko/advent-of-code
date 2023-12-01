import re
from datetime import datetime


def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [x.strip() for x in lines]

def parseData(freshlist):
  newList = []
  for line in freshlist:
    if ':' not in line and ',' in line:
      newList.append(list(map(int, line.split(','))))

  print (newList)
  return newList

def part1(A):
  parseData(A)
  return A

def part2(A):
  return (A)

def test():
  test_input = (readFile("test.txt"))
  assert part1(test_input) == 165
  #assert part2(test_input) == 208

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  #print(f"Part 2: {part2(vals)}")
  #print (datetime.now()-starttime)