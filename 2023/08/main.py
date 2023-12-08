from datetime import datetime
from collections import deque
from math import lcm


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

def parse_steps(steps):
    step_list = []
    for step in steps:
        if step == "L":
            step_list.append(0)
        else:
            step_list.append(1)
    return deque(step_list)


def part1(data):
    total_steps = 0
    position = "AAA"
    step_list = parse_steps(data[0])
    step_map = {}
    for i in range(2,len(data)):
        this_key = data[i].split(" = ")[0]
        this_left = data[i].split(" = ")[1][1:4]
        this_right = data[i].split(" = ")[1][6:9]
        step_map[this_key] = (this_left, this_right)
    while True:
        position = step_map[position][step_list[0]]
        total_steps += 1
        if position == "ZZZ":
            return total_steps
        step_list.rotate(-1)


def test_part1():
    testData = "test_input.txt"
    testData2 = "test_input2.txt"
    data = loadData(testData)
    data2 = loadData(testData2)
    assert part1(data) == 2
    assert part1(data2) == 6


##########
# Part 2 #
##########


def part2(data):
    least_common_multiple = 1
    positions = set()
    step_list = parse_steps(data[0])
    step_map = {}
    all_positions = set()

    # map data
    for i in range(2,len(data)):
        this_key = data[i].split(" = ")[0]
        # find starting points
        if this_key[2] == "A":
            positions.add(this_key)
            all_positions.add(this_key)
        this_left = data[i].split(" = ")[1][1:4]
        this_right = data[i].split(" = ")[1][6:9]
        step_map[this_key] = (this_left, this_right)
    
    # determine Z interval for each start and track LCM
    for position in positions:
        steps = 0
        zs = []
        while len(zs) < 2:
            steps += 1
            position = step_map[position][step_list[0]]
            if position[2] == "Z":
                zs.append(steps)
            step_list.rotate(-1)
        least_common_multiple = lcm(least_common_multiple,(zs[1]-zs[0]))
    return least_common_multiple


def test_part2():
    testData = "test_input3.txt"
    data = loadData(testData)
    assert part2(data) == 6


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
