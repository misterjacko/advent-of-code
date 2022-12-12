from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [list(line.strip()) for line in lines]


def process_data(data): # returns 2d list with height values, and start and end coordinates
    start, end = None, None
    p2_starts = []
    for i in range(len(data)):
        for j in range(len(data[i])):
            if data[i][j] in ["S", "E"]:
                if data[i][j] == "S":
                    start = (i,j)
                    p2_starts.append(start)
                    data[i][j] = ord("a")-97
                else:
                    end = (i,j)
                    data[i][j] = ord("z")-97
            else:
                num = ord(data[i][j])-97
                data[i][j] = num
                if num == 0 and (i == 0 or j == 0 or i == len(data)-1 or j == len(data[i])-1):
                    p2_starts.append((i,j))
    return data, start, end, p2_starts


##########
# Part 1 #
##########

def is_in_bounds(data, position):
    if position[0] < 0 or position[1] < 0 or position[0] >= len(data) or position[1] >= len(data[0]):
        return False
    return True


def part1(data):
    data, start, end, _ = process_data(data)
    visited = {start:0}
    directions = ((-1, 0), (1, 0), (0, -1), (0, 1))

    def bfs(neighbors, steps):
        new_neighbors = []
        for neighbor in neighbors:
            for direction in directions:
                n_val = data[neighbor[0]][neighbor[1]]
                new_neighbor = (neighbor[0]+direction[0], neighbor[1]+direction[1])
                if new_neighbor not in visited.keys() and is_in_bounds(data, new_neighbor):
                    nn_val = data[new_neighbor[0]][new_neighbor[1]]
                    if nn_val -1 <= n_val:
                        visited[new_neighbor] = steps + 1
                        new_neighbors.append(new_neighbor)
        if len(new_neighbors) > 0:
            bfs(new_neighbors, steps + 1)
        
    bfs([start], 0)
    return visited[end]

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 31


##########
# Part 2 #
##########

def part2(data):
    data, _, end, p2_starts = process_data(data)
    visited = {}
    for starter in p2_starts: # populates visited list
        visited[starter] = 0
    directions = ((-1, 0), (1, 0), (0, -1), (0, 1))
    def bfs(neighbors, steps):
        new_neighbors = []
        for neighbor in neighbors:
            for direction in directions:
                n_val = data[neighbor[0]][neighbor[1]]
                new_neighbor = (neighbor[0]+direction[0], neighbor[1]+direction[1])
                if is_in_bounds(data, new_neighbor): # doesnt care if it has already been visited anymore
                    nn_val = data[new_neighbor[0]][new_neighbor[1]]
                    if nn_val -1 <= n_val:
                        if new_neighbor not in visited.keys() or visited[new_neighbor] > steps + 1: # stops if it has been visited with fewer steps
                            visited[new_neighbor] = steps + 1
                            new_neighbors.append(new_neighbor)
        if len(new_neighbors) > 0:
            bfs(new_neighbors, steps + 1)
    bfs(p2_starts, 0)
    return visited[end]

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 29


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
