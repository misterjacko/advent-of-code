def readFile():
    datalist = []
    datafile = open("./data/day#data.txt", "r")
        for aline in datafile:
            transactions.append(int(aline))
    datafile.close()
    return datalist

def part1():
    return ('part1')

def part2():
    return ("part2")

def test():
    test_input = []
    assert part1(test_input) == 'part1'
    assert part2(test_input) == 'part2'

if __name__ == "__main__":
    test()
    vals = readFile()
    print(f"Part 1: {part1(vals)}")
    print(f"Part 2: {part2(vals)}")