from utils.sql_handlers import write_chore_to_table

class Chore:
    def __init__(self, name, cadence, shared, assignee, mins):
        self.name = name
        self.cadence = cadence
        self.shared = shared
        self.assignee = assignee
        self.mins = mins

    def __repr__(self):
        return(f"""
        Chore name: {self.name}
        Cadence: {self.cadence}
        Shared? {self.shared}
        Assignee: {self.assignee if self.assignee is not None else 'N/A'}
        Time to complete: {self.mins}
        """)

def add_chore():
    assignee = None
    mins = 0

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
    elif cdence == 'monthly':
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
    write_chore_to_table(new_chore)
    return new_chore