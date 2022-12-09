from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [list(line.strip()) for line in lines]


def process_data(data):
    for i in range(len(data)):
        for j in range(len(data[0])):
            data[i][j] = int(data[i][j])
    return data

def test_process_data():
    assert process_data([["3","2","1"]]) == [[3,2,1]]


##########
# Part 1 #
##########

def part1(data):
    seen = set()
    # from the sides
    for i in range(len(data)):
        # L to R
        tallest = -1
        for j in range(len(data[0])):
            if data[i][j] > tallest:
                seen.add((i,j))
                tallest = data[i][j]
        # R to L
        tallest = -1
        for j in range(len(data[0])-1,-1,-1):
            if data[i][j] > tallest:
                seen.add((i,j))
                tallest = data[i][j]
    # from the top/bottom
    for j in range(len(data[0])):
        # top to bottom
        tallest = -1
        for i in range(len(data)):
            if data[i][j] > tallest:
                seen.add((i,j))
                tallest = data[i][j]
        # bottom to top
        tallest = -1
        for i in range(len(data)-1,-1,-1):
            if data[i][j] > tallest:
                seen.add((i,j))
                tallest = data[i][j]
    return len(seen)

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part1(data) == 21


##########
# Part 2 #
##########

def in_bounds(data, i, j):
    if i < 0 or i >= len(data) or j < 0 or j >= len(data[0]):
        return False
    return True

def test_in_bounds():
    testData = "test_input.txt"
    data = loadData(testData)
    assert in_bounds(data, -1, 0) == False
    assert in_bounds(data, 5, 0) == False
    assert in_bounds(data, 0, -1) == False
    assert in_bounds(data, 0, 5) == False
    assert in_bounds(data, 4, 4) == True


def scan_trees(data, position, direction):
    distance = 0
    height = data[position[0]][position[1]]
    position = (position[0]+direction[0], position[1]+direction[1])
    while in_bounds(data, position[0], position[1]):
        distance += 1
        if data[position[0]][position[1]] >= height:
            break
        position = (position[0]+direction[0], position[1]+direction[1])
    return distance


def part2(data):
    directions = ((-1, 0),(1, 0),(0, -1),(0, 1))
    best = 0
    for row in range(1, len(data)-1):
        for col in range(1, len(data[0])-1):
            position_score = 1
            for direction in directions:
                position_score *= scan_trees(data, (row, col), direction)
            best = max(best, position_score)
    return best

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part2(data) == 8


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    data = process_data(data)
    p1 = part1(data)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part2(data)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
