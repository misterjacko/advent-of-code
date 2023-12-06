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
    total = 1
    times = [int(time) for time in data[0].split(":")[1].split()]
    dist = [int(dist) for dist in data[1].split(":")[1].split()]
    for race in range(len(times)):
        win_count = 0
        for i in range(times[race]+1):
            speed = 1*i
            race_time = times[race]-i
            distance_needed = dist[race]
            if speed*race_time > distance_needed:
                win_count += 1
        total *= win_count
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 288


##########
# Part 2 #
##########

def part2(data):
    win_count = 0
    time = ""
    dist = ""
    for num in data[0].split(":")[1]:
        if num.isdigit():
            time = f"{time}{num}"
    for num in data[1].split(":")[1]:
        if num.isdigit():
            dist = f"{dist}{num}"
    time = int(time)
    dist = int(dist)
    for i in range(time):
        speed = 1*i
        race_time = time-i
        distance_needed = dist
        if speed*race_time > distance_needed:
            win_count += 1
    return win_count

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 71503


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
