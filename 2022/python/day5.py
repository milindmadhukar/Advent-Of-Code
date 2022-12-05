from utils import getInputData

data = getInputData(year=2022, day=5)
split = data.index('')
crates, instructions = data[:split], data[split + 1:]

parsed_crates = []

for line in crates[:-1]:
    crate_line = []
    for i in range(1, len(line), 4):
        crate = line[i]
        crate_line.append(line[i])

    parsed_crates.append(crate_line)

part1_stacks = [[parsed_crates[j][i] for j in range(len(
    parsed_crates)) if parsed_crates[j][i] != ' '][::-1] for i in range(len(parsed_crates[0]))]

parsed_instructions = []
for instruction in instructions:
    split = instruction.split()
    parsed_instructions.append(list(map(int, [split[1], split[3], split[5]])))


# Part 1


def execute_part1_instruction(stacks, instruction):
    no_of_crates, start, end = instruction
    for _ in range(no_of_crates):
        crate = stacks[start - 1].pop()
        stacks[end-1].append(crate)
    return stacks


for instruction in parsed_instructions:
    part1_stacks = execute_part1_instruction(part1_stacks, instruction)

top_crates = "".join([stack[-1] for stack in part1_stacks])

print("Answer for part 1 is", top_crates)

# Part 2

part2_stacks = [[parsed_crates[j][i] for j in range(len(
    parsed_crates)) if parsed_crates[j][i] != ' '] for i in range(len(parsed_crates[0]))]


def execute_part2_instruction(stacks, instruction):
    no_of_crates, start, end = instruction
    crates = stacks[start-1][0:no_of_crates]
    stacks[end-1] = crates + stacks[end-1]
    stacks[start-1] = stacks[start-1][no_of_crates:]
    return stacks


for instruction in parsed_instructions:
    part2_stacks = execute_part2_instruction(part2_stacks, instruction)

top_crates = "".join([stack[0] for stack in part2_stacks])

print("Answer for part 2 is", top_crates)
