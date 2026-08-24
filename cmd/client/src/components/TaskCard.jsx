import { Card } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { clickableSurface } from "../styles/surfaces";
import { formatDuration } from "../utils/formatters";
import { useTaskStore } from "../stores/taskStore";
import CompletionCheckbox from "./CompletionCheckbox";

export default function TaskCard({ task }) {
  const openTaskDetails = useTaskStore(
    (state) => state.openTaskDetails
  )
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
        <CompletionCheckbox
          checked={task.completed}
          taskId={task.id}
        />
      </Stack>
    </Card>
    );
}