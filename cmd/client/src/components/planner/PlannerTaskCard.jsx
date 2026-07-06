import { Card, CardContent } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { Checkbox } from "@mui/material";

export default function PlannerTaskCard({ id, title, assignee, duration, completed, onToggleComplete, width }) {
  return (
    <Card sx= {{ 
      width: {width},
      borderRadius: 1,
      backgroundColor: completed ? "grey.200" : "background.paper"
      }}>
      <CardContent sx={{ p: 0.5, "&:last-child": { pb: 0.5 }, lineHeight: 1, }}>
        <Typography variant="body1"  
        sx = {{
          textDecoration: completed ? "line-through" : "none", 
          color: completed ? "text.secondary" : "text.primary"
          }}>
          {title}
        </Typography>
        <Typography variant="body2" sx = {{ color: 'text.secondary', textDecoration: completed ? "line-through" : "none", mb: 0.25, lineHeight: 1.15}}>
          {assignee}
        </Typography>
        <Stack 
          direction="row" 
          sx = {{ width: "100%", justifyContent:"space-between"  }}
        >
          <Typography variant="body2" sx={{ color: 'text.secondary', textDecoration: completed ? "line-through" : "none" }}>
            {duration}
          </Typography>
          <Checkbox checked={completed} size="small" sx={{ p: 0 }} onChange={() => {onToggleComplete(id);}} />
        </Stack>
      </CardContent>
    </Card>
    );
}