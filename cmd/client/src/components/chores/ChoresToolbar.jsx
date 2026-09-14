import { Box, Button, IconButton, Stack } from "@mui/material";
import FilterListIcon from "@mui/icons-material/FilterList";
import PageHeader from "../PageHeader";
import { useState } from "react";
import ChoreFilterMenu from "./ChoreFilterMenu"


export default function ChoresToolbar({ 
  editNewChore,
  hiddenUserIds,
  hiddenCadences,
  setHiddenUserIds,
  setHiddenCadences,
  resetFilters,
  toggleFilterItem,
  toggleAllFilters,
  users,
  showMyTasks
 }){

  const [choreFitersOpen, setChoreFiltersOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState(null);

  const handleClick = (event) => {
    setChoreFiltersOpen(true);
    setAnchorEl(event.currentTarget);
  }

  const handleClose = () => {
    setAnchorEl(null);
    setChoreFiltersOpen(false);
  }

  return(
    <Stack
      spacing={1}
      sx={{
        width: "95%",
        pt: 3,
      }}
    >
      <Box sx={{
        display:"flex", 
        flexDirection: "column", 
        alignItems: "center",
      }}
      >
        <PageHeader title="Chore Management" />
      </Box>
      <Stack 
        direction="row" 
        spacing={2}
        sx={{
          justifyContent: "space-between",
          alignItems: "center",
          position: "relative",
        }}
      >
        <Stack 
          direction="row" 
          spacing={2}
        >
          <IconButton 
            onClick={handleClick}
            sx={{
              position: "absolute", 
              left: 0,
              top: "75%",
              transform: "translateY(-50%)",
            }} 
          >
            <FilterListIcon fontSize="medium" />
          </IconButton>
          <ChoreFilterMenu 
            users={users}
            open={choreFitersOpen}
            handleClose={handleClose}
            anchorEl={anchorEl}
            hiddenUserIds={hiddenUserIds}
            hiddenCadences={hiddenCadences}
            setHiddenUserIds={setHiddenUserIds}
            setHiddenCadences={setHiddenCadences}
            toggleFilterItem={toggleFilterItem}
            toggleAllFilters={toggleAllFilters}
            resetFilters={resetFilters}
            showMyTasks={showMyTasks}
          />
        </Stack>
        <Button variant="contained" onClick={editNewChore}>New Chore Template</Button>
      </Stack>
    </Stack>
  )
}