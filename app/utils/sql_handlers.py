import sqlite3

db = 'chores.db'

def create_table():
    conn = sqlite3.connect(db)
    c = conn.cursor()
    c.execute('''
        CREATE TABLE IF NOT EXISTS chores (
        ID INTEGER PRIMARY KEY ASC,
        CHORE TEXT,
        CADENCE TEXT,
        SHARED INTEGER,
        ASSIGNEE TEXT,
        TIME INT)
    ''')
    conn.commit()
    conn.close()

def write_chore_to_table(chore):
    conn = sqlite3.connect(db)
    c = conn.cursor()
    c.execute(
        "INSERT INTO chores (CHORE, CADENCE, SHARED, ASSIGNEE, TIME) VALUES (?, ?, ?, ?, ?)",
        (chore.name, chore.cadence, chore.shared, chore.assignee, chore.mins)
        )
    conn.commit()
    conn.close()
    return