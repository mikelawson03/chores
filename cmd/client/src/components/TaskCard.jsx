import { Card } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { Checkbox } from "@mui/material";

export default function TaskCard({ task, openTaskDetails, toggleTaskComplete }) {
  return (
    <Card 
      onClick={() => openTaskDetails(task)}
      sx= {{ 
      width: "100%",
      borderRadius: 2,
      p: 2,
      backgroundColor: task.completed ? "grey.200" : "background.paper"
      }}>
      <Typography variant="h6" 
      gutterBottom 
      sx = {{
        textDecoration: task.completed ? "line-through" : "none", 
        color: task.completed ? "text.secondary" : "text.primary"
        }}>
        {task.title}
      </Typography>
      <Typography variant="body2" sx = {{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none"}}>
        Assigned to: {task.assignee}
      </Typography>
      <Stack 
        direction="row" 
        sx = {{ width: "100%", justifyContent:"space-between"  }}
      >
        <Typography variant="body2" sx={{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none" }}>
          {task.duration}
        </Typography>
        <Checkbox 
          checked={task.completed} 
          sx={{ p: 0 }} 
          onChange={() => {toggleTaskComplete(task);}}
          onClick={(event) => {event.stopPropagation();}}
        />
      </Stack>
    </Card>
    );
}