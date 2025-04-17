# 88 

real end game 


```py
class Solution:
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """

        last = m + n - 1

        while m > 0 and n > 0:
            if nums1[m - 1] > nums2[n - 1]:
                nums1[last] = nums1[m - 1]
                m -= 1
            else:
                nums1[last] = nums2[n - 1]
                n -= 1
            last -= 1

        while n > 0: 
            nums1[last] = nums2[n - 1]
            n -= 1
            last -= 1
```
this one was not that easy
but you di it like this
[1,2,3,0,0,0]
[2,5,6]
you start at the end of the first one and the real end of it as well so 3 and the end of 2nd one so 6 then you just
put the bigger element in place of the last variable then you decremenet the bigger ones variable 
keep doing that but there is an edge case where you can reach the end of m but there is n still so you just have a seperate loop to check 
