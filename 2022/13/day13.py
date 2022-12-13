from datetime import datetime
import json


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]


def process_data(data):
    for i in range(len(data)):
        if data[i] != "":
            data[i] = json.loads(data[i])
    return data


##########
# Part 1 #
##########

def compare(left, right):
    i = 0
    while i < len(left) and i < len(right):
        left_val = left[i]
        right_val = right[i]
        if type(left_val) != type(right_val):
            if not isinstance(left_val, (list)):
                left_val = [left_val]
            else:
                right_val = [right_val]
        # should now both be lists or both 
        if isinstance(left_val, (int)):
            if left_val > right_val:
                return False
            elif left_val < right_val:
                return True
            else:
                pass
        else:
            recurse = compare(left_val, right_val)
            if recurse == None:
                pass
            else:
                return recurse

        i += 1
    if len(left) < len(right):
        return True
    elif len(left) > len(right):
        return False
    return None


def part1(data):
    count = 0
    data = process_data(data)
    for i in range(0, len(data), 3):
        left = data[i]
        right = data[i+1]
        if compare(left, right):
            count += ((i//3)+1)
    return count

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 13


##########
# Part 2 #
##########

def mergesort(data):
    if len(data) > 1:
        left = data[:len(data)//2]
        right = data[len(data)//2:]
        left = mergesort(left)
        right = mergesort(right)
    
        d_i, l_i, r_i = 0, 0, 0
        while l_i < len(left) and r_i < len(right):
            if compare(left[l_i], right[r_i]):
                data[d_i] = left[l_i]
                l_i += 1
            else:
                data[d_i] = right[r_i]
                r_i += 1
            d_i += 1
        while l_i < len(left):
            data[d_i] = left[l_i]
            l_i += 1
            d_i += 1
        while r_i < len(right):
            data[d_i] = right[r_i]
            r_i += 1
            d_i += 1
    return data


def part2(data):
    new_list = ["[[2]]", "[[6]]"]
    for line in data:
        if line != "":
            new_list.append(line)
    data = mergesort(process_data(new_list))

    result = 1
    for i in range(len(data)):
        if data[i] == [[2]]:
            result *= (i+1)
        elif data[i] == [[6]]:
            return result * (i+1)

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 140


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
