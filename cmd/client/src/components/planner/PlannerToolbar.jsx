import { IconButton, Stack, Typography } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";

export default function PlannerToolbar({ weekStart, weekEnd, onPreviousWeek, onNextWeek }) {
  return (
    <Stack direction="row" sx={{justifyContent: "center", alignItems: "center"}}  spacing={2} >
      <IconButton onClick={onPreviousWeek}>
        <ChevronLeftIcon fontSize="large"/>
      </IconButton>
      <Typography>
        {weekStart.format("MMM D")} - {weekEnd.format("MMM D")}
      </Typography>
      <IconButton onClick={onNextWeek}>
        <ChevronRightIcon fontSize="large" />
      </IconButton>
    </Stack>
  )
}