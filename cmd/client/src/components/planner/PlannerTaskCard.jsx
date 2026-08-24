import { Card, CardContent } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { formatDuration } from "../../utils/formatters";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";


export default function PlannerTaskCard({ task, width }) {
  let bgColor;
  let txtColor;
  let txtDecoration;
  let checkboxVisible;

  if (task.canceled) {
    bgColor = "grey.100";
    txtColor = "text.disabled";
    txtDecoration = "line-through";
    checkboxVisible = false
  } else if (task.completed) {
    bgColor = "grey.200";
    txtColor = "text.secondary";
    txtDecoration = "none";
    checkboxVisible = false
  } else {
    bgColor = "background.paper";
    txtColor = "text.primary";
    txtDecoration = "none";
    checkboxVisible = false
  }

  const openTaskDetails = useTaskStore(
    (state) => state.openTaskDetails
  );

  return (
    <Card onClick={() => openTaskDetails(task)}
    sx= {[clickableSurface, { 
      width: width,
      borderRadius: 1,
      backgroundColor: bgColor,
      }]}>
      <CardContent sx={{ p: 0.5, "&:last-child": { pb: 0.5 }, lineHeight: 1, }}>
        <Typography variant="body1"  
        sx = {{
          textDecoration: txtDecoration, 
          color: txtColor
          }}>
          {task.templateName}
        </Typography>
        <Typography variant="body2" sx = {{ color: txtColor, textDecoration: txtDecoration, mb: 0.25, lineHeight: 1.15}}>
          {task.userFirstName ? task.userFirstName : "Unassigned"}
        </Typography>
        <Stack 
          direction="row" 
          sx = {{ width: "100%", justifyContent:"space-between"  }}
        >
          <Typography variant="body2" sx={{ color: txtColor, textDecoration: txtDecoration }}>
            {formatDuration(task.duration)}
          </Typography>
          <CompletionCheckbox
            checked={task.completed}
            taskId={task.id}
          />
        </Stack>
      </CardContent>
    </Card>
    );
}