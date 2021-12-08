import sys
import math


def cost(n):
    return sum(range(1, n + 1))


numbers = list(sorted(map(int, sys.stdin.readline().strip().split(","))))
# This is hack
pos = math.floor(sum(numbers) / len(numbers))
res = sum(map(lambda x: cost(abs(x - pos)), numbers))
print(res, pos)
