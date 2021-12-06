import sys

fish = [int(n) for n in sys.stdin.readline().strip().split(",")]

def proc(all_fish):
    new = []
    for i, fish in enumerate(all_fish):
        match fish:
            case 0:
                all_fish[i] = 6
                new.append(8)
            case _:
                all_fish[i] -= 1
    all_fish += new


for i in range(80):
    proc(fish)


print(len(fish))
