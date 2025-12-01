with open('day6.txt') as f:
    data = [num for num in f.readline().strip().split(',')]

fishes_dict = dict()
for i in range(9):
    fishes_dict[str(i)] = 0

for fish in data:
    fishes_dict[fish] += 1

for day in range(1,257):
    fishes_dict_cpy = fishes_dict.copy()
    for days_left in range(8, 0, -1):
        fishes_dict[str(days_left-1)] = fishes_dict_cpy[str(days_left)]
    fishes_dict['8'] = fishes_dict_cpy['0']
    fishes_dict['6'] += fishes_dict_cpy['0']

sum = 0
for val in fishes_dict.values():
    sum += val
print(sum)
