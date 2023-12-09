from datetime import datetime
from collections import deque


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]


##########
# Part 1 #
##########

def get_next_line(line):
    if len(set(line)) == 1:
        return line
    else:
        next_line = []
        for i in range(len(line)-1):
            next_line.append(line[i+1]-line[i])
        line.append(get_next_line(next_line)[-1]+line[-1])
    return line


def part1(data):
    total = 0
    for line in data:
        total += get_next_line([int(num) for num in line.split()])[-1]
    return total


def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 114


##########
# Part 2 #
##########

def get_next_line2(line):
    if len(set(line)) == 1:
        return line
    else:
        next_line = deque([])
        for i in range(len(line)-1):
            next_line.append(line[i+1]-line[i])
        line.appendleft(line[0]-get_next_line2(next_line)[0])
    return line


def part2(data):
    total = 0
    for line in data:
        total += get_next_line2(deque([int(num) for num in line.split()]))[0]
    return total


def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 2


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    p1 = part1(data)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part2(data)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
