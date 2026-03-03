import sqlite3

db = 'chores.db'

def create_table():
    conn = sqlite3.connect(db)
    c = conn.cursor()
    c.execute('''
        CREATE TABLE IF NOT EXISTS chores (
        ID INTEGER PRIMARY KEY ASC,
        CHORE TEXT UNIQUE,
        CADENCE TEXT,
        SHARED INTEGER,
        ASSIGNEE TEXT,
        ROTATION_GROUP TEXT,
        TIME INT
        )
    ''')
    conn.commit()
    conn.close()

def write_chore_to_table(chore):
    try:
        conn = sqlite3.connect(db)
        c = conn.cursor()
        c.execute(
            'INSERT INTO chores (CHORE, CADENCE, SHARED, ASSIGNEE, ROTATION_GROUP, TIME) VALUES (?, ?, ?, ?, ?, ?)',
            (chore.name, chore.cadence, chore.shared, chore.assignee, chore.rotation_group, chore.mins)
            )
        conn.commit()
        conn.close()
        return c.lastrowid
    except sqlite3.Error as e:
        conn.rollback()
        print(f"Database error: {e}")
        conn.close()


def get_chores(page_size, offset):
    conn = sqlite3.connect(db)
    c = conn.cursor()
    rows = c.execute(
        'SELECT * FROM chores ORDER BY ID LIMIT ? OFFSET ?',
        (page_size, offset)
    ).fetchall()
    conn.close()
    return rows

def get_workloads(cadence = 'g'):
    conn = sqlite3.connect(db)
    c = conn.cursor()
    if cadence == 'g':
        rows = c.execute('SELECT ROTATION_GROUP, ASSIGNEE, TIME FROM chores GROUP BY ROTATION_GROUP, ASSIGNEE').fetchall()
    else:
        rows = c.execute('SELECT ROTATION_GROUP, ASSIGNEE, TIME FROM chores WHERE CADENCE = ? GROUP BY ROTATION_GROUP, ASSIGNEE', cadence).fetchall()
    conn.close()
    return rows
    
