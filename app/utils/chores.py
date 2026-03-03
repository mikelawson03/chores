from utils.sql_handlers import write_chore_to_table, get_chores, get_workloads
from utils.balancer import balance_load, set_rotation_group
from utils.models import Chore
from tabulate import tabulate
import os, sys

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def add_chore():
    assignee = None
    mins = 0

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

    while isinstance(mins, int) and mins <= 0:
        try:
            mins = int(input('How long in minutes should this chore take? '))
        except:
            print('Invalid choice. Please enter a number greater than 0')

    new_chore = Chore(name, cadence, shared, assignee, mins)
    
    # assign to balanced group 
    if shared:
        workloads = get_workloads(cadence)
        set_rotation_group(new_chore, workloads)
    
    last_row = write_chore_to_table(new_chore)
    print(f"Chore successfully added with ID {last_row}")
    res = input('(Enter) - Continue to Main Menu \n(q) - Exit program > ')
    if res.lower() == 'q':
        sys.exit()
    else:
        return
    

def view_chores():
    page_size = 5
    offset = 0

    while True:
        cls()
        rows = get_chores(page_size, offset)
        next_row = get_chores(1, page_size + offset)
        print(tabulate(rows, headers = ['ID', 'Chore', 'd/m/w', 'Shared?', 'Assignee', 'Rotation Group', 'est time']))
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
