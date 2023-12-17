from datetime import datetime
from collections import deque, OrderedDict


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines[0].split(',')]



##########
# Part 1 #
##########

def decode_box(s: str):
    value = 0
    for char in s:
        value += ord(char)
        value *= 17
        value = value %256
    return value


def part1(data):
    total = 0
    for string in data:
        total += decode_box(string)
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 1320


##########
# Part 2 #
##########

def find_operator_index(s: str):
    i = 0
    while s[i].isalpha():
        i += 1
    return i

def part2(data):
    focusing_power = 0
    boxes = {}
    for i in range(256):
        boxes[i] = OrderedDict()
    for string in data:
        operator_index = find_operator_index(string)
        label = string[:operator_index]
        box = decode_box(label)
        if string[operator_index] == "-":
            if label in boxes[box].keys():
                del boxes[box][label]
        else: # "="
            focal_length = int(string[operator_index+1])
            boxes[box][label] = focal_length
            
    for i in range(256):
        box_value = i+1
        for slot, lenses in enumerate(boxes[i].items()):
            focusing_power += box_value * (slot+1) * lenses[1]

    return focusing_power

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 145


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
