# 121 

end game

```c++
class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if(prices.size() == 1) return 0;

        int buy = 0, sell = 1, max_profit = 0;
        while(sell < prices.size()){
            if(prices[buy] > prices[sell]){
                buy = sell;
            }
            else if(prices[buy] < prices[sell]){
                max_profit = max(max_profit, prices[sell] - prices[buy]);
            }
            sell++;
        }
        return max_profit;
    }
};
```

this one is easy af, just have a window that slides through the bes tpossible ones;