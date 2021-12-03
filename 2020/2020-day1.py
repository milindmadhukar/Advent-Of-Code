from utils import getInputData

if __name__ == "__main__":
    data = getInputData(year=2020, day=1, typecast=int)

    # Stage 1

    for i in data:
        for j in data:
            if i + j == 2020:
                print(f"{i} + {j} = 2020")
                print(i * j)
                break

    # Stage 2.
    print("\n")

    for i in data:
        for j in data:
            for k in data:
                if i + j + k == 2020:
                    print(f"{i} + {j} + {k} = 2020")
                    print(i * j * k)
                    break
