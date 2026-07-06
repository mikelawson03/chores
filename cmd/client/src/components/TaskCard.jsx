import { Card } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { Checkbox } from "@mui/material";

export default function TaskCard({ id, title, assignee, duration, completed, onToggleComplete }) {
  return (
    <Card sx= {{ 
      width: "100%",
      borderRadius: 2,
      p: 2,
      backgroundColor: completed ? "grey.200" : "background.paper"
      }}>
      <Typography variant="h6" 
      gutterBottom 
      sx = {{
        textDecoration: completed ? "line-through" : "none", 
        color: completed ? "text.secondary" : "text.primary"
        }}>
        {title}
      </Typography>
      <Typography variant="body2" sx = {{ color: 'text.secondary', textDecoration: completed ? "line-through" : "none"}}>
        Assigned to: {assignee}
      </Typography>
      <Stack 
        direction="row" 
        sx = {{ width: "100%", justifyContent:"space-between"  }}
      >
        <Typography variant="body2" sx={{ color: 'text.secondary', textDecoration: completed ? "line-through" : "none" }}>
          {duration}
        </Typography>
        <Checkbox checked={completed} sx={{ p: 0 }} onChange={() => {onToggleComplete(id);}} />
      </Stack>
    </Card>
    );
}