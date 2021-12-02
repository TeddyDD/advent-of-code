#!/usr/bin/env python3
import sys
from dataclasses import dataclass

lines = [(line[0], int(line[1:])) for line in sys.stdin]


@dataclass
class Ship:
    x: int = 0
    y: int = 0
    h: int = 90  # east

    def rotate(self, deg):
        self.h = (self.h + deg) % 360

    def do_op(self, op):
        match op:
            case ("N", n):
                self.y -= n
            case ("S", n):
                self.y += n
            case ("E", n):
                self.x += n
            case ("W", n):
                self.x -= n


ship = Ship()

for op in lines:
    match op:
        case ("F", n):
            match ship.h:
                case 0:
                    ship.do_op(("N", n))
                case 90:
                    ship.do_op(("E", n))
                case 180:
                    ship.do_op(("S", n))
                case 270:
                    ship.do_op(("W", n))
        case ("R", n):
            ship.rotate(n)
        case ("L", n):
            ship.rotate(-n)
        case op:
            ship.do_op(op)
    print(op, ship)

print(ship, abs(ship.x) + abs(ship.y))
