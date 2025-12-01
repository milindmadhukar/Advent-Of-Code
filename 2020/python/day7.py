from utils import getInputData, getSpiltList

colours_containing_shiny_gold = []
def search_for_shiny_gold(current_color:str) -> bool:
    if current_color == "shiny gold" or current_color in colours_containing_shiny_gold:
        return True
    if current_color == "no other":
        return False
    bags_contained = bag_data[current_color]
    for _, colour in bags_contained:
        isFound = search_for_shiny_gold(colour)
        if isFound:
            colours_containing_shiny_gold.append(current_color)
            return True
    else:
        return False

def sum_of_bags(current_bag = "shiny gold", prev_quanity = 1) -> int:
    sum = 0
    bags_contained = bag_data.get(current_bag, [])
    for quantity, colour in bags_contained:
        sum += quantity * prev_quanity
        sum += sum_of_bags(colour, quantity * prev_quanity)
    return sum
    
if __name__ == "__main__":
    data = getInputData(year=2020, day=7)
    data = getSpiltList(data, " contain ")

    bag_data = dict()

    for line in data:
        line[0] = line[0][:-5]
        line[1] = line[1][:-1].split(', ')
        bags = []
        for contained_bags in line[1]:
            first_space = contained_bags.index(" ")
            quantity = 0
            if contained_bags[:first_space] != "no":
                quantity = int(contained_bags[:first_space])
            else:
                first_space -= 3
            last_space = len(contained_bags) - 1 - contained_bags[::-1].index(" ")
            colour = contained_bags[first_space+1:last_space]
            bags.append([quantity, colour])

        bag_data[line[0]] = bags
    
    count = 0
    
    for colour in bag_data.keys():
        if search_for_shiny_gold(colour) and colour != "shiny gold":
            count += 1

    print("Answer for Part 1:", count)

    sum = sum_of_bags()
    print("Answer for Part 2:", sum)
