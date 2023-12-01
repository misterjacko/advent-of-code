
from datetime import datetime
import math
from math import gcd
import re
from collections import deque

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [x.strip().split() for x in lines]

def parseData(freshlist):
  newlist = []
  newmask = []
  tempdict = {}
  for line in freshlist:
    if 'mask' in line:
      newlist.append(newmask)
      newmask = []
      tempdict = {}
      newmask.append(line[2])
    else:
      tempdict[int(re.findall(r'\d+', line[0])[0])] = int(line[2])
      if len(newmask) == 2:
        newmask[1] = tempdict
      else:
        newmask.append(tempdict)
  else:
    newlist.append(newmask)
    newlist = deque(newlist)
    newlist.popleft()
  return newlist

def part1(A):
  parsed = parseData(A)
  instructionmap = {}
  for mask in parsed:
    masklist = list(mask[0])
    masklist.reverse()
    for key in mask[1]:
      binval = bin(mask[1][key]).replace('0b', '')
      memlist = list(binval)
      memlist.reverse()
      while len(memlist) < len(masklist):
        memlist.append('0')
      for i, val in enumerate(masklist):
        if val != 'X':
          memlist[i] = val
      memlist.reverse()
      newbinval = ''.join(memlist)
      instructionmap[key] = int(newbinval, 2)
  total = 0
  for key in instructionmap:
    total += instructionmap[key]
  return total

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