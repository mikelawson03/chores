## Overarching Principles

### Design Flow

Overview → Details → Action

Every module should follow it.

#### Dashboard
- Overview of work
- Click to Assignment Details
- Complete/Trade/etc.

#### Chores
- List
- Details Modal
- Edit/Create/Delete

#### Calendar
- Month View
- Day Agenda
- Assignment Details

#### Admin
- Run History
- Run Details
- Logs

## Dashboard

### Dashboard Design Principle

The dashboard should not merely show future work.

The dashboard should show future work that helps the user plan.

Routine chores (e.g., daily chores) are already visible through the user's current assignments and calendar views. Dashboard widgets should prioritize surfacing work that may require additional planning, awareness, or preparation.

Examples include:
- Weekly chores
- Monthly chores
- One-time chores
- Upcoming assignments with unusually high effort

The goal of the dashboard is to provide actionable visibility, not simply repeat information already available elsewhere in the application.

### Dashboard Goals

The dashboard should provide actionable visibility.

The dashboard should not merely repeat information available elsewhere in the application.

Dashboard widgets should prioritize surfacing:
- Current responsibilities
- Upcoming planning needs
- Recent accomplishments

The dashboard should answer:
"What should I know right now?"

## Templates Screen

Initial data loaded via:
- GET /chore-templates

Sorting and filtering are performed client-side.

Rationale:
- Expected dataset size is small.
- Sorting/filtering are presentation concerns.
- Client-side operations provide a more responsive UX.

### Filtering Philosophy

For small datasets (e.g., chore templates), filtering and sorting are performed client-side.

Filters should reflect common user workflows rather than expose generic query-building functionality.

Initial filters:
- Frequency
- Assignee

Future filters:
- Duration
- Updated Date


## Assignments Calendar

### Calendar Design Principle

Calendar is a single module containing multiple representations of assignment data.

Different views answer different user questions while operating on the same underlying assignments.

Users may switch between views without leaving the Calendar section.

### Assignment Status (UI)

The UI presents four assignment states:

- Overdue
- Scheduled
- Completed
- Canceled

"Overdue" is a derived state. An assignment is considered overdue when it is scheduled, incomplete, and its scheduled date is before today.

This is a presentation concern and does not require a separate database status.

### Frontend vs. Backend Scheduling

The backend stores recurring assignments on their due date (e.g., weekly chores are due on the last day of the week). This simplifies scheduling, balancing, and persistence.

The Planning View intentionally abstracts this implementation detail. Instead of presenting recurring chores on their due date, it groups unscheduled weekly and monthly work into planning sections ("Complete this week" / "Complete before month end"). This reflects how users naturally think about flexible work rather than how it is stored internally.

## Weekly Planner – Logical Buckets

The Weekly Planner organizes work by planning state, not by how it's stored in the backend or by recurrence.

### Scheduled Work

Displayed in the weekly calendar grid.

Represents tasks that have already been committed to a specific day.

"I know when I'm doing this."

### Weekly Work

Displayed in the Weekly Work section.

Contains flexible weekly tasks that still need to be planned before the end of the current week.

"I need to decide when to do this this week."

### Monthly Work

Displayed in the Monthly Work section.

Contains flexible monthly tasks that still need to be planned before the end of the current month.

"I need to fit this in sometime this month."

### Overdue Work (future enhancement)

Displayed in an Overdue Work section.

Contains any incomplete tasks whose due date has already passed, regardless of whether they originated as daily, weekly, monthly, or one-off work.

"I missed my target and should deal with this."

Examples:

A weekly task not completed by Sunday appears here on Monday.
A monthly task due June 30 appears here on July 1 until completed.
A one-off task moves here the day after its due date.

This prevents overdue work from being mixed into the current week's or month's planning backlog.

## Chore Details

UI components should be designed independently of their presentation container.

ChoreDetails and ChoreForm may initially be displayed within modals, but should be structured so they can
later be promoted to dedicated pages if the feature outgrows the modal workflow.


### Validation Rules (Chore Creation / Edit)

Chore Name
- Required

Frequency
- Required

Duration
- Required
- Hours and Minutes may not both equal 0
- Minutes must be 0-59

Description
- Optional

Rotate Assignment Among Household Members
- Defaults to checked

Assignee
- Hidden when rotation enabled
- Required when rotation disabled


## Admin Actions

## Duration Display

The API returns durations in minutes.

The client is responsible for converting durations
into human-readable formats appropriate for the UI.

Examples:

10  -> 10 min
60  -> 1 hr
90  -> 1 hr 30 min
150 -> 2 hr 30 min

## Date Display

Dates may be displayed in either relative or absolute form.

Recent dates:
- Today
- Yesterday
- X days ago

Older dates:
- Jun 1
- Mar 27

The goal is to optimize for readability and quick recognition rather than strict timestamp accuracy.

Date Rules:
- 0 days      -> Today
- 1 day       -> Yesterday
- 2-6 days    -> X days ago
- 7-364 days  -> Jun 10
- 365+ days   -> Jun 10, 2025

## Navigation

Current MVP navigation is intentionally flat.

Future versions may organize functionality into domain-based sections (e.g., Chore Management, Allowance, Family Calendar).

Current screens should be designed so they can be moved into a domain section without significant redesign.

