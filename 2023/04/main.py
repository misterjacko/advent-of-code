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
        winning_numbers = {}
        for val in line.split(":")[1].split(" | ")[0].split():
            winning_numbers[val] = 0
        for val in line.split(":")[1].split(" | ")[1].split():
            if val in winning_numbers:
                winning_numbers[val] += 1
        total_winners = sum(winning_numbers.values())
        if total_winners > 0:
            multiple = 1
            for i in range(total_winners-1):
                multiple *= 2
            total += multiple
    return total

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 13


##########
# Part 2 #
##########

def part2(data):
    card_counts = {}
    for i in range(len(data)):
        card_counts[i+1] = 1
    for line in data:
        winning_numbers = {}
        card_number = int(line.split(":")[0].split(" ")[-1])
        card_multiple = card_counts[card_number]
        for val in line.split(":")[1].split(" | ")[0].split():
            winning_numbers[val] = 0
        for val in line.split(":")[1].split(" | ")[1].split():
            if val in winning_numbers:
                winning_numbers[val] += 1
        total_winners = sum(winning_numbers.values())
        for i in range(total_winners):
            card_counts[card_number+i+1] += card_multiple
    return sum(card_counts.values())

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 30


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
