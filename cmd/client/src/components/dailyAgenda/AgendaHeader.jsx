import { IconButton, Stack, Typography } from "@mui/material";
import FilterListIcon from '@mui/icons-material/FilterList';
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import FilterMenu from "../FilterMenu";
import { useState } from "react";

export default function AgendaHeader({ agendaDate, onNextAgendaDay, onPreviousAgendaDay, users, filterConfig }) {
    
    const [agendaFiltersOpen, setAgendaFiltersOpen] = useState(false);
    const [anchorEl, setAnchorEl] = useState(null);

    const handleClick = (event) => {
        setAgendaFiltersOpen(true);
        setAnchorEl(event.currentTarget)
    }

    const handleClose = () => {
        setAnchorEl(null);
        setAgendaFiltersOpen(false);
    }

    return(
        <Stack 
            direction="row" 
            sx={{
                justifyContent: "center",
                alignItems: "center",
                position:"relative",
                width:"100%",       
            }}
        >
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
                open={agendaFiltersOpen}
                handleClose={handleClose}
                anchorEl={anchorEl}
                filterConfig={filterConfig}
            />
                <IconButton onClick={onPreviousAgendaDay}>
                    <ChevronLeftIcon fontSize="large"  />
                </IconButton>
                <Typography variant="h6">
                    {agendaDate.format("MMMM DD YYYY")}
                </Typography>
                <IconButton onClick={onNextAgendaDay} >
                    <ChevronRightIcon fontSize="large"  />
                </IconButton>
        </Stack>
    )
}