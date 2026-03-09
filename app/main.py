import time, os, sys
from utils.chores import add_chore, view_chores, balance_chores, edit_delete_chore
from utils.menu import run_menu

from utils.sql_handlers import create_table

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def main_menu():
    options = {
        '1': ('Add chore', add_chore),
        '2': ('Edit or Delete chore', edit_delete_chore),
        '3': ('Balance chores', balance_chores),
        '4': ('View chores', view_chores),
        'x': ('Exit program', sys.exit)
    }

    run_menu('Main Menu', options)

def main():
    if not os.path.exists('chores.db'):
        create_table()
    main_menu()

if __name__ == '__main__':
    main()