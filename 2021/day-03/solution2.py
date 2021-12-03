import sys
from collections import Counter

digits = [x.strip() for x in sys.stdin]


def bitwise_not(n):
    return n ^ (2 ** n.bit_length() - 1)


def filt(elm, pos, wanted):
    if elm[pos] == wanted:
        return True
    return False


ox = digits.copy()
while len(ox) > 1:
    for i in range(len(ox[0])):
        if len(ox) <= 1:
            break
        counter = Counter([n[i] for n in ox])
        wanted = "0"
        if counter["1"] >= counter["0"]:
            wanted = "1"

        ox = list(filter(lambda elm: filt(elm, i, wanted), ox))

print(ox)

co2 = digits.copy()
while len(co2) > 1:
    for i in range(len(co2[0])):
        if len(co2) <= 1:
            break
        counter = Counter([n[i] for n in co2])
        wanted = "1"
        if counter["0"] <= counter["1"]:
            wanted = "0"
        co2 = list(filter(lambda elm: filt(elm, i, wanted), co2))

print(co2)

print(int(ox[0], 2) * int(co2[0], 2))
