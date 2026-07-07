import { Card, CardContent } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { Checkbox } from "@mui/material";

export default function PlannerTaskCard({ task, onToggleComplete, width, openTaskDetails }) {
  return (
    <Card onClick={() => openTaskDetails(task)}
    sx= {{ 
      width: width,
      borderRadius: 1,
      backgroundColor: task.completed ? "grey.200" : "background.paper"
      }}>
      <CardContent sx={{ p: 0.5, "&:last-child": { pb: 0.5 }, lineHeight: 1, }}>
        <Typography variant="body1"  
        sx = {{
          textDecoration: task.completed ? "line-through" : "none", 
          color: task.completed ? "text.secondary" : "text.primary"
          }}>
          {task.title}
        </Typography>
        <Typography variant="body2" sx = {{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none", mb: 0.25, lineHeight: 1.15}}>
          {task.assignee}
        </Typography>
        <Stack 
          direction="row" 
          sx = {{ width: "100%", justifyContent:"space-between"  }}
        >
          <Typography variant="body2" sx={{ color: 'text.secondary', textDecoration: task.completed ? "line-through" : "none" }}>
            {task.duration}
          </Typography>
          <Checkbox checked={task.completed} size="small" sx={{ p: 0 }} onChange={() => {onToggleComplete(task.id);}} />
        </Stack>
      </CardContent>
    </Card>
    );
}