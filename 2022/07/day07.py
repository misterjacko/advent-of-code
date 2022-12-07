from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip().split(" ") for line in lines]

def process_data(data):
    result_dict = {}
    folder_stack = []
    
    i = 0
    while i < len(data):
        if data[i][0] == "$": #its a command. 
            if data[i][1] == "cd":
                if data[i][2] == "..": # backing out of the dir means adding its children to its size
                    folder_name = tuple(folder_stack) # generates "fully qualified" folder name that is hashable
                    folder_stack.pop()
                    for child in result_dict[folder_name]["children"]:
                        result_dict[folder_name]["size"] += result_dict[child]["size"]
                else:
                    folder_stack.append(data[i][2])
                    
            elif data[i][1] == "ls": # next lines are going to be contents
                folder_name = tuple(folder_stack)
                children = []
                size = 0
                while i + 1 < len(data) and data[i+1][0] != "$": # reads contents til next $ command
                    i += 1
                    if data[i][0] == "dir": # adds child fq folder to list of children
                        folder_stack.append(data[i][1])
                        children.append(tuple(folder_stack))
                        folder_stack.pop()
                    else:
                        size += int(data[i][0])
                        
                result_dict[folder_name] = {"size": size, "children": children}
        i += 1

    while len(folder_stack) > 0:
        folder_name = tuple(folder_stack)
        folder_stack.pop()
        for child in result_dict[folder_name]["children"]:
            result_dict[folder_name]["size"] += result_dict[child]["size"]    
    return result_dict


##########
# Part 1 #
##########

def part1(data):

    total_size = 0
    for k in data:
        if data[k]["size"] <= 100000:
            total_size += data[k]["size"]
    return total_size

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part1(data) == 95437


##########
# Part 2 #
##########

def part2(data):  

    free_disk = 70000000 - data[tuple(["/"])]["size"]
    need_to_free = 30000000 - free_disk
    best = 70000000 # known larger than answers

    for k in data:
        if data[k]["size"] >= need_to_free:
            best = min(best, data[k]["size"])
    return best

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    data = process_data(data)
    assert part2(data) == 24933642


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    data = process_data(data)
    p1 = part1(data)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part2(data)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
