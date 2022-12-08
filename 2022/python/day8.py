from utils import getInputData


def show_trees(trees):
    for line in trees:
        print(line)


data = getInputData(year=2022, day=8)
tree_grid = []

for line in data:
    row = []
    for tree in line:
        row.append(int(tree))
    tree_grid.append(row)

tree_grid_tr = [[tree_grid[j][i]
                 for j in range(len(tree_grid))] for i in range(len(tree_grid[0]))]

visible_trees = [[0 for i in range(len(tree_grid[0]))]
                 for j in range(len(tree_grid))]

# Marking all outer trees as visible
for row in range(len(visible_trees)):
    for column in range(len(visible_trees[row])):
        if row == 0 or row == len(visible_trees) - 1 or column == 0 or column == len(visible_trees[row]) - 1:
            visible_trees[row][column] = 1

# Part 1

# Looking from left
for row, row_val in enumerate(tree_grid):
    largest = tree_grid[row][0]
    for column, column_val in enumerate(row_val):
        if column_val > largest:
            visible_trees[row][column] = 1
            largest = column_val

# Looking from right
for row, row_val in enumerate(tree_grid):
    largest = tree_grid[row][-1]
    for column, column_val in enumerate(row_val[::-1]):
        if column_val > largest:
            visible_trees[row][len(row_val) - column - 1] = 1
            largest = column_val

# Looking from top
for row, row_val in enumerate(tree_grid_tr):
    largest = tree_grid_tr[row][0]
    for column, column_val in enumerate(row_val):
        if column_val > largest:
            visible_trees[column][row] = 1
            largest = column_val

# Looking from the bottom
for row, row_val in enumerate(tree_grid_tr):
    largest = tree_grid_tr[row][-1]
    for column, column_val in enumerate(row_val[::-1]):
        if column_val > largest:
            visible_trees[len(row_val) - column - 1][row] = 1
            largest = column_val

tree_count = 0
for line in visible_trees:
    tree_count += line.count(1)

print("Answer for Part 1 is", tree_count)

# Part 2

max_scenic_score = 0
for row, row_val in enumerate(tree_grid):
    for column, column_val in enumerate(row_val):
        if row == 0 or row == len(visible_trees) - 1 or column == 0 or column == len(visible_trees[row]) - 1:
            continue

        # Moving right
        right = 1
        for tree in row_val[column+1:]:
            if column_val > tree:
                right += 1
            else:
                break
        else:
            right -= 1

        # Moving left
        left = 1
        for tree in row_val[:column][::-1]:
            if column_val > tree:
                left += 1
            else:
                break
        else:
            left -= 1

        # Moving down
        down = 1
        for tree in tree_grid_tr[column][row + 1:]:
            if column_val > tree:
                down += 1
            else:
                break
        else:
            down -= 1

        # Moving up
        up = 1
        for tree in tree_grid_tr[column][:row][::-1]:
            if column_val > tree:
                up += 1
            else:
                break
        else:
            up -= 1

        scenic_score = up * down * left * right
        if scenic_score > max_scenic_score:
            max_scenic_score = scenic_score

print("Answer for Part 2 is", max_scenic_score)
