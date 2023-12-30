def parse_input(f):
    with open(f) as f:
        lines = f.readlines()
    return lines

def find_numeric(line):
    for c in line:
        if c.isdigit():
            return c
        
def rfind_numeric(line):
    for c in line[::-1]:
        if c.isdigit():
            return c

def main(lines):
    nums = []
    for l in lines:
        s = find_numeric(l)
        s += rfind_numeric(l)
        nums.append(int(s))
    return sum(nums)
        
if __name__ == '__main__':
    lines = parse_input('input2.txt')
    print(main(lines))