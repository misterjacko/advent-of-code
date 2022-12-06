from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines][0]

def test_loadData():
    assert loadData("test_input.txt") == "mjqjpqmgbljsphdztnvjfqwrcgsmlb"


##########
# Part 1 #
##########

def part1(data):
    for i in range(len(data)):
        if len(set(data[i:i+4])) == 4:
            return i+4
            
def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 7
    assert part1("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5
    assert part1("nppdvjthqldpwncqszvftbrmjlhg") == 6
    assert part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10


##########
# Part 2 #
##########

def part2(data):
    for i in range(len(data)):
        if len(set(data[i:i+14])) == 14:
            return i+14

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 19
    assert part2("bvwbjplbgvbhsrlpgdmjqwftvncz") == 23
    assert part2("nppdvjthqldpwncqszvftbrmjlhg") == 23
    assert part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 29


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
