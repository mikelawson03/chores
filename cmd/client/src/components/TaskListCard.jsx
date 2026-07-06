import { Card, Checkbox, List, ListItem, ListItemIcon, ListItemText, Typography } from "@mui/material";


export default function TaskListCard({ cardName, tasks, maxItems, footerText, onToggle}) {
  return (
    <Card sx={{
        width: "100%",
        borderRadius: 2,
        p: 2,
      }}>
      <Typography 
      variant="h6" 
      gutterBottom>
        {cardName}
      </Typography>
      <List>
        {tasks
          .slice(0, maxItems)
          .map(task => (
            <ListItem key={task.id}>
              <ListItemIcon>
                <Checkbox checked={task.completed} onChange={() => {onToggle(task.id);}}/>
              </ListItemIcon>
              <ListItemText primary={task.title} sx = {{ color: task.completed ? "text.secondary" : "text.primary", textDecoration: task.completed ? "line-through" : "none"}} />
            </ListItem>
          ))
        }
      </List>
      {tasks.length > maxItems && (
      <Typography>
        {footerText}
      </Typography>
      )}
    </Card>
  )
}