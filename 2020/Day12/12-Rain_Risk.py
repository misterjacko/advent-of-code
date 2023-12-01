
from datetime import datetime
import math

def readFile(filename):
  with open(filename, "r") as fh:
    lines = fh.readlines()
  return [x.strip() for x in lines]

def parseData(freshlist):
  newlist = []
  for line in freshlist:
    newlist.append([line[:1], line[1:]])
  return newlist

def changeHeading(heading, turnDirection, turnAngle):
  directions = ['E', 'S', 'W', 'N']*3
  angle = int(turnAngle)/90
  for i in enumerate(directions):
    if i[1] == heading:
      if turnDirection == 'R':
        return directions[i[0] + int(angle)]
      elif turnDirection == 'L':
        return directions[i[0] + 4 - int(angle)]

def movePosition(position, heading, distance):
  if heading == 'E':
    position[0] = position[0] + distance
  elif heading == 'W':
    position[0] = position[0] - distance
  elif heading == 'N':
    position[1] = position[1] + distance
  elif heading == 'S':
    position[1] = position[1] - distance
  return position

def part1(A):
  formattedList = parseData(A)
  position = [0, 0]
  heading = "E"
  for instruction in formattedList:
    if instruction[0] == 'R' or instruction[0] == 'L':
      heading = changeHeading(heading, instruction[0], instruction[1])
    elif instruction[0] == 'F':
      position = movePosition(position, heading, int(instruction[1]))
    else:
      position = movePosition(position, instruction[0], int(instruction[1]))
  return abs(position[0]) + abs(position[1])

def movePosition2(position, waypointPosition, distance):
  position[0] = position[0] + (waypointPosition[0] * distance)
  position[1] = position[1] + (waypointPosition[1] * distance)
  return position

def rotateWaypoint(waypointPosition, turnDirection, turnAngle):
  turnAngle /= 90
  if turnAngle % 2 == 0:
    return [int(waypointPosition[0])*(turnAngle/-2), int(waypointPosition[1])*(turnAngle/-2)]
  else:
    waypointPosition[0], waypointPosition[1] = waypointPosition[1], waypointPosition[0]
    while turnAngle > 4:
      turnAngle -= 4
    else:
      turnCommand = ''.join([turnDirection, str(math.trunc(turnAngle))])
      if turnCommand == 'R1' or turnCommand == 'L3':
        return [waypointPosition[0], int(waypointPosition[1])*-1]
      elif turnCommand == 'R3' or turnCommand == 'L1':
        return [int(waypointPosition[0])*-1, waypointPosition[1]]

def part2(A):
  formattedList = parseData(A)
  position = [0, 0]
  waypointPosition = [10, 1]
  for instruction in formattedList:
    if instruction[0] == 'R' or instruction[0] == 'L':
      waypointPosition = rotateWaypoint(waypointPosition, instruction[0], int(instruction[1]))
    elif instruction[0] == 'F':
      position = movePosition2(position, waypointPosition, int(instruction[1]))
    else:
      waypointPosition = movePosition(waypointPosition, instruction[0], int(instruction[1]))
  return abs(position[0]) + abs(position[1])

def test():
  test_input = (readFile("test.txt"))
  assert part1(test_input) == 25
  assert part2(test_input) == 286

if __name__ == "__main__":
  test()
  starttime = datetime.now()
  vals = readFile("data.txt")
  print(f"Part 1: {part1(vals)}")
  print (datetime.now()-starttime)
  print(f"Part 2: {part2(vals)}")
  print (datetime.now()-starttime)