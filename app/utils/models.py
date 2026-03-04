class Chore:
    def __init__(self, name, cadence, shared, assignee, time, rotation_group = None, chore_id = None):
        self.chore_id = chore_id
        self.name = name
        self.cadence = cadence
        self.shared = shared
        self.assignee = assignee
        self.time = time
        self.rotation_group = rotation_group

    def __repr__(self):
        return(f"""
        Chore ID: {self.chore_id}
        Chore name: {self.name}
        Cadence: {self.cadence}
        Shared? {self.shared}
        Assignee: {self.assignee if self.assignee is not None else 'N/A'}
        Rotation Group: {self.rotation_group if self.rotation_group is not None else 'N/A'}
        Time to complete: {self.time}
        """)    

class Workload:
    def __init__(self, rotation_group, time):
        self.rotation_group = rotation_group
        self.time = time
