from utils.sql_handlers import write_chore_to_table, get_chores_page, get_all_chores, get_workloads, update_rotation_group
from utils.balancer import get_new_groups, set_rotation_group, get_fairness_score
from utils.models import Chore
from tabulate import tabulate
import os, sys

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def add_chore():
    assignee = None
    time = 0

    cls()
    print('Add chore')
    print('-------------- \n')
    name = input('Chore name > ')
    cadence = input('Cadence (Daily/Weekly/Monthly) > ').lower()

    while cadence not in ['d', 'daily', 'w', 'weekly', 'm', 'monthly']:
        print('Invalid choice. Please choose daily, weeekly, or monthly')
        cadence = input('Cadence (Daily/Weekly/Monthly) > ').lower()
    
    if cadence == 'daily':
        cadence = 'd'
    elif cadence == 'weekly':
        cadence = 'w'
    elif cadence == 'monthly':
        cadence = 'm'
        
    shared_choice = input('Is this chore shared (Y/N)? ').lower()

    while shared_choice not in ['y', 'n', 'yes', 'no']:
        print('Invalid choice. Please choose either Y or N')
        shared_choice = input('Is this chore shared (Y/N)? ').lower()

    if shared_choice in ['y', 'yes']:
        shared = True;
    else:
        shared = False

    if not shared :
        assignee = input('Who should this chore be assigned to? ').lower()

    while isinstance(time, int) and time <= 0:
        try:
            time = int(input('How long in minutes should this chore take? '))
        except:
            print('Invalid choice. Please enter a number greater than 0')

    new_chore = Chore(name, cadence, shared, assignee, time)
    
    # assign to balanced group 
    if shared:
        workloads = get_workloads(cadence)
        new_chore.rotation_group = set_rotation_group(new_chore, workloads)
    
    last_row = write_chore_to_table(new_chore)
    print(f"Chore successfully added with ID {last_row}")
    res = input('(Enter) - Continue to Main Menu \n(x) - Exit program > ')
    if res.lower() == 'x':
        sys.exit()
    else:
        return
    

def view_chores():
    page_size = 5
    offset = 0

    while True:
        cls()
        rows = get_chores_page(page_size, offset)
        next_row = get_chores_page(1, page_size + offset)
        chore_list = [
            [
                chore.chore_id, 
                chore.name, 
                chore.cadence, 
                chore.shared, 
                chore.assignee, 
                chore.rotation_group, 
                chore.time
            ]
            for chore in rows
        ]
        print(tabulate(chore_list, headers = ['ID', 'Chore', 'd/m/w', 'Shared?', 'Assignee', 'Rotation Group', 'est time']))
        if next_row:
            reply = input('\n(Enter) - Show more results \n(q) - Quit to Main Menu\n')
            if reply == 'q':
                break
            offset += 5
            continue
        else:
            reply = input('End of results.\n(s) - Start from beginning\n(q) - Quit to Main Menu\n').lower()
            while True:
                if reply not in ('s', 'q'):
                    reply = input('Invalid selection. Please choose (s) to start from beginning or (q) to quit to Main Menu \n')
                    continue
                else:
                    break
            if reply == 's':
                offset = 0
                continue
            elif reply == 'q':
                break

def balance_chores():
    cls()
    workloads = get_workloads('g')
    res = get_fairness_score(workloads)
    suggest_balance = False
    if res['score'] < 0.85:
        suggest_balance = True
    print(f'Fairness rating: {round(res['score'] * 100, 2)}%')
    print(f'Group A: {res['a']} minutes')
    print(f'Group B: {res['b']} minutes')
    if res['diff'] is not None:
        print(f'Group {'A' if res['a'] > res['b'] else 'B'} currently has {res['diff']} more minutes of chores')
    
    print('---------------')
    print(f'Rebalancing {"is" if suggest_balance else "is not"} suggested at this time.')
    print(f'{"Rebalance" if suggest_balance else "Rebalance anyway"} (Y/N)?\n')
    reply = input('> ')
    while reply.lower() not in ('y', 'n', 'yes', 'no'):
        reply = input('Invalid selection. Please choose (y) or (n) > ')
    if reply.lower() in ('n', 'no'):
        return
    
    chores = get_all_chores()
    deltas = get_new_groups(chores)
    if not deltas:
        print("Chores already balanced. No updates required")
    else:
        res = update_rotation_group(deltas)
        
        if isinstance(res, Exception):
            print(f'Database error: {res}')
        elif res != len(deltas):
            print(f'Warning: Only {res} of {len(deltas)} were updated')
        else:
            updated_workloads = get_workloads()
            updated_score = get_fairness_score(updated_workloads)
            print(f'Chore balancing complete\nNew fairness score: {updated_score}')
    input('Press Enter to return to Main Menu')
        
    
