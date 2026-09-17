class DynamicArray {
    currentSize: number;
    capacity: number;
    arr: number[];
    /**
     * @constructor
     * @param {number} capacity
     */
    constructor(capacity: number) {
        this.capacity = capacity;
        this.currentSize = 0;
        this.arr = [];
    }

    /**
     * @param {number} i
     * @returns {number}
     */
    get(i: number): number {
        return this.arr[i];
    }

    /**
     * @param {number} i
     * @param {number} n
     * @returns {void}
     */
    set(i: number, n: number): void {
        this.arr[i] = n;
    }

    /**
     * @param {number} n
     * @returns {void}
     */
    pushback(n: number): void {
        if (this.currentSize == this.capacity) {
            this.resize();
        }
        this.arr[this.currentSize++] = n;
    }

    /**
     * @returns {number}
     */
    popback(): number {
        this.currentSize--;
        const val = this.arr[this.currentSize];
        return val;
    }

    /**
     * @returns {void}
     */
    resize(): void {
        this.capacity *= 2;
    }

    /**
     * @returns {number}
     */
    getSize(): number {
        return this.currentSize;
    }

    /**
     * @returns {number}
     */
    getCapacity(): number {
        return this.capacity;
    }
}
