from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [list(line.strip()) for line in lines]


##########
# Part 1 #
##########
def find_start(data):
    for i in range(len(data)):
        for j in range(len(data[i])):
            if data[i][j] == "S":
                return (i,j)

def find_posible_connection(start_position, connections, data):
    directions = [(-1,0), (0, 1), (0, -1), (1,0)]
    connecting = []
    for direction in directions:
        new_position = (start_position[0]+direction[0],start_position[1]+direction[1])
        # in bounds
        if new_position[0] >=0 and new_position[0] < len(data):
            if new_position[1] >= 0 and new_position[1] < len(data[1]):
                pipe = data[new_position[0]][new_position[1]]
                if pipe == ".":
                    continue
                for connection in connections[pipe]:
                    if (new_position[0]+connection[0],new_position[1]+connection[1]) == start_position:
                        connecting.append(new_position)
    return connecting      


def part1(data):
    connections = {
        "|": [(-1, 0), (1, 0)],
        "-": [(0, -1), (0, 1)],
        "L": [(-1, 0), (0, 1)],
        "J": [(-1, 0), (0, -1)],
        "7": [(0, -1), (1, 0)],
        "F": [(0, 1), (1, 0)],
    }
    start_position = find_start(data)
    print(start_position)
    # find first connection and just pick the first one
    position, last_position = find_posible_connection(start_position, connections, data)[0], start_position
    steps = 1
    while position != start_position:
        pipe = data[position[0]][position[1]]
        conns = connections[pipe]
        for next_direction in conns:
            next_position = (position[0]+next_direction[0],position[1]+next_direction[1])
            if next_position == last_position:
                continue
            else:
                position, last_position = next_position, position
                steps += 1
                break

    return steps//2

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 8


##########
# Part 2 #
##########

# get the proper pipe to replace the S with
def rename_start_position(start_position, connecting):
    row1 = connecting[0][0]-start_position[0]
    col1 = connecting[0][1]-start_position[1]
    row2 = connecting[1][0]-start_position[0]
    col2 = connecting[1][1]-start_position[1]
    if col1 == 0 and col2 == 0:
        return "|"
    elif row1 == 0 and row2 == 0:
        return "-"
    
    elif row1 == -1 or row2 == -1:
        if col1 == 1 or col2 == 1:
            return "L"
        else:
            return "J"
    else:
        if col1 == 1 or col2 == 1:
            return "F"
        else:
            return "7"
    


def part2(data):
    connections = {
        "|": [(-1, 0), (1, 0)],
        "-": [(0, -1), (0, 1)],
        "L": [(-1, 0), (0, 1)],
        "J": [(-1, 0), (0, -1)],
        "7": [(0, -1), (1, 0)],
        "F": [(0, 1), (1, 0)],
    }
    # collect positions of path
    visited = set()
    start_position = find_start(data)
    visited.add(start_position)
    possible_first_steps = find_posible_connection(start_position, connections, data)
    
    # rename S to actual pipe for inside/outside logic later
    data[start_position[0]][start_position[1]] = rename_start_position(start_position, possible_first_steps)
    
    # follow path
    position, last_position = possible_first_steps[0], start_position
    steps = 1
    visited.add(position)
    while position != start_position:
        pipe = data[position[0]][position[1]]
        conns = connections[pipe]
        for next_direction in conns:
            next_position = (position[0]+next_direction[0],position[1]+next_direction[1])
            if next_position == last_position:
                continue
            else:
                position, last_position = next_position, position
                visited.add(position)
                steps += 1
                break
    interior_count = 0
    
    # this will traverse each row and use the shape of the pipes in the visited set
    # to determine whether the next potential space will be inside or outside the loop
    # The idea is that we generally start outside and when we pass a | we switch between 
    # inside and outside. we ignore -. With turns, we need to do extra work to figure out
    # whether we treat them like a | (eg F--J or L--7 because they are essentially just 
    # pass through taking longer) or a - (eg L--J or F--7 because we are just hitting an 
    # edge and never going inside or outside)
    for i in range(len(data)):
        inside = False
        turn_start = None
        for j in range(len(data[i])):
            if (i,j) in visited:
                if data[i][j] == "L":
                    turn_start = "L"
                    inside = not inside
                elif data[i][j] == "F":
                    turn_start = "F"
                    inside = not inside
                elif data[i][j] == "-":
                    continue
                elif data[i][j] == "7" and turn_start == "L":
                    turn_start = None
                elif data[i][j] == "J" and turn_start == "F":
                    turn_start = None
                else:
                    inside = not inside
            elif inside:
                interior_count += 1
    return interior_count

def test_part2():
    testData = "test_input2.txt"
    data = loadData(testData)
    assert part2(data) == 8


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
