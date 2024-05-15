import re
import os

def atof(text):
    try:
        retval = float(text)
    except ValueError:
        retval = text
    return retval

def natural_keys(text):
    '''
    alist.sort(key=natural_keys) sorts in human order
    http://nedbatchelder.com/blog/200712/human_sorting.html
    (See Toothy's implementation in the comments)
    float regex comes from https://stackoverflow.com/a/12643073/190597
    '''
    return [ atof(c) for c in re.split(r'[+-]?([0-9]+(?:[.][0-9]*)?|[.][0-9]+)', text) ]

dir_list = os.listdir('./notes')
dir_list2 = os.listdir('./problems')

dir_list.sort(key=natural_keys)
dir_list2.sort(key=natural_keys)

to_write = ['# Notes for Leetcode\n', 'Welcome to my Leetcode Problems, these are my notes for the problems that I am solving<br>\n', '`python3 reloadReadme.py` to reload the readMe\n', '\n', '## Notes\n']
f = open('README.md', 'w')

for i in dir_list:
    click = i[:i.index('.')]
    file = f'./notes/{i}'
    to_write.append(f'[{click}]({file})<br>\n')

to_write.append('\n')
to_write.append('## Problems\n')

for i in dir_list2:
    click = i[:i.index('.')]
    file = f'./problems/{i}'
    to_write.append(f'[{click}]({file})<br>\n')

f.writelines(to_write)
