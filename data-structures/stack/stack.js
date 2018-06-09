class Stack {
    constructor() {
        this.store = [];
        this.top = 0;
    }

    push(item) {
        this.store[this.top++] = item;
    }

    pop() {
        if (this.top === 0){
            return -1;
        }

        return this.store.slice(--this.top, 1);
    }

    peak() {
        if (this.isEmpty()){
           return -1;
        }

        return this.store[this.top - 1];
    }

    isEmpty(){
        return this.top === 0;
    }

    getStack(){
        console.log(this.store, 1);
        return this.store;
    }

    length() {
       return this.top;
    }

}

const stack = new Stack();
stack.push('A');
stack.push('B');
stack.push('C');
console.log(stack.getStack());
stack.pop();
console.log(stack.getStack());
console.log(stack.peak());
console.log(stack.length());


