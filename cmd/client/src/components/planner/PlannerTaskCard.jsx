import { Card, CardContent, Chip } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { clickableSurface } from "../../styles/surfaces";
import { formatDuration } from "../../utils/formatters";
import { useTaskStore } from "../../stores/taskStore";
import CompletionCheckbox from "../CompletionCheckbox";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";
import { CADENCES } from "../../constants/cadences";


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
    checkboxVisible = true
  }

  const openTaskDetails = useTaskStore(
    (state) => state.openTaskDetails
  );

  const ownerColor = HOUSEHOLD_USER_COLORS[task.userColorOption].bgColor

  return (
    <Card onClick={() => openTaskDetails(task)}
    sx= {[clickableSurface, { 
      position: "relative",
      overflow: "hidden",
      width: width,
      borderRadius: 1,
      backgroundColor: bgColor,
      maxWidth: "100%",
      minWidth: 0,

      "&::before": {
        content: '""',
        position: "absolute",
        left: 0,
        top: 6,
        bottom: 6,
        width: 4,
        borderRadius: "0 4px 4px 0",
        backgroundColor: ownerColor,
      }
      }]}>
      <CardContent sx={{ 
        py: 1, 
        px: 2,
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
            label={CADENCES[task.cadence].cadenceBadge}
            size="small"
            sx={{ 
              flexShrink: 0 , 
              bgcolor: CADENCES[task.cadence].backgroundColor,
              color: CADENCES[task.cadence].color,
            }}
          />
        </Stack>
        <Typography variant="body2" sx = {{ color: txtColor, textDecoration: txtDecoration, mb: 0.25, lineHeight: 1.15}}>
          {task.userDisplayName ? task.userDisplayName : "Unassigned"}
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