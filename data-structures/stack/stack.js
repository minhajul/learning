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

        return this.store[--this.top];
    }

    peak() {
        return this.store[this.top - 1];
    }

    length() {
       return this.top;
    }

    clear() {
       this.top = 0
    }

}

const stack = new Stack();
stack.push('A');
stack.push('B');
stack.push('C');
console.log(stack.store);
stack.pop();
console.log(stack.length());
console.log(stack.peak());
console.log(stack.clear());


