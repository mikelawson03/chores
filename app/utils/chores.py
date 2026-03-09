
from utils.sql_handlers import write_chore_to_table, get_chores_page, get_all_chores, get_workloads, update_rotation_group, query_chores, update_chore, delete_chore
from utils.balancer import get_new_groups, set_rotation_group, get_fairness_score
from utils.models import Chore
from utils.menu import run_menu
from utils.field_editors import set_name, set_cadence, set_shared, set_assignee, set_time
from tabulate import tabulate
import os, sys
import copy

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def search_by_name():
    cls()
    print('Search Chores')
    print('---------------')
    print('Please enter the name of the chore you\'d like to find\n')
    choice = input('> ').lower()

    rows = query_chores(choice)
    return (choice, rows)


def display_chores(chore_list):
    l = [
        [
            chore.name, 
            chore.chore_id, 
            chore.cadence, 
            chore.shared, 
            chore.assignee, 
            chore.rotation_group, 
            chore.time
        ]
        for chore in chore_list
    ]
    if len(l) == 0:
        input('No results found. Press Enter to continue')
        return
    print(tabulate(l, headers = ['Row #', 'Chore', 'ID', 'd/m/w', 'Shared?', 'Assignee', 'Rotation Group', 'est time'], showindex=range(1, len(chore_list)+1)))
    print(f'\n{len(l)} result{'s' if len(l) > 1 else ''} found')
    



def add_chore():
    assignee = None
    time = 0

    cls()    
    print('Add chore')
    print('-------------- \n')
    new_chore = Chore()
    set_name(new_chore)
    set_cadence(new_chore)   
    set_time(new_chore)
    
    #Will implement shared chores later
    # set_shared(new_chore)

    # if not new_chore.shared :
    #     set_assignee(new_chore)        
    
    # assign to balanced group 
    workloads = get_workloads(new_chore.cadence)
    new_chore.rotation_group = set_rotation_group(new_chore, workloads)

    last_row = write_chore_to_table(new_chore)
    print(f'\nChore successfully added with ID {last_row}')
    input('Press Enter to continue')
    

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
            reply = input('\n(Enter) - Show more results \n(q) - Return to previous menu\n')
            if reply == 'q':
                break
            offset += 5
            continue
        else:
            reply = input('End of results.\n(s) - Start from beginning\n(q) - Return to previous menu\n').lower()
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
            cls()
            print(f'Chore balancing complete\nNew fairness score: {updated_score['score'] * 100}%')
    input('Press Enter to return to Main Menu')




def edit_delete_chore():
    res = None
    options = {
        's': ('Search for chore', search_by_name),
        'v': ('View chores', view_chores),
        'q': ('Return to previous menu', lambda: 'back'),
        'x': ('Exit program', sys.exit)
    }

    
    while True:
        res = run_menu('Edit or delete chore', options, True)
        if res =='back':
            return
    
        cls()
        print(f'\nResults for "{res[0]}"')
        print('---------------\n')
        chore_list = res[1]
        display_chores(chore_list)

        if len(chore_list) > 1:
            options = {
                str(i): (f"Edit chore #{i}", lambda i=i: chore_list[i-1])
                for i in range(1, len(chore_list) + 1)
            }
            options['q'] = ('Return to previous menu', lambda: 'back')
            options['x'] = ('Exit program', sys.exit)
            chore = run_menu(None, options, True)
        else:
            chore = chore_list[0]
        if chore == 'back':
            break
        
        working_chore = copy.copy(chore)

        options = {
            '1': ('Edit name', lambda: set_name(working_chore)),
            '2': ('Edit cadence', lambda: set_cadence(working_chore)),
            '3': ('Edit estimated time', lambda: set_time(working_chore)),
            's': ('Save updates', lambda: 'save'),
            'd': ('Delete chore', lambda: 'delete'),
            'q': ('Return to previous menu', lambda: 'back'),
            'x': ('Quit program', sys.exit)
        }

        while True:
            res = run_menu('Edit or delete chore', options, True)
            
            if res == 'back':
                break
            if res == 'save':
                res = update_chore(working_chore)
                if isinstance(res, Exception):
                    print(f'Database error: {res}')
                elif res == 1:
                    print(f'Chore updated successfully.')
                else:
                    print(f'Chore failed to update')
                print('Press Enter to continue')
                input()
                return
            
            if res == 'delete':
                res = delete_chore(working_chore)
                if isinstance(res, Exception):
                    print(f'Database error: {res}')
                elif res == 1:
                    print(f'Chore deleted successfully.')
                else:
                    print(f'Chore failed to delete')
                print('Press Enter to continue')
                input()
                return