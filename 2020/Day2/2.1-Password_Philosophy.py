from datetime import datetime
import re

def readFile(filename):
    passwords = []
    datafile = open(filename, "r")
    for aline in datafile:
        passwords.append(re.split('[-|:\s]', aline))
    datafile.close()
    return passwords

def part1(A):
    validcount = 0
    for i in A:
        count = i[4].count(i[2])
        if count >= int(i[0]) and count <= int(i[1]):
            validcount += 1
    return validcount

def part2(A):
    validcount = 0
    for i in A:
        index1 = int(i[0])-1
        index2 = int(i[1])-1
        passstring = i[4]
        if (passstring[index1] == i[2]) is not (passstring[index2] == i[2]):
            validcount += 1      
    return validcount

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 2
  assert part2(test_input) == 1

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  passwords = readFile("day2data.txt")
  print(f"Part 1: {part1(passwords)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(passwords)}")
  print (datetime.now()-starttime)