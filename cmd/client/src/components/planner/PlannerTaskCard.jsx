import { Card, CardContent, Chip } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { formatDuration } from "../../utils/formatters";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";
import { deepPurple, deepOrange, cyan } from "@mui/material/colors";


export default function PlannerTaskCard({ task, width }) {
  let bgColor;
  let txtColor;
  let txtDecoration;
  let checkboxVisible;
  let cadenceBadge;

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
    checkboxVisible = true
  }

  switch (task.cadence){
    case "daily":
      cadenceBadge = "D";
      break;
    case "weekly":
      cadenceBadge = "W";
      break;
    case "monthly":
      cadenceBadge = "M";
      break;
  }

  const cadenceColors = {
    daily: {
      backgroundColor: cyan[50],
      color: cyan[800],
    },
    weekly: {
      backgroundColor: deepOrange[50],
      color: deepOrange[800],
    },
    monthly: {
      backgroundColor: deepPurple[50],
      color: deepPurple[800],
    },
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
      maxWidth: "100%",
      minWidth: 0,
      }]}>
      <CardContent sx={{ 
        p: 0.5, 
        "&:last-child": 
          { pb: 0.5 }, 
        lineHeight: 1, 
        minWidth: 0,
      }}>
        <Stack direction="row">
          <Typography variant="body1"  
            noWrap
            sx = {{
              textDecoration: txtDecoration, 
              color: txtColor,
              overflow: "hidden",
              textOverflow: "ellipsis",
              flex: 1,
              minWidth: 0,
              }}>
              {task.templateName}
          </Typography>
          <Chip 
            label={cadenceBadge}
            size="small"
            sx={{ 
              flexShrink: 0 , 
              ...cadenceColors[task.cadence]
            }}
          />
        </Stack>
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
            visible={checkboxVisible}
            size="small"
          />
        </Stack>
      </CardContent>
    </Card>
    );
}