
from datetime import datetime
from collections import Counter
import math
import copy

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [list(x.strip()) for x in lines]

def addendcaps(inputlist):
  frontcap = []
  endcap = []
  for i in range(len(inputlist[len(inputlist)-1])):
    endcap.append('0')
  for i in range(len(inputlist[0])):
    frontcap.append('0')
  inputlist.insert(0, frontcap)
  inputlist.append(endcap)

  return inputlist

def linecaps(linestring):
  linestring.insert(0, '0')
  linestring.append('0')
  #print (linestring)
  return linestring

def part1(A):
  adapted = (addendcaps(A))
  for lines in adapted:
    lines = linecaps(lines)
  nextroundlist = copy.deepcopy(adapted)
  keepGoing = True
  while keepGoing == True:
    keepGoing = False
    for idx, line in enumerate(adapted):
      if 'L' in line or '#' in line:
        neighbors = list(zip(adapted[idx-1].copy(), adapted[idx].copy(), adapted[idx+1].copy()))
        for i in range(len(line)):
          if line[i] == 'L':
            immediateneighbors = []
            for x in range(-1, 2):
              immediateneighbors.append(neighbors[i+x])
            occupied = (sum(x.count('#') for x in immediateneighbors))
            if occupied == 0:
              keepGoing = True
              nextroundlist[idx][i] = '#'
          elif line[i] == '#':
            immediateneighbors = []
            for x in range(-1, 2):
              immediateneighbors.append(neighbors[i+x])
            occupied = (sum(x.count('#') for x in immediateneighbors))
            if occupied >= 5:
              keepGoing = True
              nextroundlist[idx][i] = 'L'
    adapted = copy.deepcopy(nextroundlist)
  else:
    return ((sum(x.count('#') for x in adapted)))


def part2(A):

  pattern = 1

  return (pattern)

def test():
  test_input = (readFile("test.txt"))
  #test_input2 = readFile("test2.txt")
  assert part1(test_input) == 37
  #assert part1(test_input2) == 220
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