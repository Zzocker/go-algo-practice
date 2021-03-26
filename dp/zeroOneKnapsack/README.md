# 0-1 Knapsack Problem

- Subset Sum
- Equal Sum Partition
- Count of Subset Sum
- Minimum Subset Sum Difference
- Target Sum
- Number of subset with given difference 

## General Knapsack Problem

given capacity of a knapsack (beg), we have to choose items (weight,profit) such that profit is maximum
```
    items  : I1 , I2 , I3 , I4 , I5
    weight : W1 , W2 , W3 , W4 , W5

    capacity of knapsack = W
```

### Types

- Fractional Knapsack : (greedy)
    Items can be fractioned and then place in beg
- 0-1 Knapsack : (DP)
    Item cannot be fractioned , means an item can be placed or not.
- Unbounded Knapsack: (DP)
    Each Item has unlimited supply
