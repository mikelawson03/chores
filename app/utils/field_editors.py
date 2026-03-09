def set_name(chore):
    while True:
        name = input('Chore name > ')
        if name == '':
            print('Invalid name. Please try again.')
            continue
        chore.name = name
        return

def set_cadence(chore):
    while True:
        cadence = input('Cadence (Daily/Weekly/Monthly) > ').lower()
        if cadence not in ['daily','weekly','monthly']:
            print('Invalid choice. Please choose daily, weeekly, or monthly')
            continue
        chore.cadence = cadence
        return

def set_shared(chore):
    shared = False
    while True:
        choice = input('Is this chore shared (Y/N)? ').lower()

        if choice not in ['y', 'n', 'yes', 'no']:
            print('Invalid choice. Please choose either Y or N')
            continue

        if choice in ['y', 'yes']:
            shared = True
        return

def set_assignee(chore):
    while True:
        assignee = input('Who should this chore be assigned to? ').lower()
        if assignee == '':
            print('Invalid assignee. Please try again.')
            continue
        chore.assignee = assignee
        return

def set_time(chore):
    while True:
        try:
            time = int(input('How long in minutes should this chore take? '))
            if time <= 0:
                raise ValueError
            chore.time = time
            return
        except ValueError:
            print('Invalid choice. Please enter a number greater than 0')
