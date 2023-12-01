from collections import deque, defaultdict
from datetime import datetime

def startCounting(A, countTo):
  turndict = defaultdict(list)
  nextNumber = 0
  for turn in range(1,countTo):
    if turn <= len(A):
      turndict[A[turn-1]] = deque([turn], maxlen= 2)
      nextNumber = 0
    else:
      if nextNumber in turndict:
        turndict[nextNumber].append(turn)
        nextNumber = turndict[nextNumber][1] - turndict[nextNumber][0]\
      else:
        turndict[nextNumber] = deque([turn], maxlen= 2)
        nextNumber = 0
  return (nextNumber)

def part1(A):
  return startCounting(A, 2020)

def part2(A):
  return startCounting(A, 30000000)

def test():
  assert part1([0,3,6]) == 436
  assert part1([1,3,2]) == 1
  assert part1([2,1,3]) == 10
  assert part1([1,2,3]) == 27
  assert part1([2,3,1]) == 78
  assert part1([3,2,1]) == 438
  assert part1([3,1,2]) == 1836
  # assert part2([0,3,6]) == 175594
  # assert part2([1,3,2]) == 2578
  # assert part2([2,1,3]) == 3544142
  # assert part2([1,2,3]) == 261214
  # assert part2([2,3,1]) == 6895259
  # assert part2([3,2,1]) == 18
  # assert part2([3,1,2]) == 362

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  print(f"Part 1: {part1([18,11,9,0,5,1])}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2([18,11,9,0,5,1])}")
  print (datetime.now()-starttime)