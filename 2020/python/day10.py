from utils import getInputData

def calculate_differences():
    difference_of_ones = 0
    difference_of_threes = 0

    current_rating = 0

    for i in range(len(data)):
        current_adapter = data[i]

        joltage_difference = current_adapter - current_rating

        if joltage_difference == 1:
            difference_of_ones += 1
        elif joltage_difference == 3:
            difference_of_threes += 1
        
        current_rating = current_adapter

    return difference_of_ones * difference_of_threes

def generate_combinations():
    given_adapters = data

    current_combinations = [[given_adapters[0]]]

    while True:
        existing_combinations = current_combinations.copy()
        last_elements = [combination[-1] for combination in existing_combinations]
        if all([element == given_adapters[-1] for element in last_elements]):
            break
        valid_combinations = []
        for combination in existing_combinations:
            next_valid_adapters = []
            last = combination[-1]
            if last == given_adapters[-1]:
                valid_combinations.append(combination)
                continue
            if (last + 1) in given_adapters:
                next_valid_adapters.append(last+1)
            if last + 2 in given_adapters:
                next_valid_adapters.append(last+2)
            if last + 3 in given_adapters:
                next_valid_adapters.append(last+3)

            for adapter in next_valid_adapters:
                new_combination = combination.copy()
                new_combination.append(adapter)
                valid_combinations.append(new_combination)

        current_combinations = valid_combinations.copy()
        print("Currently at", len(current_combinations), "combinations")
    return current_combinations

def part_two(data) -> int:
    # there's always at least one way to do this (so we start with 1)
    # also, from the example, the length of each arrangment is shorter than the max (last) value in the list
    points = [1] + [0] * (data[-1])

    # as we go through the numbers, we can track the number of distinct times the individual branch is used as part of the previous branches
    # then we use those values to add to the next branch
    for r in data:
        points[r] = points[r - 1] + points[r - 2] + points[r - 3]

    # the last value will be the result
    return points[-1]

if __name__ == "__main__":
    data = getInputData(year=2020, day=10, typecast=int)
    data.sort()

    print("Number of adapters:", len(data))

    device_joltage_rating = data[-1] + 3

    data.append(device_joltage_rating)

    print("Adapters", data)

    difference_product = calculate_differences()

    print("Answer for Part 1:", difference_product)

    print("Answer for Part 2:", part_two(data))

    # generated_combinations = generate_combinations()
    
    # print("Answer for Part 2:", len(generated_combinations))


