"""
Input: path = "/home//foo/"
Output: "/home/foo"

Input: path = "/home/user/Documents/../Pictures"
Input: path = "/home/user/../../Pictures"
Output: "/home/user/Pictures"

Input: path = "/.../a/../b/c/../d/./"

Output: "/.../b/d"

:D :-)
"""

def pee(stri):
    res = ""
    p = stri.split('/')
    correct_arr = []
    for i in p:
        if(i == '..'): correct_arr.pop()
        else: correct_arr.append(i + '/')
    for i in correct_arr:
        res += i
    if(res[-1] == '/'): res = res[:-1]
    return "." + res

