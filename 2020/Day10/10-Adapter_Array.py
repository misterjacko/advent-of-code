
from datetime import datetime
from collections import Counter
import math

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [int(x.strip()) for x in lines]

def part1(A):
  A.sort()
  currentjolts = 0
  step1 = 0
  step3 = 0
  for i in A:
    if i == currentjolts + 1:
      step1 += 1
    elif i == currentjolts + 2:
      continue
    elif i == currentjolts + 3:
      step3 += 1   
    else:
      break
    currentjolts = i
  step3 += 1
  return (step1 * step3)
  return (A)

def diff(data):
    return [next_ - current for current, next_ in zip(data, data[1:])]

def get_arrangements(adapters):
    if not adapters:
        return 1
    count = 0
    for i, adapter in enumerate(adapters):
        if 0 <= adapter <= 3:
            count += get_arrangements(tuple([adapter_ - adapter for adapter_ in adapters[i+1:]]))
        else:
            break
    return count

def part2(A):
  A.sort()
  #A.append(0)
  print (A)
  A += [A[-1] + 3]
  print (A)
  A = [0] + A
  print (A)
  print (tuple(A[1:]))
  counts = Counter(diff(A))
  print (counts)
  pattern = 1

  return (pattern)

def test():
  test_input = (readFile("test.txt"))
  test_input2 = readFile("test2.txt")
  assert part1(test_input) == 35
  assert part1(test_input2) == 220
  #assert part2(test_input) == 8
  #assert part2(test_input2) == 19208

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  #print(f"Part 2: {part2(vals)}")
  #print (datetime.now()-starttime)