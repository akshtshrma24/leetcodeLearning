# 12

endgame 

```c++

class Solution {
public:
    string intToRoman(int num) {    
        int values[] = {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
        string symbols[] = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
        string res = "";
        for(int i = 0; i < sizeof(values)/sizeof(*values); i++){
            while(num >= values[i]){
                num -= values[i];
                res += symbols[i];
            }
        }
        return res;
    }
};
```
this one is easy once you know the trick otherwise i have another solution where you could just use multiples and a stack
where the first one of the stack will have multiple of 1 so then you hard code 9 4 5 1 etc. then just have it that many times
