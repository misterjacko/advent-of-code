from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]

# def test_loadData():
#     assert loadData("test_input.txt") == []


# def process_data(data):
#     return data

# def test_process_data():
#     test_data = loadData("test_input.txt")
#     assert process_data(True) == True
#     assert process_data(test_data) == []


##########
# Part 1 #
##########

def part1(data):
    sum = 0
    max_dice = {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    for line in data:
        valid= True
        dice_dict = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        game_number = int(line.split(":")[0].split(" ")[1])
        dice_rolls = line.split(":")[1].split(";")
        for roll in dice_rolls:
            for dice in roll.split(","):
                color = dice.split()[1].strip()
                dice_count = int(dice.split()[0].strip())
                dice_dict[color] = max(dice_count, dice_dict[color])
        for k in max_dice.keys():
            if dice_dict[k] > max_dice[k]:
                valid = False
        if valid:
            sum += game_number
    return sum

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 8


##########
# Part 2 #
##########

def part2(data):
    sum = 0
    for line in data:
        power = 1
        dice_dict = {
            "red": 0,
            "green": 0,
            "blue": 0,
        }
        dice_rolls = line.split(":")[1].split(";")
        for roll in dice_rolls:
            for dice in roll.split(","):
                color = dice.split()[1].strip()
                dice_count = int(dice.split()[0].strip())
                dice_dict[color] = max(dice_count, dice_dict[color])
        for v in dice_dict.values():
            power *= v
        sum += power
    return sum

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 2286


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
