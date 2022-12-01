from utils import getInputData

data = getInputData(year=2022, day=3)

parsed_data = [] 
for item in data:
    half = len(item)//2
    parsed_data.append([item[:half], item[half:]])

# Part 1
    
d1 = {chr(i): i-96 for i in range(97, 123)}
d2 = {chr(i) : i-38 for i in range(65,91)}
d1.update(d2)
priority_values = d1

def get_common_item(pair):
    for char in pair[0]:
        if char in pair[1]:
            return char
    
priority_sum = 0
for pair in parsed_data:
    common_item = get_common_item(pair)
    priority_sum += priority_values[common_item]
    
print("Answer for part 1 is", priority_sum)


parsed_data = [] 
priority_value = 0
for i in range(0, len(data), 3):
    parsed_data.append([data[i], data[i+1], data[i+2]])

def get_common_item_in_three(group):
    for char in group[0]:
        if char in group[1] and char in group[2]:
            return char

for group in parsed_data:
    common_item = get_common_item_in_three(group)
    priority_value += priority_values[common_item]

print("Answer for part 2 is", priority_value)
