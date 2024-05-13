# 1143 

```py
class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        # make a 2d array and fill it in with 0s with 1 to right and 1 below for the 0th row
        dp = [[0 for j in range(len(text2) + 1)] for i in range(len(text1) + 1)]
        # loop through the 2d array backwards
        for i in range(len(text1) - 1, -1, -1):
            for j in range(len(text2) - 1, -1, -1):
                # if it is the same then diagonal value + 1
                if text1[i] == text2[j]:
                    dp[i][j] = dp[i + 1][j + 1] + 1
                # otherwise you get the max of column right and row bottom
                else:
                    dp[i][j] = max(dp[i + 1][j], dp[i][j + 1])
        # return 0th 
        return dp[0][0]

```

This one is like the one where you start from the middle and expand out, 
My initial thought is you start from middle in both if it is in both you add to ""
expand out keep expanding out.

No its not like that, its bottom up 2 dimensional dp, cant be solved like that 
what you do is you go from right down most position, if it is common, then you take the value that is diagonal 
to it then add it to current one if they dont match you look right and down then take max do that til you get 
to top