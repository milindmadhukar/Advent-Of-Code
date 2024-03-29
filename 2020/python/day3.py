from utils import getInputData


def getFromData(data, x, y):
    return data[y][x % len(data[y])]


def getTreesEncountered(data, xStep, yStep):

    x, y, count = 0, 0, 0
    for _ in range(0, (len(data) - 1) // yStep):
        x += xStep
        y += yStep

        if getFromData(data, x, y) == "#":
            count += 1

    return count


if __name__ == "__main__":
    data = getInputData(year=2020, day=3)

    # Stage 1
    print("Number of trees encountered:", getTreesEncountered(data, 3, 1))
    # Stage 2
    print(
        "Product of trees encountered:",
        getTreesEncountered(data, 1, 1)
        * getTreesEncountered(data, 3, 1)
        * getTreesEncountered(data, 5, 1)
        * getTreesEncountered(data, 7, 1)
        * getTreesEncountered(data, 1, 2),
    )
