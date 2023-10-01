from utils import getInputData

# Two numbers in nums should add up to target
def twoSum(target, nums):
    nums = set(nums)
    for num in nums:
        if target - num in nums:
            return True
    return False

def find_contiguos_set(num):
    for lower in range(len(data)):
        for upper in range(len(data)):
            if lower >= upper:
                continue
            
            contiguous_set = data[lower:upper]
            
            if sum(contiguous_set) == num:
                return contiguous_set

if __name__ == "__main__":
    data = getInputData(year=2020, day=9, typecast=int)
#     inp = """35
# 20
# 15
# 25
# 47
# 40
# 62
# 55
# 65
# 95
# 102
# 117
# 150
# 182
# 127
# 219
# 299
# 277
# 309
# 576""".split()
#     data = list(map(int, inp))

    preamble_size = 25
    num = 0
    for i in range(preamble_size, len(data)):
        preamble = data[i-preamble_size:i]
        num = data[i]
        if not twoSum(num, preamble):
            break

    
    print("Answer for Part 1:", num)

    contiguous_set = find_contiguos_set(num)
    contiguous_set.sort()

    print("Answer for Part 1:", contiguous_set[0] + contiguous_set[-1]) 


