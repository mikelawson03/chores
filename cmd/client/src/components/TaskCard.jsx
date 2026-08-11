import { Card } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { Checkbox } from "@mui/material";
import { clickableSurface } from "../styles/surfaces";
import { formatDuration } from "../utils/formatters";

export default function TaskCard({ task, openTaskDetails, toggleTaskComplete }) {
  return (
    <Card 
      onClick={() => openTaskDetails(task)}
      sx= {[clickableSurface, { 
      width: "100%",
      borderRadius: 2,
      p: 2,
      backgroundColor: task.completed ? "grey.200" : "background.paper"
      }]}>
      <Typography variant="h6" 
      gutterBottom 
      sx = {{
        textDecoration: task.completed ? "line-through" : "none", 
        color: task.completed ? "text.secondary" : "text.primary"
        }}>
        {task.templateName}
      </Typography>
      <Typography variant="body2" sx = {{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none"}}>
        {task.userFirstName ? `Assigned to: ${task.userFirstName}` : "Unassigned"}
      </Typography>
      <Stack 
        direction="row" 
        sx = {{ width: "100%", justifyContent:"space-between"  }}
      >
        <Typography variant="body2" sx={{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none" }}>
          {formatDuration(task.duration)}
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