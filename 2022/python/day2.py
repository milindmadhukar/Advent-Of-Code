from utils import getInputData

data = getInputData(year=2022, day=2)

parsed_data = []

for item in data:
    parsed_data.append(item.split())

sign = {'X': 1, 'Y':2, 'Z': 3}

total_score = 0

for combo in parsed_data:
    if (combo[0] == 'A' and combo[1] == 'Y') or (combo[0] == 'B' and combo[1] == 'Z') or (combo[0] == 'C' and combo[1] == 'X'):
        total_score += 6
        total_score += sign[combo[1]]
    elif (combo[0] == 'A' and combo[1] == 'Z') or (combo[0] == 'B' and combo[1] == 'X') or (combo[0] == 'C' and combo[1] == 'Y'):
        total_score += sign[combo[1]]
    else:
        total_score += 3
        total_score += sign[combo[1]]

print("Answer for part 1", total_score)

total_score = 0

for combo in parsed_data:
    if combo[1] == 'Z':
        total_score += 6
        if combo[0] == 'A':
            total_score += 2 # Opponent chose rock, we chose paper, we win.
        elif combo[0] == 'B':
            total_score += 3
        elif combo[0] == 'C':
            total_score += 1

    elif combo[1] == 'Y':
        total_score += 3
        if combo[0] == 'A':
            total_score += 1
        elif combo[0] == 'B':
            total_score += 2
        elif combo[0] == 'C':
            total_score += 3
    elif combo[1] == 'X':
        if combo[0] == 'A':
            total_score += 3 # Opponent chose rock, we chose scissor, we lose.
        elif combo[0] == 'B':
            total_score += 1
        elif combo[0] == 'C':
            total_score += 2

print("Answer for part 2", total_score)

        
