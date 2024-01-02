COLORS = ["red", "green", "blue"]

CUBES = {
    "red": 12,
    "green": 13,
    "blue": 14,
}

SAMPLE = """\
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"""

# def cube_max(s):
#     maximal = {c : 0 for c in COLORS}
#     upd = lambda a: int(a[0]) if int(a[0]) > maximal[a[1]] else maximal[a[1]]
#     # tokens = [list(map(lambda x: (x[1], int(x[0])), [v.split() for v in st.split(',')])) for st in s.split(';')]
#     tokens = [map(lambda x: (x[1], int(x[0])), [v.split() for v in st.split(',')]) for st in s.split(';')]

#     print(f'{tokens=}')

#     for st in tokens:
#         for sample in st:
#             if sample[1] > maximal[sample[0]]:
#                 maximal[sample[0]] |= int(sample[1])

#     # print(tokens)
#     return maximal


# Assuming games are always given in order this is fine, probably breaks for p2
# games = [(i + 1, v) for i, v in enumerate([cube_max(l) for l in SAMPLE.split("\n")])]
# print(games)

# possible = []
# print(possible)

def parse_input(f):
    with open(f) as f:
       lines = [line.rstrip() for line in f]
    return lines


lines = parse_input('input.txt')

possible = 0
for i, l in enumerate(lines, 1):
    # z = list(map(lambda a: a.split(), [yss for yss in ys.split(',')]) for ys in l.split(":")[1].split(";"))
    z = [x.split() for xs in l.split(":")[1].split(";") for x in xs.split(',')]
    # print(f'{z=}')
    
    for sample in z:
        num, col = sample
        if int(num) > CUBES[col]:
            print(f'Impossible game {i}, {num} {col} > maximum {CUBES[col]}')
            break  # impossible game
    else:
        possible += i

print(possible)
