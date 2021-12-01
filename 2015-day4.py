from main import getInputData
import hashlib
import threading
from functools import partial


def checkStart(start, num):
    n = num
    while True:
        key = data[0] + str(n)
        md5_hex = hashlib.md5(key.encode()).hexdigest()
        if md5_hex.startswith(start):
            print(md5_hex)
            print(key)
            exit()
        # print(out, end='\r')
        print(n, end="\r")
        n += 1


if __name__ == "__main__":
    data = getInputData(year=2015, day=4)

    t1 = threading.Thread(target=partial(checkStart, "000000", 0))
    t2 = threading.Thread(target=partial(checkStart, "000000", 10000))
    t3 = threading.Thread(target=partial(checkStart, "000000", 20000))
    t4 = threading.Thread(target=partial(checkStart, "000000", 30000))
    t5 = threading.Thread(target=partial(checkStart, "000000", 40000))

    t1.start()
    t2.start()
    t3.start()
    t4.start()
    t5.start()

    t1.join()
    t2.join()
    t3.join()
    t4.join()
    t5.join()

