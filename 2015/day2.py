from .utils import getInputData
from .utils import getSpiltList


def surface_area(length, width, height):
    return 2 * ((length * width) + (width * height) + (height * length))


def smallest_side(length, width, height):
    return min(length * width, width * height, height * length)


def perimeter(length, width, height):
    return min((2 * (length + width), 2 * (length + height), 2 * (height + width)))


def volume(length, width, height):
    return length * width * height


if __name__ == "__main__":
    data = getInputData(year=2015, day=2)
    data = getSpiltList(data, "x", typecast=int)

    # Stage 1

    sum = 0
    for dimension in data:
        sum += surface_area(dimension[0], dimension[1], dimension[2]) + smallest_side(
            dimension[0], dimension[1], dimension[2]
        )
    print("Total Gift Paper Required", sum)

    # Stage 2

    sum = 0
    for dimension in data:
        sum += perimeter(dimension[0], dimension[1], dimension[2]) + volume(
            dimension[0], dimension[1], dimension[2]
        )

    print("Total Ribbon Needed", sum)
