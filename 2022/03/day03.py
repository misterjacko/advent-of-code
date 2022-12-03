from datetime import datetime


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

def part1(data):
    total = 0
    for bag in data:
        pocket1 = set([*bag[:len(bag)//2]])
        pocket2 = set([*bag[len(bag)//2:]])
        intersect = list(pocket1.intersection(pocket2))[0]
        if ord(intersect) >= ord("a"):
            total += ord(intersect) - ord("a") + 1
        else:
            total += ord(intersect) - ord("A") + 27
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 157


##########
# Part 2 #
##########

def part2(data):
    total = 0
    for i in range(0, len(data), 3):
        intersect = set(data[i]).intersection(set(data[i+1]), set(data[i+2]))
        intersect =list(intersect)[0]
        if ord(intersect) >= ord("a"):
            total += ord(intersect) - ord("a") + 1
        else:
            total += ord(intersect) - ord("A") + 27
    return total

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 70


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
