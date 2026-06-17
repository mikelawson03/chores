# Scheduler

## Definitions
**Chore Templates** - Definitions of recurring chores created by users (e.g., "take out the trash every week," "wash the dishes every day")
**Assignments** - Scheduled instances of chore templates; these say "the chore described by the template is scheduled to be done on a specific date"
**Scheduling Cadence** - The frequency an assignment repeats (daily, weekly, monthly)
**Scheduling / Planning Horizon** - The time window the scheduler uses to plan assignments (currently 28 days)
**Assignment Lookup Window** - The time window of assignments that the scheduler retrieves for ensuring duplicate assignments are not made

## Scheduling Assignments
When the scheduler runs, it will look at the scheduling cadence within the scheduler horizon (see "Scheduler Horizon and Assignment Lookup Window" below). The scheduler will then:

1. Retrieve users, templates, and existing assignments from the SQLite database
2. Run the daily, weekly, and monthly schedulers in that order.
3. For each cadence, the scheduler will find the templates that meet that cadence and check to see if an assignment has been made for each cadence window within the horizon (e.g., for daily, the scheduler will loop through all days in the scheduling horizon and check that a task with the appropriate chore template ID has been scheduled on that day)
4. If assignments are missing for a specified cadence window, the scheduler will create an assignment and schedule it for the *last day of the cadence window* (e.g, for weekly, this means the last day of the week; for monthly, this means the last day of the month)
5. New Assignments are added to a list that is passed along to the balancer for delegating to users

## Canceled and Completed Assignments
For the purpose of the scheduler, assignments that have been marked as canceled or completed are still considered to have been scheduled on their `scheduled_for` date. These assignments will not be recreated by the scheduler.

## Assignment Ownership Rules
Assignment creation happens when the scheduler is run, but they are only delegated to an assignee if the Chore Template specifies one. Otherwise, the scheduler sets the assignee to "" and the assignee is delegated during balancing (see `docs/balancing.md`)

## Assignments for Shared vs Assigned Templates
The scheduler ignores the concept of a shared template. It only looks at whether or not an assignee is included. Rules for shared and assigned templates are below. Note that validity is enforced as time of template creation:

| Shared | Assignee  | Valid? | Meaning |
|--------|-----------|--------|---------|
| true   | ""        | ✅     | Shared chore. Scheduler/balancer can assign it to anyone. |
| true   | "user-id" | ✅     | Assigned chore. This user is the intended owner. Scheduler should generate assignments for that user. |
| false  | ""        | ❌     | Invalid. Chore is neither shared nor assigned, so nobody can receive it. |
| false  | "user-id" | ✅*    | Treated the same as assigned. The assignee takes precedence. |

## Scheduler Horizon and Assignment Lookup Window

### Scheduler Planning Horizon
The scheduler uses a configurable planning horizon (currently 28 days) to determine which cadence windows require assignment generation.

**Examples:**
Horizon Start: June 15
Horizon End:   July 13

The scheduler evaluates all cadence windows touched by the horizon:
Daily:
   June 15 - July 12

Weekly:
   All weeks overlapping June 15 - July 13

Monthly:
   June and July

### Monthly Scheduling Rule
Monthly chores use **month coverage** rather than strict **horizon coverage**. If the planning horizon touches any portion of a month, the scheduler will generate an assignment for that month if one does not already exist.

**Example:**
Horizon:
   June 15 - July 13

Monthly chore:
   Clean gutters

Assignments generated:
June 30
July 31

Note that July 31 > Horizon End. This is intentional. The scheduler guarantees that every month touched by the horizon contains a monthly assignment.

### Assignment Lookup Window
Because monthly assignments may be generated beyond the planning horizon, assignment retrieval cannot stop at the horizon end date.

Instead:
Scheduler loads assignments through the first day of the month following the month containing Horizon End.

The lookup window uses:
Start = inclusive
End   = exclusive

**Example:**
Horizon End:
   July 13

Assignment Lookup End:
   August 1

Assignments loaded:
July 1 - July 31

This ensures that monthly assignments generated outside the planning horizon remain visible to future scheduler runs and prevents duplicate monthly assignment creation.

### Design Principle
The scheduler must always be able to retrieve every assignment it is capable of generating.

If assignment generation extends beyond the planning horizon, the lookup window must be extended accordingly.

## Planned Future Enhancements
- Enable multiple occurrences within cadence window (e.g., 3X/week or 2X/month)
- Create interval scheduling (e.g., every 3 days, every other week, quarterly)
- Separate template owner from assignee
    - Currently these are both managed under assignee
    - There are likely cases where an owner might assign templates to another user (e.g., a parent assigning chores to a child, admins managing templates)
    - These will likely depend on developing a permissions system and creating the concept of a "household"
- Separating assignment due date from schedule date
    - Due date - scheduler concern that communicates the latest possible date an assignment should be completed
    - Scheduled date - user planning concern that communicates date user prefers to complete assignment
    - These are currently collapsed under scheduled_for, but separation would provide a better UX
- Scheduler refactor
    - Currently, each cadence has its own helper function that manages scheduling for that specific cadence
    - Much of the code is repeated and follows similar logic
    - This could be refactored into a single helper that accepts a domain.Cadence value and manages both the broad shared logic as well as the cadence-specific logic based on the value provided to the function
