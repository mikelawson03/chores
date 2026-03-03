def balance_load():
    return

def get_balance(workloads):
    a = 0
    b = 0
    for row in workloads:
        if row[0] == 'A':
            a += row[2]
        elif row[0] == 'B':
            b += row[2]
    return (a, b)

def set_rotation_group(chore, workloads):
    (a, b) = get_balance(workloads)
    if a <= b:
        chore.rotation_group = 'A'
    else:
        chore.rotation_group = 'B'