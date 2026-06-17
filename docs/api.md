# API Documentation

## Purpose
The API represents the transport layer between the HTTP requests the server receives and the app-layer logic. The API/transport layer (as well as the other layers within the application) are generally broken down into four high-level concepts:

- Users
- Templates
- Assignments
- Scheduler

## Endpoints
### Authorization
Current: provide user ID in header:
```X-User-ID : {id}```

### Users
#### POST /users
*Authentication: none*
Create new user
Request: 
```
{
    "username":"{username}"
}
```

Response: `201 Created`
```{
    "ID":"{user_id}",
    "Username":"{username}",
    "Role":"{role}",
    "CreatedAt":"{created timestamp}",
    "UpdatedAt":"{last updated timestamp}"
    }
```

#### GET /users
*Auth: user ID in header*
Get all current users
Response: `200 OK` 
```[
    {
        "ID":"{user_id}",
        "Username":"{username}",
        "Role":"{role}",
        "CreatedAt":"{created timestamp}",
        "UpdatedAt":"{last updated timestamp}"
    },
    {
        "ID":"{user_id}",
        "Username":"{username}",
        "Role":"{role}",
        "CreatedAt":"{created timestamp}",
        "UpdatedAt":"{last updated timestamp}"
    },
    ...
    ]
```

#### GET /users/{id}
*Auth: user ID in header*
Get single user
Response: `200 OK`
```{
    "ID":"{user_id}",
    "Username":"{username}",
    "Role":"{role}",
    "CreatedAt":"{created timestamp}",
    "UpdatedAt":"{last updated timestamp}"
    }
```

#### GET /users/{id}/assignments
*Auth: user ID in header*
View all assignments assigned to a single user
TBD post-implementation

#### PUT /users/{id}
*Auth: user ID in header*
Edit user details. Note: user may only edit self
Request:
```
{
    "username":{new username}
}
```

#### DELETE /users/{id}
*Auth: user ID in header*
Delete user Note: user may only delete self
Response: `204 No Content`

### Templates
#### POST /chore-templates
*Auth: user ID in header*
Add new chore template
Request:
```
{
    "name": "{template name}",
    "cadence": "{daily/weekly/monthly}",
    "shared": {true/false},
    "assignee": "{user id of assignee or empty string}",
    "duration": {time in minutes}
}
```
Response: `201 Created`
```
{
    "ID":"{id}",
    "Name":"{template name}",
    "Cadence":"{cadence}",
    "Shared": {true/false},
    "Assignee":"",
    "Duration":{time},
    "CreatedAt":"{created date}",
    "UpdatedAt":"{last updated date}"
}
```

#### GET /chore-templates
*Auth: user ID in header*
Get all chores
Response: `200 OK`
```
[
    {
        "ID":"{id}",
        "Name":"{template name}",
        "Cadence":"{cadence}",
        "Shared": {true/false},
        "Assignee":"",
        "Duration":{time},
        "CreatedAt":"{created date}",
        "UpdatedAt":"{last updated date}"
    },
    {
        "ID":"{id}",
        "Name":"{template name}",
        "Cadence":"{cadence}",
        "Shared": {true/false},
        "Assignee":"",
        "Duration":{time},
        "CreatedAt":"{created date}",
        "UpdatedAt":"{last updated date}"
    },
    ...
]
```


#### GET /chore-templates/{id}
*Auth: user ID in header*
View a specific chore
Response: `200 OK`
```
{
    "ID":"{id}",
    "Name":"{template name}",
    "Cadence":"{cadence}",
    "Shared": {true/false},
    "Assignee":"",
    "Duration":{time},
    "CreatedAt":"{created date}",
    "UpdatedAt":"{last updated date}"
}
```

#### PUT /chore-templates/{id}
*Auth: user ID in header*
Edit chore. Note: may only be edited by chore owner
Request:
```
{
    "id":"{id}",
    "name":"{updated name}",
    "cadence":"{updated cadence}",
    "shared":{true/false},
    "assignee":"{updated assignee}",
    "duration":{updated duration}
}

```
Response: `200 OK`
```
{
    "ID":"{id}",
    "Name":"{template name}",
    "Cadence":"{cadence}",
    "Shared": {true/false},
    "Assignee":"",
    "Duration":{time},
    "CreatedAt":"{created date}",
    "UpdatedAt":"{last updated date}"
}
```

#### DELETE /chore-templates/{id}
*Auth: user ID in header*
Delete chore. Note: may only be deleted by chore owner
Response: `204 No Content`


#### GET /chore-templates/{id}/assignments
*Auth: user ID in header*
View all scheduled instances of template
TBD after changing to persistence layer

### Assignments
#### POST /assignments
*Auth: user ID in header*
Add assignment. Note: May only assign to self 
Request:
```
{
    "chore_id": "{template id}",
    "assigned_user_id": "{user id}",
    "schedule_date": "{scheduled timestamp}"
}
```
Response: `201 Created`
```
{
    "ID":"{id}",
    "TemplateID":"{template id}",
    "AssignedUserID":"{user id}",
    "ScheduledFor":"{scheduled timestamp}",
    "Completed":{true/false},
    "Canceled":{true/false},
    "CreatedAt":"{Created Timestamp}",
    "UpdatedAt":"{Last Updated Timestamp}",
    "CompletedAt":"{Completed Timestamp}",
    "CanceledAt":"{Canceled Timestamp}"
}
```

#### GET /assignments
*Auth: user ID in header*
View all assignments

Optional Query Parameters:

| Parameter   | Description |
|------------|-------------|
| user_id    | Return assignments assigned to the specified user |
| template_id| Return assignments generated from the specified template |

Examples:
GET /assignments
GET /assignments?user_id={id}
GET /assignments?template_id={id}

Notes
- If no query parameters are supplied, all assignments are returned.
- Only one filter parameter may currently be specified.
- Supplying multiple filters returns a 400 Bad Request.

Response: `200 OK`
```
[
    {
        "ID":"{id}",
        "TemplateID":"{template id}",
        "AssignedUserID":"{user id}",
        "ScheduledFor":"{scheduled timestamp}",
        "Completed":{true/false},
        "Canceled":{true/false},
        "CreatedAt":"{Created Timestamp}",
        "UpdatedAt":"{Last Updated Timestamp}",
        "CompletedAt":"{Completed Timestamp}",
        "CanceledAt":"{Canceled Timestamp}"
    },
    {
        "ID":"{id}",
        "TemplateID":"{template id}",
        "AssignedUserID":"{user id}",
        "ScheduledFor":"{scheduled timestamp}",
        "Completed":{true/false},
        "Canceled":{true/false},
        "CreatedAt":"{Created Timestamp}",
        "UpdatedAt":"{Last Updated Timestamp}",
        "CompletedAt":"{Completed Timestamp}",
        "CanceledAt":"{Canceled Timestamp}"
    }
    ...
]
```

#### GET /assignments/{id}
*Auth: user ID in header*
View a specific assignment
Response: `200 OK`
```
{
    "ID":"{id}",
    "TemplateID":"{template id}",
    "AssignedUserID":"{user id}",
    "ScheduledFor":"{scheduled timestamp}",
    "Completed":{true/false},
    "Canceled":{true/false},
    "CreatedAt":"{Created Timestamp}",
    "UpdatedAt":"{Last Updated Timestamp}",
    "CompletedAt":"{Completed Timestamp}",
    "CanceledAt":"{Canceled Timestamp}"
}
```

#### PUT /assignments/{id}
*Auth: user ID in header*
Edit an assigned chore. Note: may only edit chores assigned to self
Request:
```
{
    "assigned_user_id":"{updated user id}",
    "schedule_date":"{scheduled timestamp}"
}
```
Response: `200 OK`
```
{
    "ID":"{id}",
    "TemplateID":"{template id}",
    "AssignedUserID":"{user id}",
    "ScheduledFor":"{scheduled timestamp}",
    "Completed":{true/false},
    "Canceled":{true/false},
    "CreatedAt":"{Created Timestamp}",
    "UpdatedAt":"{Last Updated Timestamp}",
    "CompletedAt":"{Completed Timestamp}",
    "CanceledAt":"{Canceled Timestamp}"
}
```


#### POST /assignments/{id}/cancel
*Auth: user ID in header*
Cancel an assigned chore. Note: may only delete chores assigned to self
Response: `200 OK`
```
{
    "ID":"{id}",
    "TemplateID":"{template id}",
    "AssignedUserID":"{user id}",
    "ScheduledFor":"{scheduled timestamp}",
    "Completed":{true/false},
    "Canceled":{true},
    "CreatedAt":"{Created Timestamp}",
    "UpdatedAt":"{Last Updated Timestamp}",
    "CompletedAt":"{Completed Timestamp}",
    "CanceledAt":"{Canceled Timestamp}"
}
```


#### POST /assignments/{id}/complete
*Auth: user ID in header*
Complete an assigned chore. Note: may only delete chores assigned to self
Response: `200 OK`
```
{
    "ID":"{id}",
    "TemplateID":"{template id}",
    "AssignedUserID":"{user id}",
    "ScheduledFor":"{scheduled timestamp}",
    "Completed":{true},
    "Canceled":{true/false},
    "CreatedAt":"{Created Timestamp}",
    "UpdatedAt":"{Last Updated Timestamp}",
    "CompletedAt":"{Completed Timestamp}",
    "CanceledAt":"{Canceled Timestamp}"
}
```

### Scheduler
#### POST /scheduler/run
*Auth: user ID in header*
Schedule tasks for current horizon window (see `./docs/scheduler.md`)
Response: `200 OK`
```
"scheduler completed sucessfully"
```

## Future Authentication

Current State
-------------
User identity is supplied by request parameters/headers for development.

Planned State
-------------
Bearer-token authentication.

Future Enhancement
------------------
Role-based authorization:
- User
- Admin