'''
    Author: @Alex
    Date: 2022.1.13
    read names from file or string var
    and mkdir in batch in current directory
'''
import io
import os
text = """base
comparisons
constants
control
conversions
extended
loads
math
stack
stores"""

def mkdir_all(list):
    for e in list:
        os.mkdir(e)


buf = io.StringIO(text)
file_names = buf.readlines()
for idx,s in enumerate(file_names):
    file_names[idx]=s.rstrip('\n')
print(file_names)
mkdir_all(file_names)