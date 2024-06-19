# 981 

end game 

```c++

class TimeMap {
    unordered_map<string,vector<pair<string, int>>> mTime;
public:
    TimeMap() {
        
    }
    
    void set(string key, string value, int timestamp) {        
        mTime[key].push_back({value, timestamp});
    }
    
    string get(string key, int timestamp) {        
        int res = -1;
        if( mTime.find(key) != mTime.end()){            
            int start = 0, end = mTime[key].size()-1;
            
            while(start <= end){
                int mid = start + (end-start)/2;
                if(mTime[key][mid].second == timestamp){
                    return mTime[key][mid].first;
                }else if(mTime[key][mid].second < timestamp){
                    res = mid;
                    start = mid+1;
                }else
                    end = mid-1;
            }           
        }
        
        return (res != -1) ? mTime[key][res].first : "";
        
    }
};
```
this one I need to pay attention to thte checks I do because although I did binary search too many checks made my shit slow;