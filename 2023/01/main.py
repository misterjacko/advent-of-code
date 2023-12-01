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
    total = 0
    for line in data:
        numbers = []
        for i in line:
            if i.isnumeric():
                numbers.append(i)
        total += (10*int(numbers[0])+int(numbers[-1]))
                
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 142


##########
# Part 2 #
##########

def part2(data):
    total = 0
    # i am not proud of this...
    number_dict = {
        "one": "o1e",
        "two": "t2o",
        "three": "t3e", 
        "four": "4",
        "five": "5e",
        "six": "6",
        "seven": "7n",
        "eight": "e8t",
        "nine": "n9e",
    }

    for line in data:
        numbers = []
        for k, v in number_dict.items():
            if k in line:
                line = line.replace(k, v)
        for i in line:
            if i.isnumeric():
                numbers.append(i)
        total += (10*int(numbers[0])+int(numbers[-1]))
    return total

def test_part2():
    testData = "test_input2.txt"
    data = loadData(testData)
    assert part2(data) == 281


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
