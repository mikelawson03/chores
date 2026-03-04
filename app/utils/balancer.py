from utils.models import Workload

def get_new_groups(chores):
    a_workload = Workload('A', 0)
    b_workload = Workload('B', 0)
    sorted_chores = sorted(chores, key=lambda c: c.time, reverse=True)
    deltas = []
    
    for chore in sorted_chores:
        new_group = set_rotation_group(chore, [a_workload, b_workload])
        if new_group == 'A':
            a_workload.time += chore.time
        else:
            b_workload.time += chore.time

        if new_group != chore.rotation_group:
            deltas.append((new_group, chore.chore_id))

    
    return deltas


def get_balance(workloads):
    a = 0
    b = 0
    
    for workload in workloads:
        if workload.rotation_group == 'A':
            a += workload.time
        elif workload.rotation_group == 'B':
            b += workload.time
    
    return (a, b)

def set_rotation_group(chore, workloads):
    (a, b) = get_balance(workloads)
    if a <= b:
        return 'A'
    else:
        return 'B'

def get_fairness_score(workloads):
    diff = 0
    higher = None

    (a, b) = get_balance(workloads)
    total = a + b
    if total == 0:
        return(1, higher, 0)
    if a > b:
        diff = a - b
    elif b > a:
        diff = b - a
    score = 1 - ( diff / total )
    res = {
        'score' : score,
        'diff' : diff,
        'a' : a,
        'b' : b
    }
    return res