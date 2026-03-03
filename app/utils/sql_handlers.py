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
        TIME INT)
    ''')
    conn.commit()
    conn.close()

def write_chore_to_table(chore):
    try:
        conn = sqlite3.connect(db)
        c = conn.cursor()
        c.execute(
            "INSERT INTO chores (CHORE, CADENCE, SHARED, ASSIGNEE, TIME) VALUES (?, ?, ?, ?, ?)",
            (chore.name, chore.cadence, chore.shared, chore.assignee, chore.mins)
            )
        conn.commit()
        print(f"Chore successfully added with ID {c.lastrowid}")
    except sqlite3.Error as e:
        conn.rollback()
        print(f"Database error: {e}")
    finally:    
        conn.close()


def get_chores(page_size, offset):
    conn = sqlite3.connect(db)
    c = conn.cursor()
    rows = c.execute(
        "SELECT * FROM chores ORDER BY ID LIMIT ? OFFSET ?",
        (page_size, offset)
    ).fetchall()
    conn.close()
    return rows
    
