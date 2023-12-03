from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]


def process_data(data):
    symbol_map = {}
    for i in range(len(data)):
        for j , char in enumerate(data[i]):
            if not char.isdigit() and char != ".":
                symbol_map[(i,j)] = char

    return symbol_map

def symbol_search(digit_start, digit_end, symbol_map):
    i = digit_start[0]
    # check above
    for j in range(digit_start[1]-1, digit_end[1]+2):
        if (i-1, j) in symbol_map:
            return True
    # check front and back
    if (i, digit_start[1]-1) in symbol_map or (i, digit_end[1]+1) in symbol_map:
        return True
    # check below
    for j in range(digit_start[1]-1, digit_end[1]+2):
        if (i+1, j) in symbol_map:
            return True
    return False

def gear_search(digit_start, digit_end, symbol_map):
    i = digit_start[0]
    # check above
    for j in range(digit_start[1]-1, digit_end[1]+2):
        if (i-1, j) in symbol_map and symbol_map[(i-1, j)] == "*":
            return (i-1, j)
    # check front and back
    if (i, digit_start[1]-1) in symbol_map and symbol_map[(i, digit_start[1]-1)] == "*":
        return (i, digit_start[1]-1)
    if (i, digit_end[1]+1) in symbol_map and symbol_map[(i, digit_end[1]+1)] == "*":
        return (i,digit_end[1]+1)
    # check below
    for j in range(digit_start[1]-1, digit_end[1]+2):
        if (i+1, j) in symbol_map and symbol_map[(i+1, j)] == "*":
            return (i+1, j)
    return False

##########
# Part 1 #
##########

def part1(data):
    not_parts = 0
    symbol_map = process_data(data)
    for i in range(len(data)):
        j = 0
        while j < len(data[i]):
            if data[i][j].isnumeric():
                digits = data[i][j]
                digit_start = (i,j)
                while j+1 < len(data[i]) and data[i][j+1].isnumeric():
                    j+=1
                    digits = f"{digits}{data[i][j]}"
                digit_end = (i,j)
                if symbol_search(digit_start, digit_end, symbol_map):
                    not_parts += int(digits)
            j+= 1
    return not_parts

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 4361


##########
# Part 2 #
##########

def part2(data):
    gears = 0
    gear_map = {}
    symbol_map = process_data(data)
    for i in range(len(data)):
        j = 0
        while j < len(data[i]):
            if data[i][j].isnumeric():
                digits = data[i][j]
                digit_start = (i,j)
                while j+1 < len(data[i]) and data[i][j+1].isnumeric():
                    j+=1
                    digits = f"{digits}{data[i][j]}"
                digit_end = (i,j)
                gear_id = gear_search(digit_start, digit_end, symbol_map)
                if gear_id:
                    if gear_id in gear_map:
                        gear_map[gear_id].append(int(digits))
                    else:
                        gear_map[gear_id] = [int(digits)]
            j+= 1
    for v in gear_map.values():
        if len(v) == 2:
            gears += (v[0]*v[1])
    return gears

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 467835


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
