import { Button, IconButton, Stack, Typography } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import { queryClient } from "../../query/queryClient";
import { useAuth } from "../../auth/useAuth";
import dayjs from "dayjs";

export default function PlannerToolbar({ weekStart, weekEnd, onPreviousWeek, onNextWeek, onResetWeek, currentDay }) {
  const { user } = useAuth();
  const isThisWeek = currentDay.isSame(dayjs(), "day")
  return (
    <Stack>
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
        {/* <Button 
          onClick={() => {
            queryClient.invalidateQueries({
              queryKey: ["assignments", user.id],
            });
          }}
        >
          Refresh Task Data
        </Button> */}
      </Stack>
      <Stack sx={{ alignItems: "center"}}>
        <Button 
          disabled={isThisWeek}
          onClick={onResetWeek}
        >
          Today
        </Button>
      </Stack>
    </Stack>
  )
}