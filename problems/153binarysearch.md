# 153 

end game 

```c++

class Solution {
public:
    int findMin(vector<int>& nums) {
        int left = 0, right = nums.size() - 1;
        int minimum = 5004;
        while(left <= right){
            int tempMin = min(nums[left], nums[right]);
            minimum = min(tempMin, minimum);
            int mid = (left + right) / 2;
            minimum = min(minimum, nums[mid]);
            if(nums[mid] > minimum) left = mid + 1;
            else right = mid -1;
        }
        return minimum;
    }
};
```

This one was really easy just binary search with some extra checks;