from utils import getInputData

if __name__ == "__main__":
    data = getInputData(year=2015, day=3)

    x = 0
    y = 0

    visited = [(x, y)]

    for direction in data:
        if direction == "^":
            y += 1

        elif direction == "v":
            y -= 1

        elif direction == ">":
            x += 1

        elif direction == "<":
            x -= 1

        visited.append((x, y))

    print("Santa visited", len(set(visited)), "houses")

    # Stage 2

    santa_x = 0
    santa_y = 0

    robo_santa_x = 0
    robo_santa_y = 0

    visited = [(santa_x, santa_y), (robo_santa_x, robo_santa_y)]

    for i in range(len(data)):
        if data[i] == "^":
            if i % 2 == 0:
                santa_y += 1
            else:
                robo_santa_y += 1

        elif data[i] == "v":
            if i % 2 == 0:
                santa_y -= 1
            else:
                robo_santa_y -= 1

        elif data[i] == ">":
            if i % 2 == 0:
                santa_x += 1
            else:
                robo_santa_x += 1

        elif data[i] == "<":
            if i % 2 == 0:
                santa_x -= 1
            else:
                robo_santa_x -= 1

        if i % 2 == 0:
            visited.append((santa_x, santa_y))
        else:
            visited.append((robo_santa_x, robo_santa_y))

    print("Both visited", len(set(visited)), "houses")
