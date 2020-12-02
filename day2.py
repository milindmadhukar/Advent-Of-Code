import main

if __name__ == "__main__":
    data = main.getInputData(day=2)

    data = main.getSpiltList(data=data, splitChar=':')

    rangeVals = []
    charVals = []
    password_vals = []

    for line in data:
        rangeVals.append(list(map(int,line[0][0:-2].split('-'))))
        charVals.append(line[0][-1])
        password_vals.append(line[1][1:])

    # Part 1

    valid = 0
    for i in range(1000):
        count = 0
        for char in password_vals[i]:
            if char == charVals[i]:
                count+=1

        if count >= rangeVals[i][0] and count <= rangeVals[i][1]:
            valid += 1

    print(valid)

    # Part 2
    valid = 0
    for i in range(1000):
        lower_char = password_vals[i][(rangeVals[i][0]-1)]
        higher_char = password_vals[i][(rangeVals[i][1]-1)]
        if charVals[i] ==  lower_char or charVals[i] == higher_char:
            if lower_char != higher_char:
                valid += 1

    print(valid)