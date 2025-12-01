
with open("day7.txt") as f:
    data = [int(num) for num in f.readline().strip().split(",")]

fuels = []
for fin_pos in range(max(data)):
    fuel = 0
    for init_pos in data:
        fuel += abs(fin_pos - init_pos)
    fuels.append(fuel)

print(min(fuels))

additorial = lambda n : (n*n+1)//2
fuels = []
for fin_pos in range(max(data)):
    fuel = 0
    for init_pos in data:
        fuel += (additorial(abs(fin_pos - init_pos)))
    fuels.append(fuel)

print(min(fuels))
