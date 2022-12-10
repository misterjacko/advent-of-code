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


def part2(data):
    result = [" "] * 240 # looks cleaner for output
    sprite_pos = set()
    for i in range(0, 201, 40):
        sprite_pos |= {i, i+1, i+2}
    i = 0
    for line in data:
        if i in sprite_pos:
            result[i] = "#"
        if line[0] == "noop":
            i += 1
        else:
            if i+1 in sprite_pos:
                result[i+1] = "#"
            i += 2
            sprite_pos = {pos + int(line[1]) for pos in sprite_pos}
    print_crt(result)
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
