# 11. Container With Most Water

## Input Format

```text
n
h1 h2 h3 ... hn
```

## Example

```text
9
1 8 6 2 5 4 8 3 7
```

Output:

```text
49
```

## Idea

Use two pointers. The area is limited by the shorter side, so move the pointer on the shorter side inward.

Time complexity: `O(n)`

Space complexity: `O(1)`
