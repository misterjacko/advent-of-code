from datetime import datetime
import copy

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
    return [x.strip().split(' ') for x in lines]


def part1(A):
  accval = 0
  hit = set()
  i = 0
  while i not in hit and i < len(A):
    hit.add(i)
    if 'nop' in A[i]:
      i += 1
    elif 'acc' in A[i]:
      accval += int(A[i][1])
      i += 1
    elif 'jmp' in A[i]:
      i += int(A[i][1])
  return accval

def totheend(B):
  hit = set()
  i = 0
  while i < len(B):
    if i not in hit:
      hit.add(i)
      if 'jmp' in B[i]:
        i += int(B[i][1])
      else:
        i += 1
      continue
    return False
  return True


def part2(A):
  for line in range(len(A)):
    if A[line][0] == 'nop':
      B = copy.deepcopy(A)
      B[line][0] = 'jmp'
      if totheend(B) == True:
        return part1(B)
    elif A[line][0] == 'jmp':
      B = copy.deepcopy(A)
      B[line][0] = 'nop'
      if totheend(B) == True:
        return part1(B)
    else:
      continue

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 5
  assert part2(test_input) == 8


if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)