import { Card, List, ListItem, ListItemIcon, ListItemText, Typography } from "@mui/material";
import { clickableText } from "../styles/typography";
import { useTaskStore } from "../stores/taskStore";
import CompletionCheckbox from "./CompletionCheckbox";
import { NavLink } from "react-router-dom";

export default function TaskListCard({ cardTitle, tasks, maxItems, footerText, route, subHead=" " }) {
  const openTaskDetails = useTaskStore(
    (state) => state.openTaskDetails
  )

  return (
    <Card sx={{
        width: "100%",
        borderRadius: 2,
        p: 2,
      }}
    >
      <Typography 
        variant="h6" 
      >
          {cardTitle.toUpperCase()}
      </Typography>
        <Typography
          sx ={{
            color: "text.secondary"
          }}
        >
          {subHead}
        </Typography>
      <List>
        {tasks
          .slice(0, maxItems)
          .map(task => (
            <ListItem key={task.id}>
              <ListItemIcon>
                <CompletionCheckbox
                  checked={task.completed}
                  taskId={task.id}
                />
              </ListItemIcon>
              <ListItemText onClick={() => openTaskDetails(task)} primary={task.templateName} sx = {[clickableText, { color: task.completed ? "text.secondary" : "text.primary", textDecoration: task.completed ? "line-through" : "none"}]} />
            </ListItem>
          ))
        }
      </List>
      {tasks.length > maxItems && (
      <Typography component={NavLink} to={route} sx={clickableText}>
        {footerText}
      </Typography>
      )}
    </Card>
  )
}