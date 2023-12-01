
from datetime import datetime
import math
from math import gcd

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [x.strip().split(',') for x in lines]

def parseData(freshlist):
  newlist = []
  for line in freshlist:
    if line != 'x':
     newlist.append(int(line))
    else:
      newlist.append(line)
  return newlist

def part1(A):
  time = int(A[0][0])
  bus = parseData(A[1])
  closestTime = time
  closestBus = ''
  for ID in bus:
    if ID != 'x':
      nextID = math.ceil(time/int(ID))*int(ID)
      if nextID - time < closestTime:
        closestTime = nextID - time
        closestBus = ID
  return int(closestBus)*closestTime

def nextPair(step, val1, val2, iter):
  val1 = step
  print (step)
  while (val1 + iter) % val2 != 0:
    val1 += step
    print (val1)
  return val1

def part2(A):
  bus = parseData(A[1])
  lastval = 0
  step = 1
  time = 0
  for idx, ID in enumerate(bus):
    if ID != 'x':
      time = nextPair(step, 0, ID, idx)
      print (time)
      if idx == 0:
        step = time
      else:
        step = time * (ID + idx)

  # for x in range(len(bus)):
  #   print (bus[x])
  #   if bus[x] != 'x':
  #     time = lcd(time, bus[x] - x)
  #   else:
  #     continue
  # while valid == False:
  #   valid = True
  #   time += int(bus[0])
  #   for x in range(len(bus)):
  #     if bus[x] != 'x':
  #       if (time + x) % int(bus[x])!= 0: #==0 when number is valid
  #         valid = False
  #         break
  # print('-----')
  print (time)
    
  return (time)

def test():
  test_input = (readFile("test.txt"))
  assert part1(test_input) == 295
  assert part2(test_input) == 1068781

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  #print(f"Part 2: {part2(vals)}")
  #print (datetime.now()-starttime)