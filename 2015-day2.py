from main import getInputData
from main import getSpiltList

if __name__ == "__main__":
    data = getInputData(year=2015,day=2)
    data = getSpiltList(data,'x', typecast=int)

    # Stage 1
    surface_area = lambda l,w,h : 2*((l*w) + (w*h) + (h*l))
    smallest_side = lambda l,w,h : min(l*w, w*h, h*l)

    sum = 0
    for dimension in data:
        sum += surface_area(dimension[0],dimension[1],dimension[2]) + smallest_side(dimension[0],dimension[1],dimension[2])
    print("Total Gift Paper Required", sum)

    # Stage 2

    perimeter = lambda l,w,h: min(( 2*(l+w), 2*(l+h), 2*(h+w) ))
    volume = lambda l,w,h : l*w*h
    
    sum = 0
    for dimension in data:
        sum += perimeter(dimension[0],dimension[1],dimension[2]) + volume(dimension[0],dimension[1],dimension[2])

    print("Total Ribbon Needed",sum)