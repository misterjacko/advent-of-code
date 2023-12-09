from datetime import datetime
from collections import Counter


#############
# Prep Work #
#############

def loadData(dataFile):
    with open(dataFile, "r") as f:
        lines = f.readlines()
    return [line.strip() for line in lines]


def map_hand(hand_map):
    if len(hand_map) == 1:
        return 5
    elif len(hand_map) == 2: # 1:4 or 2:3
        if 1 in hand_map.values():
            return 4
        else:
            return 3.2
    elif len(hand_map) == 3: # 3:1:1 or 2:2:1
        if 3 in hand_map.values():
            return 3
        else:
            return 2.2
    elif len(hand_map) == 4: 
        return 2
    else:
        return 1


def map_joker_hand(hand_map):
    if "J" in hand_map.keys():
        jokers = hand_map["J"]
        if jokers == 5:
            return 5
        del hand_map["J"]
        # add jokers to highest count
        hand_map[sorted(hand_map.items(), key=lambda x:x[1])[-1][0]] += jokers
    return map_hand(hand_map)


def format_cards(cards, joker=False):
    formatted_cards = ""
    value_map = {
        "2": "b",
        "3": "c",
        "4": "d",
        "5": "e",
        "6": "f",
        "7": "g",
        "8": "h",
        "9": "i",
        "T": "j",
        "J": "k",
        "Q": "l",
        "K": "m",
        "A": "n",
    }
    if joker:
        value_map["J"] = "a"
    
    for card in list(cards):
        formatted_cards = f"{formatted_cards}{value_map[card]}"
    return formatted_cards

def sort_groups(cards_in_group):
    return sorted(cards_in_group, key=lambda x: (x[0]))


##########
# Part 1 #
##########

def part1(data):
    total_winnings = 0
    multiplier = 1
    groups = {5: [], 4: [], 3.2: [], 3: [], 2.2: [], 2: [], 1: []}
    for hand in data:
        value = int(hand.split()[1])
        raw_hand = dict(Counter(hand.split()[0]))
        sortable_cards = format_cards(hand.split()[0])
        groups[map_hand(raw_hand)].append([sortable_cards, value]) 

    group_order = sorted(groups)
    print(groups)
    for group in group_order: # smallest to largest
        sorted_group = sort_groups(groups[group]) # smallest to largest
        for hand in sorted_group:
            total_winnings += hand[1]*multiplier
            multiplier+=1
        
    return total_winnings

def test_part1():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part1(data) == 644000


##########
# Part 2 #
##########

def part2(data):
    total_winnings = 0
    # sorting will go from smallest to largest so will the multiplier
    multiplier = 1
    groups = {5: [], 4: [], 3.2: [], 3: [], 2.2: [], 2: [], 1: []}
    for hand in data:
        value = int(hand.split()[1])
        raw_hand = dict(Counter(hand.split()[0]))
        sortable_cards = format_cards(hand.split()[0], True)
        if "J" in hand.split()[0]:
            groups[map_joker_hand(raw_hand)].append([sortable_cards, value]) 
        else:
            groups[map_hand(raw_hand)].append([sortable_cards, value]) 

    group_order = sorted(groups)
    for group in group_order: # smallest to largest
        sorted_group = sort_groups(groups[group]) # smallest to largest
        for hand in sorted_group:
            total_winnings += hand[1]*multiplier
            multiplier+=1
        
    return total_winnings

def test_part2():
    testData = "test_input.txt"
    data = loadData(testData)
    assert part2(data) == 5905


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
