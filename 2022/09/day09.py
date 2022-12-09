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

def tail_is_too_far(head, tail):
    if abs(head[0] - tail[0]) > 1 or abs(head[1] - tail[1]) > 1:
        return True
    return False


def move(head, tail, direction):
    new_head = (head[0] + direction[0], head[1] + direction[1])
    if tail_is_too_far(new_head, tail):
        tail = head
    return new_head, tail


def part1(data):
    head_visits = set()
    tail_visits = set()
    move_dir = {
        "U": (1,0),
        "D": (-1,0),
        "R": (0,1),
        "L": (0,-1)
    }
    head = (0,0)
    tail = (0,0)
    for step in data:
        for i in range(int(step[1])):
            head, tail = move(head, tail, move_dir[step[0]])
            head_visits.add(head)
            tail_visits.add(tail)
    return len(tail_visits)

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 13


##########
# Part 2 #
##########

def move_two(snake, direction):
    diag = False
    diag_direction = (0, 0)
    snake[0] = (snake[0][0] + direction[0], snake[0][1] + direction[1])
    for i in range(1, len(snake)):
        leader, follower = snake[i-1], snake[i]
        if tail_is_too_far(leader, follower):
            if leader[0] == follower[0]: # move LR 
                diag = False # being "in line" cancels/resets diagonal movement
                if leader[1] < follower[1]: # move L
                    snake[i] = (follower[0], follower[1] - 1)
                else: # move R
                    snake[i] = (follower[0], follower[1] + 1)
            
            elif leader[1] == follower[1]: # move UD
                diag = False # being "in line" cancels/resets diagonal movement
                if leader[0] < follower[0]: # move D
                    snake[i] = (follower[0] - 1, follower[1])
                else: # move U
                    snake[i] = (follower[0] + 1, follower[1])
            else: # diagonal
                if diag: # if follower needs to move after the previous movement was diagonal, it moves in the same direction
                    snake[i] = (follower[0] + diag_direction[0], follower[1] + diag_direction[1])
                else:
                    ud = 0
                    lr = 0
                    if leader[0] - follower[0] > 0:
                        ud = 1
                    else:
                        ud = -1
                    if leader[1] - follower[1] > 0:
                        lr = 1
                    else:
                        lr = -1
                    diag_direction = (ud, lr)
                    snake[i] = (follower[0] + diag_direction[0], follower[1] + diag_direction[1])
        else:
            break
    return snake


def print_map(tail_visits):
    h = [0, 0]
    w = [0, 0]
    for visit in tail_visits:     
        h = [min(h[0], visit[0]), max(h[1], visit[0])]
        w = [min(w[0], visit[1]), max(w[1], visit[1])]
    for i in range(h[1], h[0]-1, -1):
        line = ""
        for j in range(w[0], w[1]+1):
            if (i,j) == (0,0):
                line += "s"
            elif (i,j) in tail_visits:
                line += "#"
            else:
                line += "."
        print(line)


def part2(data):
    tail_visits = set()
    move_dir = {
        "U": (1,0),
        "D": (-1,0),
        "R": (0,1),
        "L": (0,-1)
    }
    snake = [(0,0) for _ in range(10)]

    for step in data:
        for _ in range(int(step[1])):
            snake = move_two(snake, move_dir[step[0]])
            tail_visits.add(snake[-1])
    # print_map(tail_visits)
    return len(tail_visits)

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 1
    testData = "test_input2.txt"
    data = loadData(testData)
    assert part2(data) == 36


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
