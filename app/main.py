import time, os, sys
from utils.chores import add_chore, view_chores
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
    print('4. View chores')
    print('5. Exit')
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
    while True:
        cls()
        choice = menu()
        if choice == 1:
            add_chore()
        elif choice == 2:
            continue
            # updated_chore = edit_chore()
            # print(f'---------------\nChore updated: {updated_chore}')
        elif choice == 3:
            cls()
            continue
        elif choice == 4:
            view_chores()
            continue
        elif choice == 5:
            sys.exit()
        else:
            cls()
            print(f'\nInvalid choice. Please make a valid selection\n\n')
    

        


if __name__ == '__main__':
    main()