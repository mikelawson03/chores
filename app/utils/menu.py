from utils.chores import add_chore, view_chores, balance_chores, update_chore
import os, sys

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def menu():
    print('Main Menu')
    print('-----------------')
    print('What would you like to do?')
    print('1. Add a chore')
    print('2. Balance chores')
    print('3. Edit a chore')
    print('4. Remove a chore')
    print('5. View chores')
    print('-----------------\n')
    print('6. Exit')
    choice = input('> ')
    return choice

def menu_handler(choice):
    if choice == '1':
        add_chore()
    elif choice == '2':
        balance_chores()
    elif choice == '3':
        update_chore()
    elif choice == '4':
        cls()
        return
    elif choice == '5':
        view_chores()
        return
    elif choice == '6':
        sys.exit()
    else:
        cls()
        print(f'\nInvalid choice. Please make a valid selection')