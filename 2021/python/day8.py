with open("day8.txt") as f:
    data = []
    for line in f.readlines():
        signal, output = line.split("|")
        data.append([list(map(set, signal.split())), output.split()])

count = 0
for line in data:
    for digit in line[1]:
        if len(digit) in [2, 4, 3, 7]:
            count += 1

print(count)


display_char = {
    "t": "",
    "b": "",
    "tl": "",
    "tr": "",
    "bl": "",
    "br": "",
    "m": "",
}


def get_seven(sets):
    for setVal in sets:
        if len(setVal) == 3:
            return setVal


def get_one(sets):
    for setVal in sets:
        if len(setVal) == 2:
            return setVal


def get_four(sets):
    for setVal in sets:
        if len(setVal) == 4:
            return setVal


def get_eight(sets):
    for setVal in sets:
        if len(setVal) == 7:
            return setVal


def get_zero_six_and_nine(sets):
    vals = []
    for setVal in sets:
        if len(setVal) == 6:
            vals.append(setVal)
    return vals


def get_two_three_and_five(sets):
    vals = []
    for setVal in sets:
        if len(setVal) == 5:
            vals.append(setVal)
    return vals

def get_number_from_codes(codes: str, display_char: dict):
    number = []
    for code in codes:
        sl = {
            "t": False,
            "b": False,
            "tl": False,
            "tr": False,
            "bl": False,
            "br": False,
            "m": False,
        }
        for char in code:
            sl[list(display_char.keys())[list(display_char.values()).index(char)]] = True 

        # elif sl["t"] and sl["tl"] and sl["bl"] and sl["b"] and sl["br"] and sl["tr"] and sl["m"]:
        if all(sl.values()):
            number.append("8")
        elif sl["t"] and sl["tl"] and sl["bl"] and sl["b"] and sl["br"] and sl["tr"] and not sl["m"]:
            number.append("0")
        elif not sl["t"] and not sl["tl"] and not sl["bl"] and not sl["b"] and sl["br"] and sl["tr"] and not sl["m"]:
            number.append("1")
        elif sl["t"] and not sl["tl"] and sl["bl"] and sl["b"] and not sl["br"] and sl["tr"] and sl["m"]:
            number.append("2")
        elif sl["t"] and not sl["tl"] and not sl["bl"] and sl["b"] and sl["br"] and sl["tr"] and sl["m"]:
            number.append("3")
        elif not sl["t"] and sl["tl"] and not sl["bl"] and not sl["b"] and sl["br"] and sl["tr"] and sl["m"]:
            number.append("4")
        elif sl["t"] and sl["tl"] and not sl["bl"] and sl["b"] and sl["br"] and not sl["tr"] and sl["m"]:
            number.append("5")
        elif sl["t"] and sl["tl"] and sl["bl"] and sl["b"] and sl["br"] and not sl["tr"] and sl["m"]:
            number.append("6")
        elif sl["t"] and not sl["tl"] and not sl["bl"] and not sl["b"] and sl["br"] and sl["tr"] and not sl["m"]:
            number.append("7")
        elif sl["t"] and sl["tl"] and not sl["bl"] and sl["b"] and sl["br"] and sl["tr"] and sl["m"]:
            number.append("9")

    return int("".join(number))

numbers = []

for line in data:
    seven = get_seven(line[0])
    one = get_one(line[0])
    display_char["t"] = list((seven - one))[0]

    four = get_four(line[0])
    zerosixnine = get_zero_six_and_nine(line[0])
    zero = None
    six = None
    nine = None
    for number in zerosixnine:
        # print("num-val", number)
        print("\n\n\n\n\n")
        print("zsn", zerosixnine)
        print("four =", four)
        print("seven =", seven)
        print("number =", number)
        # print("sevenfour", seven | four)
        # print("ZSN", zerosixnine)
        if len(number - (four | seven)) == 1:
            print("Why did I come here?", len(number), len(number - (four | seven)), number - (four | seven))
            nine = number
            display_char["b"] = list(number - (four | seven))[0]
            
        if len(four - number) == 1:
            zero = number
            display_char["m"] = list(four-number)[0]

    eight = get_eight(line[0])
    print("nine", nine)
    display_char["bl"] = list(eight - nine)[0]
    display_char["tr"] = list(eight - six)[0]

    display_sum = set(display_char["t"] + display_char["b"]+ display_char["tr"] + display_char["bl"] + display_char["m"])

    display_char["br"] = list(seven - display_sum)[0]

    display_sum = set(display_char["t"] + display_char["b"]+ display_char["tr"] + display_char["bl"] + display_char["br"] + display_char["m"])

    display_char["tl"] = list(eight- display_sum)[0]

    numbers.append(get_number_from_codes(line[1], display_char))

print(sum(numbers))


