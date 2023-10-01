from typing import List

with open("day5.txt") as f:
    data = [
        [list(map(int, coord.split(","))) for coord in line.strip().split("->")]
        for line in f.readlines()
    ]


def get_points_covered(xmin: int, ymin: int, xmax: int, ymax: int) -> List:
    covered_points = []

    for x in range(xmin, xmax + 1):
        for y in range(ymin, ymax + 1):
            covered_points.append([x, y])

    return covered_points


def get_diagonals(x1: int, y1: int, x2: int, y2: int) -> List:
    xmin = min(x1, x2)
    xmax = max(x1, x2)
    ymin = min(y1, y2)
    ymax = max(y1, y2)
    covered_points = []
    if (abs(y2 - y1) / abs(x2 - x1)) == 1:
        print("here", x1, y1, x2, y2)
        covered_points.append([xmin, ymin])
        covered_points.append([xmax, ymax])

        for x in range(xmin + 1, xmax + 1):
            for y in range(ymin + 1, ymax + 1):
                if x == y:
                    covered_points.append([x, y])

    return covered_points


def isHorizontalOrVertical(x1: int, y1: int, x2: int, y2: int) -> bool:
    if x1 == x2 or y1 == y2:
        return True
    return False


def get_matrix(xmax: int, ymax: int) -> List:
    return [[0 for _ in range(xmax + 1)] for _ in range(ymax + 1)]


matrix = get_matrix(9, 9)

for line in data:
    first, second = line
    x1, y1 = first
    x2, y2 = second
    xmin = min(x1, x2)
    xmax = max(x1, x2)
    ymin = min(y1, y2)
    ymax = max(y1, y2)
    if isHorizontalOrVertical(x1, y1, x2, y2):
        points_covered = get_points_covered(xmin, ymin, xmax, ymax)
        for point in points_covered:
            x, y = point
            matrix[y][x] += 1

count = 0

for line in matrix:
    for value in line:
        if value >= 2:
            count += 1

print(count)

matrix = get_matrix(9, 9)

for line in data:
    first, second = line
    x1, y1 = first
    x2, y2 = second
    xmin = min(x1, x2)
    xmax = max(x1, x2)
    ymin = min(y1, y2)
    ymax = max(y1, y2)
    points_covered = []
    if isHorizontalOrVertical(x1, y1, x2, y2):
        points_covered = get_points_covered(xmin, ymin, xmax, ymax)
    else:
        diagonals = get_diagonals(x1, y1, x2, y2)
        for point in diagonals:
            points_covered.append(point)
    for point in points_covered:
        x, y = point
        matrix[y][x] += 1

for line in matrix:
    print(line)

count = 0

for line in matrix:
    for value in line:
        if value >= 2:
            count += 1

print(count)
