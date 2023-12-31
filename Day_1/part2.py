import re 

TEXT_NUMBERS = '(?=(one|two|three|four|five|six|seven|eight|nine|\d))'
DIGITS = {
    'zero': '0',
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9'
}

def parse_input(f):
    with open(f) as f:
        lines = f.readlines()
    return lines

def convert(a):
    return a if a.isdigit() else DIGITS[a]

def main(lines):
    total = 0
    for l in lines:
        res = re.findall(TEXT_NUMBERS, l.strip())
        a, b = res[0], res[-1]
        total += int(f'{convert(a)}{convert(b)}')
    return total

if __name__ == '__main__':
    lines = parse_input('input.txt')
    print(main(lines))
