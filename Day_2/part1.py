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

possible = 0
for i, l in enumerate(lines, 1):
    z = [x.split() for xs in l.split(":")[1].split(";") for x in xs.split(',')]
    
    for sample in z:
        num, col = sample
        if int(num) > CUBES[col]:
            print(f'Impossible game {i}, {num} {col} > maximum {CUBES[col]}')
            break  # impossible game
    else:
        possible += i

print(possible)
