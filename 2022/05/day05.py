from datetime import datetime
from collections import deque


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line for line in lines]


def process_data(data):
    stacks = [deque() for _ in range(len(data[0])//4)]
    ln = 0
    for l, line in enumerate(data):
        if line[1] == "1":
            ln = l + 2 # allow us to resume at instructions
            break
        for i in range(0, len(line), 4):
            if line[i+1] != " ":
                stacks[i//4].appendleft(line[i+1])
    return stacks, data[ln:] # returns list of queues and instruction section


##########
# Part 1 #
##########

def part1(data):
    data, instructions = process_data(data)
    for inst in instructions:
        inst = inst.split(" ")
        for _ in range(int(inst[1])):
            val = data[int(inst[3])-1].pop()
            data[int(inst[5])-1].append(val)

    ans = ""
    for stack in data:
        if len(stack) >= 1:
            ans += stack[-1]
    return ans

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == "CMZ"


##########
# Part 2 #
##########

def part2(data):
    data, instructions = process_data(data)
    for inst in instructions:
        inst = inst.split(" ")
        tempq = deque()
        for _ in range(int(inst[1])):
            tempq.appendleft(data[int(inst[3])-1].pop())
        data[int(inst[5])-1].extend(tempq)

    ans = ""
    for stack in data:
        if len(stack) >= 1:
            ans += stack[-1]
    return ans

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == "MCD"


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
