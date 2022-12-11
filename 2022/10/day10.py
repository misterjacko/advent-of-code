from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(" ") for line in lines]

##########
# Part 1 #
##########

def part1(data):
    total = 0
    check_cycles = [220, 180, 140, 100, 60, 20]
    register = 1
    cycle = 1
    i = 0
    while len(check_cycles) > 0:
        if data[i][0] == "noop":
            cycle += 1
        else:
            if cycle + 1 == check_cycles[-1]:
                total += (register * check_cycles[-1])
                check_cycles.pop()
            cycle += 2
            register += int(data[i][1])
        i += 1
        if len(check_cycles) > 0 and cycle >= check_cycles[-1]:
            total += (register * check_cycles[-1])
            check_cycles.pop()
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 13140


##########
# Part 2 #
##########

def print_crt(data):
    for start in range(0, 201, 40):
        print("".join(data[start:start+40]))


def set_i(i, result):
    if i + 1 == 40: # print and reset line
        print("".join(result))
        result = [" "] * 40
        return 0, result
    return i + 1, result


def part2(data):
    result = [" "] * 40 # looks cleaner for output
    sprite_pos = {0,1,2}
    i = 0
    for line in data:
        if i in sprite_pos:
            result[i] = "#"
        if line[0] == "noop":
            i, result = set_i(i, result)
        else:
            i, result = set_i(i, result)
            if i in sprite_pos:
                result[i] = "#"
            i, result = set_i(i, result)
            sprite_pos = {pos + int(line[1]) for pos in sprite_pos}
    return True

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == False # failing test to print output


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
