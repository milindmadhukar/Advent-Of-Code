from utils import getInputData

def get_pos(arrangement, min, max, upper_char, lower_char):
    for pos in arrangement:
        mid = (min + max) // 2
        if pos == upper_char:
            max = mid
        elif pos == lower_char:
            min = mid + 1

    return max

if __name__ == "__main__":
    data = getInputData(year=2020, day=5)

    max_id = 0

    seats = dict()
    for row in range(128):
        for col in range(7):
            seats[(row, col)] = False

    for boarding_pass in data:
        row = get_pos(boarding_pass[:7], 0, 127, "F", "B")
        col = get_pos(boarding_pass[7:], 0, 7, "L", "R")
        seats[(row, col)] = True
        id = row * 8 + col
        if id > max_id:
            max_id = id


    missing_seats = []
    for row_col, isFull in seats.items():
        if not isFull and row_col[0] != 0 and row_col[0] != 127:
            missing_seats.append(row_col)

    missing_seats_ids = [seat[0] * 8 + seat[1] for seat in missing_seats]

    print("Answer for Part 1:", max_id)

    print("Answer for Part 2:", missing_seats_ids[0])
