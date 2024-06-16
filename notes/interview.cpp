#include <string.h>
#include <iostream>
#include <stack>
#include <vector>

using namespace std;

class Solution
{
public:
    string filePath(string s)
    {
        string delimiter = "/";
        size_t pos = 0;
        string token;
        vector<string> split;
        stack<string> stack;
        while ((pos = s.find(delimiter)) != string::npos)
        {   
            token = s.substr(0, pos);
            s.erase(0, pos + delimiter.length());
            split.push_back(token);
        }
        split.push_back(s);
        for(string& s : split){
            if(s == ".."){
                stack.pop();
            }
            else{
                stack.push(s);
            }
        }
        string res = "";
        while(!stack.empty()){
            res = stack.top() + "/" + res;
            stack.pop();
        }
        res += "\n";
        return "." + res;
    }
};

int main()
{
    Solution sol;
    string s = "/home/user/Documents/../Pictures";
    string result = sol.filePath(s);
    cout << result;
    return 0;
}

// Input: path = "/home//foo/"
// Output: "/home/foo"

// Input: path = "/home/user/Documents/../Pictures"
// Input: path = "/home/user/../../Pictures"
// Output: "/home/user/Pictures"

// Input: path = "/.../a/../b/c/../d/./"

// Output: "/.../b/d"

// :D :-)