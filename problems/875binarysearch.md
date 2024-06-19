# 875 

end game

```c++
class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int right = -1;
        int left = 1;
        for(int& i : piles){
            right = max(right, i);
        }

        while(left < right){
            int mid = (left + right) / 2;
            int total = 0;
            for(int i : piles)
            {
                total = total + ceil((double)i / mid);
            }
            if(total <= h){
                right = mid;
            }else{
                left = mid + 1;
            }
        }
        return left;
    }
};
```

this one is a bit different because I need to keep the bounds of int in mind