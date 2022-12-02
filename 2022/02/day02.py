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
    score = 0 

    # A | X Rock
    # B | Y Paper
    # C | Z Scissors
    points = {
        "A X": 4, 
        "A Y": 8, 
        "A Z": 3,
        "B X": 1, 
        "B Y": 5, 
        "B Z": 9,
        "C X": 7, 
        "C Y": 2, 
        "C Z": 6,  
    }
    for round in data:
        score += points[round]
    return score

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 15


##########
# Part 2 #
##########

def part2(data):
    score = 0

    # X | A Rock
    # Y | B Paper
    # Z | C Scissors
    points = {
        "A X": 3, 
        "A Y": 4, 
        "A Z": 8,
        "B X": 1, 
        "B Y": 5, 
        "B Z": 9,
        "C X": 2, 
        "C Y": 6, 
        "C Z": 7,  
    }
    for round in data:
        score += (points[round])
    return score

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 12


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
