from datetime import datetime
from itertools import chain


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(":") for line in lines]


def process_data(data):
    for i in range(len(data)):
        for j in range(len(data[0])):
            line = data[i][j]
            x = int(line.split("=")[1].strip(", y"))
            y = int(line.split("=")[2])
            data[i][j] = (x,y)
            
    return data

##########
# Part 1 #
##########


def clear_to_beacon(sensor, distance, q_row):
    if sensor[1] + distance >= q_row or sensor[1] - distance <= q_row:
        row_offset = abs(q_row - sensor[1])
        start = sensor[0] - (distance - row_offset)
        end = sensor[0] + (distance - row_offset)
        return (start, end)
    return None


def part1(data, q_row):
    beacons_sensors = set()
    searched = []
    for line in data:
        beacons_sensors.add(line[0])
        beacons_sensors.add(line[1])

        distance = abs(line[0][0] - line[1][0]) + abs(line[0][1] - line[1][1])
        searched.append(clear_to_beacon(line[0], distance, q_row))
    searched = combine_ranges(searched)
    total = 0
    for rng in searched:
        total += ((rng[1]) - rng[0])
        for bs in beacons_sensors:
            if bs[1] == q_row and bs[0] >= rng[0] and bs[0] <= rng[0]:
                total -= 1
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part1(data, 10) == 26


##########
# Part 2 #
##########

def add_range(grid, sensor, distance, mx):
    iter_start = max(sensor[1], 0)
    iter_end = min((sensor[1] + (distance+1)), mx)
    d = distance
    for i in range(iter_start, iter_end):
        range_start = max(sensor[0]-(d), 0)
        range_end = min(sensor[0]+(d), mx)
        d -= 1
        grid[i].append((range_start, range_end))
    d = distance
    iter_end = max((sensor[1] - (distance+1)), -1)
    for i in range(iter_start, iter_end, -1):
        range_start = max(sensor[0]-(d), 0)
        range_end = min(sensor[0]+(d), mx)
        d -= 1
        grid[i].append((range_start, range_end))
    return grid


def combine_ranges(ranges):
        ranges.sort(key=lambda x: x[0])
        c_ranges = [ranges[0]]
        for i in range(1, len(ranges)):
            if ranges[i][0] <= c_ranges[-1][1]:
                c_ranges[-1] = (min(c_ranges[-1][0], ranges[i][0]), max(c_ranges[-1][1], ranges[i][1]))
            else:
                c_ranges.append(ranges[i])
        return c_ranges


def part2(data, mx):
    grid = [[] for _ in range(mx+1)]
    for line in data:
        print(line) # to show progress
        distance = abs(line[0][0] - line[1][0]) + abs(line[0][1] - line[1][1])
        grid = add_range(grid, line[0], distance, mx)

    for i in range(len(grid)):
        if i%100_000 == 0:
            print(i) # to show progress
        combined = combine_ranges(grid[i])
        if combined[0] != (0,mx):
            print(i)
            print(combined)
            return i + (4_000_000 * (combined[0][-1]+1))
    return True

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part2(data, 20) == 56_000_011


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    data = process_data(data)

    p1 = part1(data, 2_000_000)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part2(data, 4_000_000)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
