from utils import getInputData, getSpiltList

def get_accumulator():

    for line in data:
        line[2] = False

    accumulator = 0
    
    line_number = 0;
    while line_number < len(data):
        if data[line_number][2]:
            break

        data[line_number][2] = True


        match(data[line_number][0]):
            case 'jmp':
                line_number += data[line_number][1] - 1
            case 'acc':
                accumulator += data[line_number][1]

        line_number += 1

    
    didTerminateCorrectly = line_number == len(data)
    return accumulator, didTerminateCorrectly

def fix_instruction_bruteforce():
    for i in range(len(data)):
        if data[i][0] == 'acc':
            continue
        if data[i][0] == 'jmp':
            data[i][0] = 'nop'
        elif data[i][0] == 'nop':
            data[i][0] = 'jmp'

        accumulator, didTerminateCorrectly = get_accumulator()
        if didTerminateCorrectly:
            break
        else:
            if data[i][0] == 'jmp':
                data[i][0] = 'nop'
            elif data[i][0] == 'nop':
                data[i][0] = 'jmp'

    return accumulator

if __name__ == "__main__":
    data = getInputData(year=2020, day=8)
    data = getSpiltList(data, " ")

    for line in data:
        line[1] = int(line[1])
        line.append(False)


    accumulator, _ = get_accumulator()

    print("Answer for Part 1:", accumulator)

    accumulator = fix_instruction_bruteforce()
    
    print("Answer for Part 2:", accumulator)
