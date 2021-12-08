from itertools import groupby, product
import more_itertools
import sys

# R&D

sets = {
    0: set("abcefg"),
    2: set("acdeg"),
    3: set("acdfg"),
    5: set("abdfg"),
    6: set("abdefg"),
    9: set("abcdfg"),
    # known
    1: set("cf"),
    4: set("bcdf"),
    7: set("acf"),
    8: set("abcdefg"),
}


for e in product(sets.keys(), sets.keys()):
    if e[0] == e[1]:
        continue
    a = sets[e[0]]
    b = sets[e[1]]

    # if len(a) != len(b):
    #     continue

    d = a.difference(b)
    i = a.intersection(b)

    # print(e, len(i), len(d))


# test = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"
test = (
    "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"
)
(input, output) = test.split(" | ")
input = list(map(sorted, input.split(" ")))

input = sorted(input, key=len)

# Solution


def sort(lst):
    return sorted(map(lambda x: "".join(sorted(x)), lst), key=len)


def diff(a, b):
    return len(set(a).difference(set(b)))


def intr(a, b):
    return len(set(a).intersection(set(b)))


def parse_input(s):
    (input, output) = s.split(" | ")
    res = {}
    iter = groupby(sort(input.split(" ")), len)
    for k, v in iter:
        res[k] = list(v)
    res["output"] = list(map(lambda x: "".join(sorted(x)), output.split(" ")))
    return res


def decode(input):
    mapping = {
        1: input[2][0],
        4: input[4][0],
        7: input[3][0],
        8: input[7][0],
    }

    # figure out fives
    fives = input[5]

    three = list(filter(lambda x: diff(mapping[7], x) == 0, fives))[0]
    fives.remove(three)
    mapping[3] = three

    two = list(filter(lambda x: intr(mapping[4], x) == 2, fives))[0]
    fives.remove(two)
    mapping[2] = two

    # 5 left
    mapping[5] = fives[0]

    # figure out sixes
    sixes = input[6]

    six = list(filter(lambda x: intr(mapping[7], x) == 2, sixes))[0]
    sixes.remove(six)
    mapping[6] = six

    nine = list(filter(lambda x: diff(mapping[4], x) == 0, sixes))[0]
    sixes.remove(nine)
    mapping[9] = nine
    # 0 left
    mapping[0] = sixes[0]

    return dict((v, k) for k, v in mapping.items())


def decode_output(out, mapping):
    return list(map(lambda d: mapping[d], out))


def process_line(line):
    input = parse_input(line)
    print(input)
    output = input["output"]
    mapping = decode(input)
    res = int("".join(map(str, decode_output(output, mapping))))
    return res


input = [line.strip() for line in sys.stdin]

res = sum(map(process_line, input))
print(res)
