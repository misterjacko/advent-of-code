from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(",") for line in lines]




##########
# Part 1 #
##########

def part1(data):
    count = 0
    for line in data:
        r1 = line[0].split("-")
        r2 = line[1].split("-")
        s1 = set()
        s2 = set()

        for i in range(int(r1[0]), int(r1[1])+1):
            s1.add(i)
        for i in range(int(r2[0]), int(r2[1])+1):
            s2.add(i)
        if len(s1.intersection(s2)) == len(s1) or len(s1.intersection(s2)) == len(s2):
            count += 1
    return count

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 2


##########
# Part 2 #
##########

def part2(data):
    count = 0
    for line in data:
        r1 = line[0].split("-")
        r2 = line[1].split("-")
        s1 = set()
        s2 = set()

        for i in range(int(r1[0]), int(r1[1])+1):
            s1.add(i)
        for i in range(int(r2[0]), int(r2[1])+1):
            s2.add(i)
        if len(s1.intersection(s2)) > 0:
            count += 1
    return count

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 4


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
