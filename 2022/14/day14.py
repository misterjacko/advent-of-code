from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(" -> ") for line in lines]

def process_data(data):
    for i in range(len(data)):
        for j in range(len(data[i])):
            data[i][j] = (int(data[i][j].split(",")[0]),int(data[i][j].split(",")[1]))
    return data

##########
# Part 1 #
##########

def process_rock(data):
    rock = set()
    limit = 0
    for line in data:
        for i in range(1, len(line)):
            if line[i-1][0] == line[i][0]:
                for j in range(min(line[i-1][1], line[i][1]), max(line[i-1][1], line[i][1])+1):
                    rock.add((line[i-1][0], j))
                    limit = max(limit, j)
            else:
                limit = max(limit, line[i][1])
                for j in range(min(line[i-1][0], line[i][0]), max(line[i-1][0], line[i][0])+1):
                    rock.add((j, line[i-1][1]))
    return rock, limit


def drop_sand(rock, limit):
    grain = [500,0]
    while grain[1] < limit:
        if (grain[0], grain[1]+1) not in rock:
            grain[1] += 1
        elif (grain[0]-1, grain[1]+1) not in rock:
            grain[0] -= 1
            grain[1] += 1
        elif (grain[0]+1, grain[1]+1) not in rock:
            grain[0] += 1
            grain[1] += 1
        else:
            return (grain[0], grain[1])
    return False # fell out bottom


def part1(data):
    rock, limit = process_rock(data)
    sand = set()
    while True:
        new_sand = drop_sand(rock, limit)
        if new_sand:
            sand.add(new_sand)
            rock.add(new_sand)
        else:
            return len(sand)

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part1(data) == 24


##########
# Part 2 #
##########


def drop_sand2(rock, limit):
    if (500,0) in rock:
        return False
    grain = [500,0]
    while grain[1] < limit:
        if (grain[0], grain[1]+1) not in rock:
            grain[1] += 1
        elif (grain[0]-1, grain[1]+1) not in rock:
            grain[0] -= 1
            grain[1] += 1
        elif (grain[0]+1, grain[1]+1) not in rock:
            grain[0] += 1
            grain[1] += 1
        else:
            return (grain[0], grain[1])
    return ((grain[0], grain[1])) # settles at lowest level in open spot


def part2(data):
    rock, limit = process_rock(data)
    limit += 1 # this is the lowest point sand can BE
    sand = set()
    while True:
        new_sand = drop_sand2(rock, limit)
        if new_sand:
            sand.add(new_sand)
            rock.add(new_sand)
        else:
            return len(sand)

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part2(data) == 93


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
