from utils import getInputData

data = getInputData(year=2021, typecast=int)
print(
    "Answer to part 1:",
    len([data[i] for i in range(0, len(data) - 1) if data[i + 1] > data[i]]),
)

print(
    "Answer to part 2:",
    len(
        [
            i
            for i in range(0, len(data) - 1)
            if data[i] + data[i + 1] + data[i + 2]
            < data[i + 1] + data[i + 2] + data[i + 3]
        ]
    ),
)
