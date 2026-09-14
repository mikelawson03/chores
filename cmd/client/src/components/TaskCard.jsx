import { Card, CardContent, Chip } from "@mui/material";
import { Typography } from "@mui/material";
import { Stack } from "@mui/material";
import { clickableSurface } from "../styles/surfaces";
import { formatDuration } from "../utils/formatters";
import { useTaskStore } from "../stores/taskStore";
import CompletionCheckbox from "./CompletionCheckbox";
import { CADENCES } from "../constants/cadences";
import { HOUSEHOLD_USER_COLORS } from "../constants/colorPalette";

export default function TaskCard({ task }) {
  const openTaskDetails = useTaskStore(
    (state) => state.openTaskDetails
  )

  const ownerColor = HOUSEHOLD_USER_COLORS[task.userColorOption]?.bgColor ?? null;

  return (
    <Card 
      onClick={() => openTaskDetails(task)}
      sx= {[clickableSurface, { 
        position: "relative",
        overflow: "hidden",
        width: "100%",
        borderRadius: 2,
        p: 1,
        backgroundColor: task.completed ? "grey.200" : "background.paper",
      
        "&::before": {
          content: '""',
          position: "absolute",
          left: 0,
          top: 1,
          bottom: 1,
          width: 4,
          borderRadius: "0 4px 4px 0",
          backgroundColor: ownerColor,
        }
      }]}>
        <CardContent sx={{ 
        py: 1, 
        px: 1,
        "&:last-child": 
          { pb: 0.5 }, 
        lineHeight: 1, 
        minWidth: 0,
        }}>
          <Stack 
          direction="row" 
          sx = {{ width: "100%", justifyContent:"space-between"  }}
          >
            <Typography variant="h6" 
            sx = {{
              textDecoration: task.completed ? "line-through" : "none", 
              color: task.completed ? "text.secondary" : "text.primary"
              }}>
              {task.templateName}
            </Typography>
            <Chip
              label={CADENCES[task.cadence].cadenceBadge}
              size="small"
              sx={{
                flexShrink: 0,
                bgcolor: CADENCES[task.cadence].backgroundColor,
                color: CADENCES[task.cadence].color,
              }}
            />
          </Stack>
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
      </CardContent>
    </Card>
    );
}