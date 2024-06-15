# 706 

end game

```c++
class MyHashMap {
public:
    vector<int> map;
    MyHashMap() {
        for(int i = 0; i < 1000001; i++){
            map.push_back(-1);
        }
    }
    
    void put(int key, int value) {
        map[key] = value;
    }
    
    int get(int key) {
        return map[key];
    }
    
    void remove(int key) {
        map[key] = -1;
    }
};

/**
 * Your MyHashMap object will be instantiated and called as such:
 * MyHashMap* obj = new MyHashMap();
 * obj->put(key,value);
 * int param_2 = obj->get(key);
 * obj->remove(key);
 */
```

hashmap implementation really easy if we set the 100000 to like 1000 then we will need to deal with collisions which we can deal with by chaining.