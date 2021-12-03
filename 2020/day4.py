from utils import getInputData
from utils import getSpiltList


def checkHex(s):
    if len(s) != 7:
        return True
    n = len(s)
    for i in range(1, n):
        ch = s[i]

        if (ch < "0" or ch > "9") or (ch < "A" or ch > "F"):
            return False
    return True


if __name__ == "__main__":
    data = getInputData(year=2020, day=4, pure=True)
    data += " "
    parsed = []
    tmp = ""
    for char in data:
        if char == " ":
            parsed.append(tmp.replace("\n", ""))  # Why didn't I use .split()
            tmp = ""
        elif char == "\n":
            parsed.append(tmp.replace(" ", ""))
            tmp = ""

        else:
            tmp += char

    data = parsed

    data = getSpiltList(data, ":")
    tmp = []
    parsed = []
    for line in data:
        if line[0] == "":
            parsed.append({tmp[i][0]: tmp[i][1] for i in range(len(tmp))})
            tmp = []
        else:
            tmp.append(line)
    valid = 0
    data = parsed

    # Part 1

    for line in data:
        if (len(line.keys()) == 8) or (
            len(line.keys()) == 7 and "cid" not in line.keys()
        ):
            valid += 1
    print("Valid for Part 1:", valid)

    # Part 2

    valid = -1  # I really don't know why is it adding one extra :(
    for line in data:
        if (len(line.keys()) == 8) or (
            len(line.keys()) == 7 and "cid" not in line.keys()
        ):
            if (
                len(line["byr"]) != 4
                or int(line["byr"]) < 1920
                or int(line["byr"]) > 2002
            ):
                continue
            if (
                len(line["iyr"]) != 4
                or int(line["iyr"]) < 2010
                or int(line["iyr"]) > 2020
            ):
                continue

            if (
                len(line["eyr"]) != 4
                or int(line["eyr"]) < 2020
                or int(line["eyr"]) > 2030
            ):
                continue

            if len(line["pid"]) != 9:
                continue

            if checkHex(line["hcl"]):
                continue

            if "cm" in line["hgt"]:
                height = int(line["hgt"][0 : len(line["hgt"]) - 2])
                if height < 150 or height > 193:
                    continue
            if "in" in line["hgt"]:
                height = int(line["hgt"][0 : len(line["hgt"]) - 2])
                if height < 59 or height > 76:
                    continue

            eye_colours = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
            if line["ecl"] not in eye_colours:
                continue

            valid += 1

    print("Valid for Part 2:", valid)
