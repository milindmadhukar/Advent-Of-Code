from main import getInputData

def getTreesEncountered(data, xStep, yStep):

    x,y,count = 0,0,0
    for _ in range(0,(len(data)-1)//yStep):
        x += xStep
        y += yStep

        if data[y][x] == "#":
            count += 1
    
    return count
    
if __name__ == "__main__":
    data = getInputData()

    for i in range(len(data)):
        data[i] += data[i]*200 #Hardcode Gang yeahhhh

    # Stage 1
    print("Number of trees encountered:",getTreesEncountered(data,3,1))
    # Stage 2
    print("Product of trees encountered:", getTreesEncountered(data,1,1)*getTreesEncountered(data,3,1)*getTreesEncountered(data,5,1)*getTreesEncountered(data,7,1)*getTreesEncountered(data,1,2))