from datetime import datetime
from collections import deque
from itertools import combinations

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
    return [int(x.strip()) for x in lines]

def part1(A, preamble):
  d = deque([], maxlen = preamble)
  for i in range(0,preamble):
    d.append(int(A[i]))
  for i in range(preamble, len(A)):
    valid = False
    for val1, val2 in combinations(d, 2):
      if (val1 + val2) == int(A[i]):
        valid = True
        break
    if valid == False:
      return int(A[i])
    else:
      d.append(int(A[i]))

def part2(A, preamble):
  invalidnumber = part1(A, preamble)
  for i in range(len(A)):
    testlist = []
    for j in range(i,len(A)):
      testlist.append(A[j])
      if sum(testlist) >= invalidnumber:
        break
    if sum(testlist) == invalidnumber:
      testlist.sort()
      return (testlist[0] + testlist[len(testlist)-1])

def test():
  preamble = 5
  test_input = readFile("test.txt")
  assert part1(test_input, preamble) == 127
  assert part2(test_input, preamble) == 62


if __name__ == "__main__":
  preamble = 25
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals, preamble)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals, preamble)}")
  print (datetime.now()-starttime)