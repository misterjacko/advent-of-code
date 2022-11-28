from datetime import datetime
import itertools


def part1(data):
    return False

def part2(data):
    return False


def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(',') for line in lines]

def test_loadData():
    assert loadData("test_input.txt") == ["1","3","5","3"]


def t_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == "1" 


def t_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == True 


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
