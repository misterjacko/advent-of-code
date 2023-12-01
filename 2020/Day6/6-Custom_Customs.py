from datetime import datetime

def combinelines(tempdata):
  onelinedict = {'total': len(tempdata)}
  for i in range(len(tempdata)):
    for j in range(len(tempdata[i])):
      if tempdata[i][j] in onelinedict:
       onelinedict[tempdata[i][j]] += 1
      else:
       onelinedict[tempdata[i][j]] = 1
  return onelinedict

def readFile(filename):
  dataset = []
  datafile = open(filename, "r")
  tempdata = []
  for aline in datafile:
    if aline != "\n":
      tempdata.append(aline.strip())
    else: 
      dataset.append(combinelines(tempdata))
      tempdata = []
  dataset.append(combinelines(tempdata))
  datafile.close()
  return dataset

def part1(A):
  total = 0
  for i in range(len(A)):
    total += len(A[i]) -1
  return (total)

def part2(A):
  semitotal = 0
  total = 0
  for i in range(len(A)):
    for j in A[i]:
      if A[i]['total']  == A[i][j]:
        semitotal += 1
    total += semitotal -1
    semitotal = 0
  return (total)

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 11
  assert part2(test_input) == 6

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)