class Solution {
    /**
     * @param {number[]} prices
     * @return {number}
     */
    maxProfit(prices: number[]): number {
        let windowStart = -1;
        let windowEnd = -1;

        let lastMin = Infinity;
        let lastMaxProfit = 0;

        for (let i = 0; i < prices.length; i++) {
            if (windowStart === -1 || prices[i] < lastMin) {
                lastMin = prices[i];
                windowStart = i;
                windowEnd = i;
                let newProfit = 0;
                for (let j = i + 1; j < prices.length; j++) {
                    if (prices[j] - lastMin > newProfit) {
                        newProfit = prices[j] - lastMin;
                    }
                }

                if (newProfit > lastMaxProfit) {
                    lastMaxProfit = newProfit;
                }
            }
        }

        return lastMaxProfit
    }
}
