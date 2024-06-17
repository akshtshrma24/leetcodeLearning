# 75

end game

```c++
class Solution {
public:
    void sortColors(vector<int>& nums) {
        // Buckt sort
        int zeros = 0;
        int ones = 0;
        int twos = 0;
        for(int& i : nums){
            if(i == 0) zeros++;
            else if(i == 1) ones++;
            else twos++;
        }
        for(int i = 0; i < zeros; i++){
            nums[i] = 0;
        }
        for(int i = zeros; i < zeros + ones; i++){
            nums[i] = 1;
        }
        for(int i = ones + zeros; i < zeros + ones + twos; i++){
            nums[i] = 2;
        }
    }
};
```
Easy bucket sort