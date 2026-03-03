class Chore:
    def __init__(self, name, cadence, shared, assignee, mins, rotation_group = None):
        self.name = name
        self.cadence = cadence
        self.shared = shared
        self.assignee = assignee
        self.mins = mins
        self.rotation_group = rotation_group

    def __repr__(self):
        return(f"""
        Chore name: {self.name}
        Cadence: {self.cadence}
        Shared? {self.shared}
        Assignee: {self.assignee if self.assignee is not None else 'N/A'}
        Time to complete: {self.mins}
        """)    