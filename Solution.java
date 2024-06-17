class Solution {
public:
    int partitionString(string s) {
        set<char> set;
        int res = 1;
        for(char& c : s){
            int prev_size = set.size();
            set.insert(c);
            if(prev_size == set.size()){
                res++;
                set.clear();
                set.insert(c);
            }
        }
        return res;
    }
};