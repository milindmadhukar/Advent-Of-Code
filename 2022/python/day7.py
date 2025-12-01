from utils import getInputData

class File:
    def __init__(self, size, name, loc):
        self.name = name
        self.size = size
        self.loc = loc

class Folder:
    def __init__(self, name, loc, files):
        self.name = name
        self.loc = loc
        self.files = files
        self.size = 0

def go_one_directory_behind(location):
    folders = location.split("/")
    return "/".join(folders[:-1])

def change_directory(loc):
    global curr_dir
    if loc == "/":
        curr_dir = ""
    elif loc == "..":
        curr_dir = go_one_directory_behind(curr_dir)
    else:
        if curr_dir == "":
            curr_dir += loc
        else:
            curr_dir += f"/{loc}"

data = getInputData(year=2022, day=7)
data.append("$")
curr_dir = ""

command_output_pairs = []
command_output = []

for line in data:
    if line.startswith('$') and command_output:
        command_output_pairs.append(command_output)
        command_output = []
    command_output.append(line)

folder_with_contents = {}

for pair in command_output_pairs:
    if pair[0][2:].startswith("cd"):
        change_directory(pair[0][5:])
    if pair[0][2:].startswith("ls"):
        folder_with_contents[curr_dir] = pair[1:]

fs = []
total_size = 0

for location, items in folder_with_contents.items():
    for item in items:
        if item.startswith('dir'):
            folder = Folder(item[4:], location, [])
            fs.append(folder)
        else:
            size, name = item.split(' ')
            file = File(int(size), name, location)
            fs.append(file)

# Putting files in their correct folders
for item in fs:
    if isinstance(item, File):
        if item.loc == '':
            total_size += item.size
            continue
        for folder in fs:
            if isinstance(folder, File):
                continue
            elif folder.loc == go_one_directory_behind(item.loc) and folder.name == item.loc.split('/')[-1]:
                folder.files.append(item)
                folder.size += item.size
                total_size += item.size
                break

fs.sort(key=lambda x: len(x.loc.split('/')), reverse=True)

# Putting folders in their correct folders
for item in fs:
    if isinstance(item, Folder):
        if item.loc == '':
            continue
        for folder in fs:
            if isinstance(folder, File):
                continue
            elif folder.loc == go_one_directory_behind(item.loc) and folder.name == item.loc.split('/')[-1]:
                folder.files.append(item)
                folder.size += item.size
                break

part_1_size = 0

for folder in fs:
    if isinstance(folder, Folder):
        if folder.size < 100000:
            part_1_size += folder.size

print("Answer for Part 1 is", part_1_size)

folder_sizes = [item.size for item in fs if isinstance(item, Folder)]
folder_sizes.sort()
folder_sizes.append(total_size)

unused_space = 70000000 - total_size
space_required = 30000000 - unused_space

for size in folder_sizes:
    if size > space_required:
        print("Answer for Part 2 is", size) # 2877389
        break