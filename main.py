from datetime import date
import requests
from datetime import date

def getInputData(year=2020, day=int(date.today().strftime('%d')), inputFileName=None, typecast=str):
    if inputFileName != None:
        f = open(f'{inputFileName}.txt','r')
        data = []
        for line in f.readlines():
            data.append(line[0:-1])
        f.close()
    
    else:
        session_key = open('sessionkey.txt', 'r').read()
        input_request = requests.get(f"https://adventofcode.com/{year}/day/{day}/input", cookies={'session': session_key}).text
        data = []
        tmp = ""
        for char in input_request:
            if char == '\n':
                data.append(tmp)
                tmp = ""

            else:
                tmp += char

    return list(map(typecast,data))

def getSpiltList(data,splitChar):
    split = []
    for line in data:
        split.append(line.split(splitChar))

    return split

if __name__ == "__main__":
    print(getInputData())