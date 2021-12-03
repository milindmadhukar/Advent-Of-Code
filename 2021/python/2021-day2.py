from utils import getInputData, getSpiltList

data = getInputData(year=2021)
data = getSpiltList(data, " ")

horizontal = 0
depth = 0

for instruction in data:
    direction = instruction[0]
    instruction[1] = int(instruction[1])
    value = instruction[1]
    if direction == "forward":
        horizontal += value
    elif direction == "down":
        depth += value
    elif direction == "up":
        depth -= value

print("Part one: ", horizontal * depth)

horizontal = 0
depth = 0
aim = 0

for instruction in data:
    direction = instruction[0]
    value = instruction[1]
    if direction == "forward":
        horizontal += value
        depth += value * aim

    elif direction == "down":
        aim += value
    elif direction == "up":
        aim -= value


print("Part two: ", horizontal * depth)
