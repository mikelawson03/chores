# Balancer

## Balancing Design
The balancer expects a list of assignments and a list of users. These can already be delegated to assignees or they may be unassigned so far. The balancer will then:

1. Split assignments by cadence
2. Starting with daily cadence, the balancer will:
    a. Create and maintain a running total of the number of chore minutes assigned to each user
    b. Add previously delegated assignments to the total
    c. Sort remaining assignments by duration in descending order
    c. Use a "greedy" algorithm to delegate the remaining assignments to users where the largest available assignment goes to the user with the lowest total of chore minutes
    d. Continue assignments in this manner until all available assignments have been delegated
3. Repeat #2 for weekly and monthly chore assignments