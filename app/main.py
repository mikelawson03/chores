import time
import os
from utils.chores import add_chore
from utils.sql_handlers import create_table

def cls():
    os.system('cls' if os.name=='nt' else 'clear')
    

def menu():
    print('Main Menu')
    print('-----------------')
    print('What would you like to do?')
    print('1. Add a chore')
    print('2. Edit a chore')
    print('3. Remove a chore')
    print('4. View chores\n')
    choice = input('> ')
    return int(choice)

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
    choice = menu()
    print(type(choice))
    if choice == 1:
        cls()
        new_chore = add_chore()
        print(f'''---------------\nNew chore added: {new_chore}''')

        


if __name__ == '__main__':
    main()