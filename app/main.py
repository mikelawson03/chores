import time, os, sys
from utils.menu import menu, menu_handler

from utils.sql_handlers import create_table

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def week_type():
    week = int(time.strftime('%U'))
    week_type = ''

    if week % 2 == 0:
        week_type = 'even'

    else:
        week_type = 'odd'
    
    print(week_type)

def main():
    if not os.path.exists('chores.db'):
        create_table()
    while True:
        cls()
        choice = menu()
        menu_handler(choice)
    
if __name__ == '__main__':
    main()