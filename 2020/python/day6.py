from utils import getInputData, getSpiltList

if __name__ == "__main__":
    data = getInputData(year=2020, day=6)
    data.append("")

    tmp = []
    groups = []

    for line in data:
        if line == "":
            groups.append(tmp)
            tmp = []
        else:
            tmp.append(line)
    
    sum1 = 0
    sum2 = 0

    for group in groups:
        questions = ""
        for question in group:
            questions += question

        sum1 += len(set(questions))
    
    for group in groups:
        question_sets = []
        for questions in group:
            question_sets.append(set(questions))

        sum2 += len(question_sets[0].intersection(*question_sets[1:]))


    print("Answer for Part 1:", sum1)

    print("Answer for Part 2:", sum2)
    
