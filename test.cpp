#include <iostream>
#include <unordered_set>
#include <utility>

int main() {
    std::unordered_set<std::pair<int,int>> set;
    set.insert(std::make_pair(1,1));
    return 0;
}