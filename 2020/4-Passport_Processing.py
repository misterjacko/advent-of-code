from datetime import datetime
from schema import Schema, And, Use, Optional
import re

def combinelines(tempdata):
  oneline = {}
  for i in range(len(tempdata)):
    for j in range(len(tempdata[i])):
      tempval = tempdata[i][j].split(':')
      oneline[tempval[0]] = tempval[1]
      #oneline.append(tempdata[i][j])
  return oneline

def readFile(filename):
  dataset = []
  datasetindex = 0
  datafile = open(filename, "r")
  tempdata = []
  for aline in datafile:
    if aline != "\n":
      tempdata.append(re.split('[ |\n]', aline.strip()))
    else: 
      # add data from temp to dataset in order? should it be a dictionary?
      dataset.append(combinelines(tempdata))
      datasetindex += 1
      tempdata = []
  dataset.append(combinelines(tempdata))
  datafile.close()
  return dataset

def part1(A):
  validcount = 0
  val = re.compile('cid')
  for passport in A:
    if len(passport) == 8:
      validcount += 1
    else:
      if len(passport) == 7 and not list(filter(val.match, passport)):
        validcount += 1
  return (validcount)

def hgttest(string):
  booo = False
  if 'cm' in string:
    if  150 <= int(string.replace('cm', '')) <= 193:
      booo = True
  elif 'in' in string:
    if  59 <= int(string.replace('in', '')) <= 76:
      booo = True
  else: 
    booo = False
  return booo

def validatedata (dict):
  schema = Schema({
    'byr': And(Use(int), lambda n: 1920 <= n <= 2002),
    'iyr': And(Use(int), lambda n: 2010 <= n <= 2020),
    'eyr': And(Use(int), lambda n: 2020 <= n <= 2030),
    'hgt': And(Use(str), lambda n: hgttest(dict['hgt'])),
    'hcl': And(Use(str), lambda n: re.match('#([A-Fa-f0-9]{6})', n)),
    'ecl': And(Use(str), lambda n: re.match('amb|blu|brn|gry|grn|hzl|oth', n)), # 
    'pid': And(Use(str), lambda n: re.match('^\d{9}$', n)),
    Optional('cid'): str
  })
  valid = schema.is_valid(dict)
  return valid

def part2(A):
  validcount = 0
  val = re.compile('cid')
  for passport in A: 
    if len(passport) == 8:
      if validatedata(passport):
        validcount += 1
    else:
      if len(passport) == 7 and not list(filter(val.match, passport)):
        if validatedata(passport):
          validcount += 1
  return (validcount)

def test():
  test_input = readFile("test.txt")
  assert part1(test_input) == 10
  assert part2(test_input) == 6

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)