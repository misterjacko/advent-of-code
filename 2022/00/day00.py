from datetime import datetime
import itertools


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]

def test_loadData():
    assert loadData("test_input.txt") == []


def process_data(data):
    return data

def test_process_data():
    test_data = loadData("test_input.txt")
    assert process_data(True) == True
    assert process_data(test_data) == []


##########
# Part 1 #
##########

def part1(data):
    return False

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == False


##########
# Part 2 #
##########

def part2(data):
    return False

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == False


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
