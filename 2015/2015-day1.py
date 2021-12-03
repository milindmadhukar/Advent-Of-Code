from main import getInputData

if __name__ == "__main__":
    data = getInputData(year=2015, day=1)
    # Stage 1
    floor = 0
    for char in data:
        if char == "(":
            floor += 1

        elif char == ")":
            floor -= 1

    print(floor)

    # Stage 2

    floor = 0
    for i in range(len(data)):
        if data[i] == "(":
            floor += 1

        elif data[i] == ")":
            floor -= 1

        if floor == -1:
            print("Reached Basement at character", (i + 1))
            break

