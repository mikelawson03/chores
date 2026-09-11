import { Button, IconButton, Stack, Typography } from "@mui/material";
import FilterListIcon from '@mui/icons-material/FilterList';
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import dayjs from "dayjs";
import { useState } from "react";
import FilterMenu from "../FilterMenu";

export default function PlannerToolbar({ 
  weekStart, 
  weekEnd, 
  onPreviousWeek, 
  onNextWeek, 
  onResetWeek, 
  currentDay,
  users,
  filterConfig
}) {
  const isThisWeek = currentDay.isSame(dayjs(), "day")

  const [plannerFiltersOpen, setPlannerFiltersOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState(null);
  
  const handleClick = (event) => {
    setPlannerFiltersOpen(true);
    setAnchorEl(event.currentTarget)
  }

  const handleClose = () => {
    setAnchorEl(null);
    setPlannerFiltersOpen(false);
  }

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
      </Stack>
      <Stack sx={{ alignItems: "center", position: "relative"}}>
        <IconButton 
          onClick={handleClick}
          sx={{
            position: "absolute", 
            left: 0
          }} 
        >
          <FilterListIcon/>
        </IconButton>      
        <FilterMenu 
          users={users} 
          open={plannerFiltersOpen}
          handleClose={handleClose}
          anchorEl={anchorEl}
          filterConfig={filterConfig}
        />
           
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