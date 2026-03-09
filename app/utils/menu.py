import os, sys

def cls():
    os.system('cls' if os.name=='nt' else 'clear')

def run_menu(title, options, return_result = False):
    while True:
        if title is not None:
            cls()
            print(title)
            print('-' * len(title))
        for key, (label, _) in options.items():
            print(f'{key.upper()} - {label}')

        choice = input('> ').lower()

        

        if choice not in options:
            print('\nInvalid choice.')
            input('Press Enter to continue')
            continue
    
        action = options[choice][1]
        result = action()

        if return_result:
            return result

        if result == 'back':
            return

