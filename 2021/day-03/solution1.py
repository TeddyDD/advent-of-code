import sys
from collections import Counter

digits = [x.strip() for x in sys.stdin]

def bitwise_not(n):
    return n ^ (2 ** n.bit_length() - 1)

gamma = ""

for i in range(len(digits[0])):
    c = Counter([d[i] for d in digits])
    if c["0"] > c["1"]:
        gamma += "0"
    else:
        gamma += "1"
    print(c)

gamma = int(gamma, 2)
epsilon = bitwise_not(gamma)

print(gamma, epsilon, gamma * epsilon)
