# 146

end game

```c++
    class LRUCache {
    public:
        class Node{
            public: 
                int key;
                int val;
                Node* prev;
                Node* next;

                Node(int key, int val){
                    this->key = key;
                    this->val = val;
                }
        };

        Node* head = new Node(-1, -1);
        Node* tail = new Node(-1, -1);

        int cap;
        unordered_map<int, Node*> m;

        LRUCache(int capacity) {
            cap = capacity;
            head -> next = tail;
            tail -> prev = head;
        }

        void addNode(Node* newnode){
            Node* temp = head -> next;

            newnode -> next = temp;
            newnode -> prev = head;

            head -> next = newnode;
            temp -> prev = newnode;
        }

        void deleteNode(Node* delnode){
            Node* prevv = delnode -> prev;
            Node* nextt = delnode -> next;

            prevv -> next = nextt;
            nextt -> prev = prevv;
        }
        
        int get(int key) {
            if(m.find(key) != m.end()){
                Node* resNode = m[key];
                int ans = resNode -> val;

                m.erase(key);
                deleteNode(resNode);
                addNode(resNode);

                m[key] = head -> next;
                return ans;
            }
            return -1;
        }
        
        void put(int key, int value) {
            if(m.find(key) != m.end()){
                Node* curr = m[key];
                m.erase(key);
                deleteNode(curr);
            }

            if(m.size() == cap){
                m.erase(tail -> prev -> key);
                deleteNode(tail -> prev);
            }

            addNode(new Node(key, value));
            m[key] = head -> next;
        }
    };
```


```py
class Node: 
    def __init__(self, key, val):
        self.key, self.val = key,val
        self.prev = self.next = None

class LRUCache:

    def __init__(self, capacity: int):
        self.cap = capacity
        self.cache = {}
        # left is LR, right ms MR
        self.left, self.right = Node(0,0), Node(0,0)
        self.left.next, self.right.prev = self.right, self.left

    # remove from the list
    def remove(self, node):
        prev, nxt = node.prev, node.next
        prev.next, nxt.prev = nxt, prev

    # insert at right
    def insert(self, node):
        prev, nxt = self.right.prev, self.right
        prev.next = nxt.prev = node
        node.next, node.prev = nxt, prev

    def get(self, key: int) -> int:
        if key in self.cache: 
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            return self.cache[key].val
        return -1 

    def put(self, key: int, value: int) -> None:
        if key in self.cache: 
            self.remove(self.cache[key])
        self.cache[key] = Node(key, value)
        self.insert(self.cache[key])

        if len(self.cache) > self.cap:
            lru = self.left.next
            self.remove(lru)
            del self.cache[lru.key]

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
```
this one is aids but its just a hashmap of doubley linked lists