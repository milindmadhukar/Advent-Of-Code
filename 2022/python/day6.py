from utils import getInputData

contents = getInputData(year=2022, day=6, pure=True)

def check_unique_chars(till):
    for idx in range(len(contents)-till):
        received_chars = contents[idx:idx+till]
        if len(set(received_chars)) == len(received_chars):
            return idx + till

print("Answer for part 1", check_unique_chars(4))
print("Answer for part 2", check_unique_chars(14))

