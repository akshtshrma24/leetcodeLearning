# 232

THIS IS THE ENDGAME

```c++
class MyQueue {
    stack<int> in; 
    stack<int> out;
public:
    MyQueue() {
        
    }
    
    void push(int x) {
        in.push(x);
    }
    
    int pop() {
        int front = this->peek();
        if(front == -1){
            return -1;
        }
        out.pop();
        return front;
    }
    
    int peek() {
        if(out.empty()){
            while(!in.empty()){
                out.push(in.top());
                in.pop();
            }
        }
        return out.empty() ? -1 : out.top() ;
    }
    
    bool empty() {
        return out.size() == 0 && in.size() == 0;
    }
};

/**
 * Your MyQueue object will be instantiated and called as such:
 * MyQueue* obj = new MyQueue();
 * obj->push(x);
 * int param_2 = obj->pop();
 * int param_3 = obj->peek();
 * bool param_4 = obj->empty();
 */
```

so this you use 2 stacks
you push to the in stack and pop from both but the 
peek function works like this:
    you check if out is empty if it is, pop all elements from the other stack 
    adding it to this one, then return the top(), its like reversing it twice
