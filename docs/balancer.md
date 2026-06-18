# Balancer

## Balancing Design
The balancer expects a list of assignments and a list of users. These can already be delegated to assignees or they may be unassigned so far. The balancer will then:

1. Get planning horizon window and monthly planning end
2. Retrieve all users and assignments for the horizon (for daily and weekly assignments) or planning window (for monthly assignments) - see Assignment Retrieval below for details
3. Create user loads for all users
4. Add previously assigned assignment durations to user loads
5. Split assignments by cadence
6. Starting with daily cadence, the balancer will:
    a. Sort remaining assignments by duration in descending order
    b. Determine the lowest user load for that cadence
    c. Delegate the first unassigned assignment to the user with the lowest load
    d. Update the user load (note that this is done by passing pointers so that the incrementation happens in-place)
    e. Append new assignment to slice of new assignments for that cadence
    f. Continue assignments in this manner (find lowest load, assign, increment load) until all available assignments have been delegated
    g. Return newDailyAssignments
4. Repeat #6 for weekly and monthly chore assignments
5. Send new assignments to persistence layer
6. Print metrics to the console (future: will output detailed results to log file and return metrics to caller)

## Assignment Retrieval

The balancer does not load all assignments.

Assignments are filtered according to balancing rules:

- Daily assignments must fall within the planning horizon.
- Weekly assignments must fall within the planning horizon.
- Monthly assignments may extend beyond the planning horizon and are included through the end of any month touched by the horizon.
- Canceled assignments are excluded.
- Completed assignments are included because completed work contributes to fairness calculations.

The store layer encapsulates these rules through `GetAssignmentsForBalancing()`.

## Load Calculation
Workload is measured in estimated duration minutes rather than assignment count.

Examples:
- Take out trash: 5 minutes
- Clean bathroom: 45 minutes

Balancing by assignment count would treat these equally and produce unfair distributions. Therefore, all balancing decisions are based on cumulative duration.

## Existing Assignments

The balancer includes already-assigned assignments when calculating current user workloads.

This ensures:
- Fixed assignments affect future balancing decisions.
- Shared chores are distributed around existing commitments.
- Completed assignments within the planning window count toward fairness because the work has already been performed.

Canceled assignments are excluded because no work will occur.

## Assignment Ordering

Assignments are sorted by duration descending before balancing.

This allows larger chores to be assigned first and generally produces more balanced results than assigning chores in arbitrary or chronological order.

Future window-based balancing may introduce additional ordering rules within cadence windows.

## Date Range Conventions
The application uses inclusive-start, exclusive-end date ranges.

Example:
    Start: 2026-07-01
    End:   2026-08-01

Includes:
    2026-07-01 through 2026-07-31

Excludes:
    2026-08-01

This convention is used by:
- Scheduler horizon calculations
- Monthly planning windows
- Assignment retrieval queries
- Balancer assignment scope

## Planning Horizon

The planning horizon defines the period over which scheduling
and balancing occur.

Current implementation:

- Start: Beginning of current week
- End: Four weeks after horizon start

Monthly assignments use an extended planning window because monthly chores are scheduled at month-end. This ensures that monthly assignments created by the scheduler are visible to the balancer even when their scheduled date falls outside the standard horizon.

## Why does balanceUserLoads return AssignmentWithMetadata?

The balancer requires duration and cadence information while assigning chores. Although persistence only needs the underlying Assignment struct, metadata is preserved after balancing so that callers can:

- Generate assignment statistics
- Produce logs
- Calculate workload metrics
- Support future analytics features

The balancer then extracts the embedded Assignment(s) to a new slice of just Assignment types to send to the persistence layer

## Bulk Assignment Update - Persistence Layer
We update all of the new assignments in bulk after balancing. We do this in a single transaction. If all of the transactions commit successfully, we currently return 200 OK response with a body that says "balancer run successful." If any fail, we roll back the updates and return an 500 Internal Server Error along with the transaction error.

## Future Enhancements
- Window-based balancing - the app currently only balances across the horizon window. After a few months, we will revisit and decide if balancing within a cadence window (e.g., one week at a time for weekly cadence rather than all weekly assignments in a horizon window) would increase short-term fairness
- Status and history endpoints - ability to look at current status of balancer (e.g., last run, # of unassigned tasks, etc) as well as history of last X runs (e.g., successful run, # assigned for each cadence, # skipped for each cadence, etc)