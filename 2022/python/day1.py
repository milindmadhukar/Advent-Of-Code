from utils import getInputData

data = getInputData(year=2022, day=1)
data.append('')

# Part 1

elf_data = []
elf_items = []
for item in data:
    if item == '':
        elf_data.append(elf_items)
        elf_items = []
        continue
    elf_items.append(int(item))

# print(elf_data)

elf_sums = [sum(calories) for calories in elf_data]

print("Answer for part 1 is", max(elf_sums))

elf_sums.sort(reverse=True)

print("Answer for part 2 is", sum(elf_sums[0:3]))
