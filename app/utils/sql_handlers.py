import sqlite3
from utils.models import Chore, Workload

db = 'chores.db'

def create_table():
    with sqlite3.connect(db) as conn:
        c = conn.cursor()
        c.execute('''
            CREATE TABLE IF NOT EXISTS chores (
            CHORE_ID INTEGER PRIMARY KEY ASC,
            CHORE TEXT UNIQUE,
            CADENCE TEXT,
            SHARED INTEGER,
            ASSIGNEE TEXT,
            ROTATION_GROUP TEXT,
            TIME INT
            )
        ''')
        
def row_to_chore(row):
    chore = Chore(
        row['CHORE'],
        row['CADENCE'],
        row['SHARED'],
        row['ASSIGNEE'],
        row['TIME'],
        row['ROTATION_GROUP'],
        row['ID']
        )
    return chore

def row_to_workload(row):
    workload = Workload(
        row['ROTATION_GROUP'],
        row['sum(TIME)']
    )
    return workload

def write_chore_to_table(chore):
    try:
        with sqlite3.connect(db) as conn:
            c = conn.cursor()
            c.execute(
                'INSERT INTO chores (CHORE, CADENCE, SHARED, ASSIGNEE, ROTATION_GROUP, TIME) VALUES (?, ?, ?, ?, ?, ?)',
                (chore.name, chore.cadence, chore.shared, chore.assignee, chore.rotation_group, chore.time)
                )
            
            return c.lastrowid
    except sqlite3.Error as e:
        print(f"Database error: {e}")

def update_rotation_group(deltas):
    print(deltas)
    try:
        with sqlite3.connect(db) as conn:
            c = conn.cursor()
            c.executemany('''
                UPDATE chores
                SET ROTATION_GROUP = ?
                WHERE ID = ?
                ''',
                deltas
                )
            return c.rowcount
    except sqlite3.Error as e:
        return e

def get_all_chores():
    with sqlite3.connect(db) as conn:
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        rows = c.execute('SELECT * FROM chores ORDER BY ID'
        ).fetchall()
        return [row_to_chore(row) for row in rows]

def get_chores_page(page_size, offset):
    with sqlite3.connect(db) as conn:
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        rows = c.execute(
            'SELECT * FROM chores ORDER BY ID LIMIT ? OFFSET ?',
            (page_size, offset)
        ).fetchall()
        return [row_to_chore(row) for row in rows]

def get_workloads(cadence = 'g'):
    with sqlite3.connect(db) as conn:
        conn.row_factory = sqlite3.Row
        c = conn.cursor()
        if cadence == 'g':
            rows = c.execute('SELECT ROTATION_GROUP, sum(TIME) FROM chores GROUP BY ROTATION_GROUP').fetchall()
        else:
            rows = c.execute('SELECT ROTATION_GROUP, CADENCE, sum(TIME) FROM chores WHERE CADENCE = ? GROUP BY ROTATION_GROUP', cadence).fetchall()
        return [row_to_workload(row) for row in rows]

