from datetime import datetime


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]


def process_data(data):
    seeds = [int(number) for number in data[0].split(": ")[1].split()]
    map_sets = {}
    map_name = ""
    i = 2
    while i < len(data):
        if data[i] == "":
            pass
        elif data[i].split()[0].isdigit():
            values = [int(val) for val in data[i].split()]
            map_sets[map_name].append(values)
        else:
            map_name = data[i].split()[0]
            map_sets[map_name] = []
        i += 1
    
    # format dict
    for k, v in map_sets.items():
        #sort by start range
        v.sort(key=lambda x: x[1])
        formatted_v = []
        if v[0][1] != 0:
            formatted_v.append([0,0])
        for i, rng in enumerate(v):
            formatted_v.append([rng[1],rng[0]-rng[1]]) 
            if i+1 >= len(v) or v[i+1][1] != rng[1]+rng[2]:
                formatted_v.append([rng[1]+rng[2], 0])
        map_sets[k] = formatted_v
    return seeds, map_sets


def calcualte_diff(in_value, tranfer_table):
    for i in range(len(tranfer_table)-1):
        if in_value < tranfer_table[i+1][0]:
            return in_value+tranfer_table[i][1]
    return in_value


def calcualte_diff2(in_range_list, tranfer_table):
    new_ranges = []
    for rng in in_range_list:
        range_min = min(rng)
        range_max = max(rng)

        for i in range(len(tranfer_table)-1):
            # current range contains at least some of transfer
            if range_min < tranfer_table[i+1][0]:
                # all within one range
                if range_max < tranfer_table[i+1][0]:
                    new_ranges.append(range(range_min+tranfer_table[i][1], range_max+tranfer_table[i][1]+1))
                    range_min = float('inf')
                # split range in two
                else:
                    # add lower half to new ranges
                    new_ranges.append(range(range_min+tranfer_table[i][1], (tranfer_table[i+1][0]-1)+tranfer_table[i][1]+1))
                    range_min = tranfer_table[i+1][0]
        if range_min != float('inf'):
            new_ranges.append(range(range_min, range_max+1))
    return new_ranges

##########
# Part 1 #
##########

def part1(seeds, maps):
    final_map = {}
    for seed in seeds:
        final_map[seed] = {}
        final_map[seed]["soil"] = calcualte_diff(seed, maps["seed-to-soil"])
        final_map[seed]["fertilizer"] = calcualte_diff(final_map[seed]["soil"], maps["soil-to-fertilizer"])
        final_map[seed]["water"] = calcualte_diff(final_map[seed]["fertilizer"], maps["fertilizer-to-water"])
        final_map[seed]["light"] = calcualte_diff(final_map[seed]["water"], maps["water-to-light"])
        final_map[seed]["temperature"] = calcualte_diff(final_map[seed]["light"], maps["light-to-temperature"])
        final_map[seed]["humidity"] = calcualte_diff(final_map[seed]["temperature"], maps["temperature-to-humidity"])
        final_map[seed]["location"] = calcualte_diff(final_map[seed]["humidity"], maps["humidity-to-location"])

    lowest_location = float('inf')
    for v in final_map.values():
        lowest_location = min(lowest_location, v["location"])
    return lowest_location

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    seeds, maps = process_data(data)
    assert part1(seeds, maps) == 35


##########
# Part 2 #
##########

def part2(seeds, maps):
    # this takes ~8 minutes and that isnt great... 
    lowest_location = float('inf')
    seed_ranges = []
    i = 0
    while i < len(seeds):
        seed_ranges.append(range(seeds[i], seeds[i]+seeds[i+1]))
        i+=2
    for i, rng in enumerate(seed_ranges):
            print(f"starting {i+1} of {len(seed_ranges)}")
 
            soil = calcualte_diff2([rng], maps["seed-to-soil"])
            fertilizer = calcualte_diff2(soil, maps["soil-to-fertilizer"])
            water = calcualte_diff2(fertilizer, maps["fertilizer-to-water"])
            light = calcualte_diff2(water, maps["water-to-light"])
            temperature = calcualte_diff2(light, maps["light-to-temperature"])
            humidity = calcualte_diff2(temperature, maps["temperature-to-humidity"])
            location = calcualte_diff2(humidity, maps["humidity-to-location"])
            for rng in location:
                lowest_location = min(lowest_location, min(rng))
    return lowest_location

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    seeds, maps = process_data(data)
    assert part2(seeds, maps) == 46


if __name__ == "__main__":
    dataFile = "input.txt"
    startTime = datetime.now()
    data = loadData(dataFile)
    seeds, maps = process_data(data)
    p1 = part1(seeds, maps)
    if p1:
        print(f"Part 1 Solution: {p1}")
        print(f"Part 1 Exe Time: {datetime.now() - startTime}")

    p2 = part2(seeds, maps)
    if p2:
        print(f"Part 2 Solution: {p2}")
        print(f"Part 2 Exe Time: {datetime.now() - startTime}")
