from datetime import datetime
from itertools import combinations


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

def parse_dataset(data, expanded_amount):
    rows_with_stars = set()
    cols_with_stars = set()
    star_locations = []
    for i in range(len(data)):
        for j in range(len(data[i])):
            if data[i][j] == "#":
                rows_with_stars.add(i)
                cols_with_stars.add(j)
                star_locations.append((i,j))
    # find row offsets
    row_offsets = calculate_offsets(rows_with_stars, len(data), expanded_amount)
    col_offsets = calculate_offsets(cols_with_stars, len(data[i]), expanded_amount)

    # transform stars with offset values
    for i in range(len(star_locations)):
        star_locations[i] = (star_locations[i][0]+row_offsets[star_locations[i][0]],star_locations[i][1]+col_offsets[star_locations[i][1]])
    return star_locations


def calculate_offsets(star_set, range_max, expanded_amount):
    offset = 0
    offset_map = {}
    for i in range(range_max):
        if i not in star_set:
            offset += expanded_amount-1
        offset_map[i] = offset
    return offset_map


def part1(data, expanded_amount):
    total_distance = 0
    star_locations  = parse_dataset(data, expanded_amount)
    for stars in combinations(star_locations, 2):
        total_distance += (abs(stars[0][0]-stars[1][0]) + abs(stars[0][1]-stars[1][1]))
    return total_distance


def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data, 2) == 374


##########
# Part 2 #
##########

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data, 10) == 1030


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    p1 = part1(data, 2)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part1(data, 1000000)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
