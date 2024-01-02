from functools import reduce

COLORS = ["red", "green", "blue"]

CUBES = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

def parse_input(f):
    with open(f) as f:
       lines = [line.rstrip() for line in f]
    return lines

lines = parse_input('input.txt')
powers = []
for l in lines:
    minimums = {c: 1 for c in COLORS}
    z = [x.split() for xs in l.split(":")[1].split(";") for x in xs.split(',')]
    
    for sample in z:
        num, col = sample
        if int(num) > minimums[col]:
            minimums[col] = int(num)
    powers += [reduce(lambda a, b: a * b, minimums.values())]

print(sum(powers))
