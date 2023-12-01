from datetime import datetime

def readFile(filename):
  dataset = []
  datafile = open(filename, "r")
  for aline in datafile:
      dataset.append(aline.strip('\n'))
  datafile.close()
  return dataset

def findrow(rowstring):
  rowstring = rowstring.replace('B', '1')
  binaryrow = rowstring.replace('F', '0')
  rownumb = int(binaryrow, 2)
  return(rownumb)

def findseat(seatstring):
  seatstring = seatstring.replace('R', '1')
  binaryseat = seatstring.replace('L', '0')
  seatnumb = int(binaryseat, 2)
  return(seatnumb)

def part1(A):
  highseat = 0
  for boardingpass in A:
    row = findrow(boardingpass[0:7])
    seat = findseat(boardingpass[7:10])
    tempseat = (row * 8) + seat
    if tempseat > highseat:
      highseat = tempseat
  return (highseat)

def part2(A):
  seatdict = {}
  for boardingpass in A:
    row = findrow(boardingpass[0:7])
    seat = findseat(boardingpass[7:10])
    tempseat = (row * 8) + seat
    seatdict[tempseat] = "@"
  iminus1 = -1
  #will ignore the first stretch of missing seats and return the first missing
  #seat after the first occupied seat.
  for i in range(0,885):     
    if i not in seatdict:
      if i - 1 != iminus1:
        return i
      iminus1 += 1
  return ("error")

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 820
  #assert part2(test_input) == #no test

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)