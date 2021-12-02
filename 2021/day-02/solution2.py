#!/usr/bin/env python3
import sys
from dataclasses import dataclass

lines = [line.strip().split(" ") for line in sys.stdin]
lines = list(map(lambda x: (x[0], int(x[1])), lines))


@dataclass
class Position:
    x: int = 0
    depth: int = 0

    def result(self):
        return self.x * self.depth


pos = Position()
aim = 0
for line in lines:
    match line:
        case ("down", n):
            aim += n
        case ("up", n):
            aim -= n
        case ("forward", n):
            pos.x += n
            pos.depth += aim * n


print(pos, pos.result())
