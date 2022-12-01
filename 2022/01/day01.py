from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]

def test_loadData():
    assert loadData("test_input.txt") == ["1000","2000","3000","","4000","","5000","6000","","7000","8000","9000","","10000"]


##########
# Part 1 #
##########

def part1(data):
    max_elf = 0
    cur = 0
    for cals in data:
        if cals == "":
            cur = 0
        else:
            cur += int(cals)
            max_elf = max(max_elf, cur)
    return max_elf
        

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 24000


##########
# Part 2 #
##########

def part2(data):
    l = [0] * 3
    cur = 0
    for cals in data:
        if cals == "":
            if cur > l[0]:
                l[0] = cur
                l.sort()
            cur = 0
        else:
            cur += int(cals)
    if cur > l[0]:
        l[0] = cur


    return sum(l)
        

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 45000


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
