from more_itertools import pairwise, windowed
import sys

input = [int(line) for line in sys.stdin]

print(
    list(map(lambda x: x[0] < x[1], pairwise(map(sum, windowed(input, 3))))).count(True)
)
