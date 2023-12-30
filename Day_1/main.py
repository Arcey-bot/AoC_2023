def parse_input(f):
    with open(f) as f:
        lines = f.readlines()
    return lines

def build_int(line):
    return int(f'{find(line)}{find(line[::-1])}')

def find(line):
    for c in line:
        if c.isdigit():
            return c

def main(lines):
    return sum([build_int(l) for l in lines])
        
if __name__ == '__main__':
    lines = parse_input('input2.txt')
    print(main(lines))