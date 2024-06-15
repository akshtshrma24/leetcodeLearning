# 268 

end game

```c++

class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int max_range = -1;
        int sum_nums = 0;
        for(const int& i: nums){
            max_range = max(max_range, i);
            sum_nums += i;
        }
        if(max_range + 1 == nums.size()) return max_range + 1;
        int sum_real_nums = 0;
        for(int i = 0; i <= max_range; i++){
            sum_real_nums += i;
        }
        return sum_real_nums - sum_nums;
    }
};

```
THis is super easy
