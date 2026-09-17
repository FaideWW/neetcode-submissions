interface LLNode {
    next: LLNode | null;
    val: number;
}

class LinkedList {
    head: LLNode | null;
    constructor() {
        this.head = null;
    }

    /**
     * @param {number} index
     * @return {number}
     */
    get(index: number): number {
        if (this.head === null) return -1;
        let curr = this.head;
        for (let i = 0; i < index; i++) {
            curr = curr.next;
            if (curr === null) return -1;
        }
        return curr.val;
    }

    /**
     * @param {number} val
     * @return {void}
     */
    insertHead(val: number): void {
        this.head = {
            val,
            next: this.head,
        };
    }

    /**
     * @param {number} val
     * @return {void}
     */
    insertTail(val: number): void {
        if (this.head === null) {
            this.head = {
                val,
                next: null,
            }
            return;
        }

        let curr = this.head;
        while (curr.next !== null) {
            curr = curr.next;
        }
        curr.next = {
            val,
            next: null,
        };
    }

    /**
     * @param {number} index
     * @return {boolean}
     */
    remove(index: number): boolean {
        if (index === 0) {
            if (this.head === null) {
                return false;
            }
            this.head = this.head.next;
            return true;
        }

        let prev = this.head;
        for (let i = 1; i < index; i++) {
            prev = prev.next;
            if (prev === null) return false;
        }

        if (prev.next === null) return false;
        prev.next = prev.next.next;
        return true;
    }

    /**
     * @return {number[]}
     */
    getValues(): number[] {
        const vals = [];
        let curr = this.head;
        while (curr !== null) {
            vals.push(curr.val);
            curr = curr.next;
        }

        return vals;
    }
}
