from datetime import datetime
from collections import deque
import math


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]

# def test_loadData():
#     assert loadData("test_input.txt") == []


def process_data(data):
    dict = {}
    for i in range(0, len(data), 7):
        monkey = int(data[i].split(" ")[1].strip(":"))
        items = deque()
        for item in data[i+1].split(":")[1].split(", "):
            items.append(int(item))
        operation_line = data[i+2].split(" ")
        operation_sign = operation_line[-2]
        if operation_line[-1] == "old":
            operation_value = "old"
        else:
            operation_value = int(operation_line[-1])
        operation = (operation_sign, operation_value)
        test = int(data[i+3].split(" ")[-1])
        true = int(data[i+4].split(" ")[-1])
        false = int(data[i+5].split(" ")[-1])
        dict[monkey] = {
            "items":items,
            "operation":operation,
            "test":test,
            "true":true,
            "false":false,
            "actions":0,
        }
    # print(dict)
    return dict

# def test_process_data():
#     test_data = loadData("test_input.txt")
#     # assert process_data(True) == True
#     assert process_data(test_data) == []


##########
# Part 1 #
##########

def inspect_item(item, operation):
    op_value = operation[1]
    if op_value == "old":
        op_value = item
    if operation[0] == "+":
        return (item+op_value)
    else:
        return (item*op_value)

def test_inspect_item():
    assert inspect_item(79, ("*", 19)) == 1501
    assert inspect_item(98, ("*", 19)) == 1862
    assert inspect_item(54, ("+", 6)) == 60
    assert inspect_item(65, ("+", 6)) == 71
    assert inspect_item(79, ("*", "old")) == 6241
    assert inspect_item(60, ("*", "old")) == 3600


def part1(data):
    data = process_data(data)
    for _ in range(20):
        for monkey in range(len(data)):
            while len(data[monkey]["items"])> 0:
                item = data[monkey]["items"].popleft()
                data[monkey]["actions"] += 1
                item = inspect_item(item, data[monkey]["operation"])//3
                if item % data[monkey]["test"] == 0:
                    data[data[monkey]["true"]]["items"].append(item)
                else:
                    data[data[monkey]["false"]]["items"].append(item)
    
    actions = []
    for monkey in data.keys():
        actions.append(data[monkey]["actions"])
    actions.sort()
    return actions[-1]*actions[-2]

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 10605


##########
# Part 2 #
##########

def part2(data):
    data = process_data(data)
    lcm = 1
    for monkey in data.keys():
        lcm *= data[monkey]["test"]
  
    for _ in range(10000):
        for monkey in range(len(data)):
            while len(data[monkey]["items"])> 0:
                item = data[monkey]["items"].popleft()
                data[monkey]["actions"] += 1
                item = inspect_item(item, data[monkey]["operation"])
                item %= lcm # essentially subtracts out n*lcm of every item. 
                if item % data[monkey]["test"] == 0:
                    data[data[monkey]["true"]]["items"].append(item)
                else:
                    data[data[monkey]["false"]]["items"].append(item)
    
    actions = []
    for monkey in data.keys():
        actions.append(data[monkey]["actions"])
    actions.sort()
    return actions[-1]*actions[-2]

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 2713310158


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    p1 = part1(data)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")
    data = loadData(dataFile)

    p2 = part2(data)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
