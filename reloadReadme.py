import os
dir_list = os.listdir('./notes')
dir_list2 = os.listdir('./problems')

dir_list.sort()
dir_list2.sort()

to_write = ['# Notes for Leetcode\n', 'Welcome to my Leetcode Problems, feel free to clone and just delete my work to do it yourself<br>\n', '`python3 reloadReadme.py` to reload the readMe\n', '\n', '## Notes\n']
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
