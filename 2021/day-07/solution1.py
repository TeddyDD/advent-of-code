from statistics import median
import sys

numbers = list(map(int, sys.stdin.readline().strip().split(",")))
m = median(numbers)
res = sum(map(lambda x: abs(x - m), numbers))
print(res)
