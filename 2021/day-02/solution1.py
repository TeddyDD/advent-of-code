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
for line in lines:
    match line:
        case ("down", n):
            pos.depth += n
        case ("up", n):
            pos.depth -= n
        case ("forward", n):
            pos.x += n


print(pos, pos.result())
