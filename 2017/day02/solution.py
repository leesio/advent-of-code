import os 

def get_answer(n):
    rows = n.split('\n')
    total = 0

    for row in rows:
        if row == '':
            continue

        cells = row.split('\t')
        min = None
        max = None
        for cell in cells:
            cell_val = int(cell)
            if min == None or cell_val < min:
                min = cell_val
            if max == None or cell_val > max:
                max = cell_val

        delta = max - min
        total += delta
    return total

def get_answer_part_two(n):
    rows = n.split('\n')
    total = 0

    for row in rows:
        if row == '':
            continue

        cells = row.split('\t')

        for i, cell_a in enumerate(cells):
            for j, cell_b in enumerate(cells):
                if i == j:
                    continue
                cell_a_val = int(cell_a)
                cell_b_val = int(cell_b)
                if cell_a_val % cell_b_val == 0:
                    total += cell_a_val / cell_b_val
    return total

dir_path = os.path.dirname(os.path.realpath(__file__))
with open(os.path.join(dir_path, 'input'), 'r') as myfile:
    input=myfile.read()
    print(get_answer(input))
    print(get_answer_part_two(input))


