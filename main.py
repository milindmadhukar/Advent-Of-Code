from datetime import date
import requests
from datetime import date

def getInputData(year=2020, day=int(date.today().strftime('%d')), inputFileName=None, typecast=str):
    if inputFileName != None:
        f = open(inputFileName,'r')
        if '\n' in f.read():
            data = []
            for line in f.readlines():
                data.append(line[0:-1])

        else:
            data = f.readlines()
            f.close()
            return data
    
    else:
        session_key = open('sessionkey.txt', 'r').read()
        input_request = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies={'session': session_key}).text
        data = []
        tmp = ""
        if '\n' in input_request:
            for char in input_request:
                if char == '\n':
                    data.append(tmp)
                    tmp = ""

                else:
                    tmp += char
        else:
            data = input_request

    return list(map(typecast,data))

def getSpiltList(data,splitChar,typecast=str):
    split = []
    for line in data:
        split.append(list(map(typecast, line.split(splitChar))))

    return split

if __name__ == "__main__":
    print(getInputData())