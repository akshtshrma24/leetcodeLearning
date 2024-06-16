# 56 

end game 

```c++

class Solution {
public:
    vector<vector<int>> merge(vector<vector<int>>& intervals) {
        int start = -1;
        int end = -1;
        vector<vector<int>> res;
        sort(intervals.begin(), intervals.end(), [](const vector<int>& a, const vector<int>& b) {
            return a[0] < b[0];
        });
        for(vector<int>& i : intervals){
            if(end < i[0]){
                if(start != -1) res.push_back({start,end});
                start = i[0];
                end = i[1];
            }else if(end >= i[0]){
                if(end <= i[1]) end = i[1];
            }
        }
        res.push_back({start,end});
        return res;
    }
};
```

this one is easy as well make sure the the intervals array is sorted