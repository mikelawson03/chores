import { Button, IconButton, Stack, Typography } from "@mui/material";
import FilterListIcon from '@mui/icons-material/FilterList';
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import dayjs from "dayjs";
import FilterMenu from "../FilterMenu";
import { useState } from "react";

export default function CalendarToolbar({ currentDate, onPreviousMonth, onNextMonth, onResetDate, users, filterConfig }) {
    const isToday = currentDate.isSame(dayjs(), "day")

    const [calendarFiltersOpen, setCalendarFiltersOpen] = useState(false);
    const [anchorEl, setAnchorEl] = useState(null);

    const handleClick = (event) => {
        setCalendarFiltersOpen(true);
        setAnchorEl(event.currentTarget)
    }

    const handleClose = () => {
        setAnchorEl(null);
        setCalendarFiltersOpen(false);
    }

    return (
        <Stack>
            <Stack 
                direction="row" 
                spacing={2} 
                sx ={{justifyContent: "center", 
                    alignItems: "center"
                }} 
            >
                <IconButton onClick={onPreviousMonth} >
                    <ChevronLeftIcon fontSize="large" />
                </IconButton>
                <Typography variant="h5">
                    {currentDate.format("MMMM YYYY")}
                </Typography>
                <IconButton onClick={onNextMonth}>
                    <ChevronRightIcon fontSize="large" />
                </IconButton>
            </Stack>
            <Stack sx={{ alignItems: "center", position:"relative"}}>
                <IconButton 
                    onClick={handleClick}
                    sx={{
                        position:"absolute",
                        left: 0
                    }}
                >
                    <FilterListIcon />
                </IconButton>
                <FilterMenu
                  users={users}
                  open={calendarFiltersOpen}
                  handleClose={handleClose}
                  anchorEl={anchorEl}
                  filterConfig={filterConfig}
                />
                <Button 
                disabled={isToday}
                onClick={onResetDate}
                >
                Today
                </Button>
            </Stack>
        </Stack>
    )
}