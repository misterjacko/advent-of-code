from datetime import datetime

def readFile(filename):
  datalist = []
  datafile = open(filename, "r")
  for aline in datafile:
    datalist.append(list(aline.strip()))
  datafile.close()
  return datalist
  
def findacrossindex (acrossindex, right, setindexlength):
  position = acrossindex + right
  while position > setindexlength:
    position = position - setindexlength - 1
  return position

def part1(A):
  setindexlength = len(A[0])-1  
  right = 3
  acrossindex = 0
  treesencountered = 0
  for i in A:
    if i[acrossindex] == "#":
      treesencountered += 1
    acrossindex = findacrossindex(acrossindex,right,setindexlength)  
  return (treesencountered)

def part2(A):
  setindexlength = len(A[0])-1  
  right = [1,3,5,7,1]
  down = [1,1,1,1,2]
  treeproduct = 1
  i=0
  while i < len(right):
    treesencountered = 0
    acrossindex = 0
    j = 0
    while j < len(A):
      if A[j][acrossindex] == "#":
        treesencountered += 1
      acrossindex = findacrossindex(acrossindex,right[i],setindexlength)
      j += down[i]
    treeproduct *= treesencountered 
    i += 1
  return (treeproduct)

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 7
  assert part2(test_input) == 336

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)