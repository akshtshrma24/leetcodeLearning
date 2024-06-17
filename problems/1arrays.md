# 1 

end game 

```c++
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> res;
        map<int, int> map;
        for(int i = 0; i < nums.size(); i++){
            if(map.find(target - nums[i]) == map.end()){
                map[nums[i]] = i;
             }else{
                res.push_back(map[ target - nums[i]]);
                res.push_back(i);
                break;
            }
        }
        return res;
    }
};
```
hash map 