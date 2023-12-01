from datetime import datetime

def readFile(filename):
  dataset = []
  datafile = open(filename, "r")
  for aline in datafile:
    lineset = {}
    line = aline.split(" ")
    outerbag =" ".join(line[:2])
    innerbag = {}
    for index in range(len(line)):
      if line[index].isnumeric():
        innerbag[" ".join(line[index + 1:index + 3])] = line[index]
    lineset[outerbag] = innerbag
    dataset.append(lineset)
  datafile.close()
  return dataset

def searchbags(querybags, A):
  valid = []
  for query in querybags:
    for outerbag in A:
      for innerbag in outerbag:
        if query in outerbag[innerbag]:
          valid.append(innerbag)
  return valid

def part1(A):
  acceptedbags = []
  querybags = ['shiny gold']
  while len(querybags) > 0:
    validbags = searchbags(querybags, A)
    querybags = []
    for bag in validbags:
      if bag not in acceptedbags:
        acceptedbags.append(bag)
        querybags.append(bag)
  return (len(acceptedbags))

#def fillbags():
  

def part2(A):
  return ('-')

def test():
  test_input = readFile("test.txt")
  test_input2 = readFile("test2.txt")
  assert part1(test_input) == 4
  #assert part2(test_input) == 32
  #assert part2(test_input2) == 126

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  #print(f"Part 2: {part2(vals)}")
  #print (datetime.now()-starttime)