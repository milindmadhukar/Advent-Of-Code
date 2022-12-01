from datetime import date
import requests


def getInputData(
    year,
    day=int(date.today().strftime("%d")),
    inputFileName=None,
    typecast=str,
    pure=False,
):
    if inputFileName is not None:
        f = open(inputFileName, "r")

        data = f.read().strip()
        f.close()
        return data

    else:
        session_key = open("sessionkey.txt", "r").read().strip()
        input_request = requests.get(
            f"https://adventofcode.com/{year}/day/{day}/input",
            cookies={"session": session_key},
        ).text
        data = []
        tmp = ""

        if pure:
            return input_request
        if "\n" in input_request:
            for char in input_request:
                if char == "\n":
                    data.append(tmp)
                    tmp = ""

                else:
                    tmp += char
        else:
            data = input_request

    return list(map(typecast, data))


def getSpiltList(data, splitChar, typecast=str):
    split = []
    for line in data:
        split.append(list(map(typecast, line.split(splitChar))))

    return split


if __name__ == "__main__":
    print(getInputData())