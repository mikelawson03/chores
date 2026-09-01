import { Button, IconButton, Stack, Typography } from "@mui/material";
import ChevronLeftIcon from "@mui/icons-material/ChevronLeft";
import ChevronRightIcon from "@mui/icons-material/ChevronRight";
import dayjs from "dayjs";

export default function CalendarToolbar({ currentDate, onPreviousMonth, onNextMonth, onResetDate }) {
    const isToday = currentDate.isSame(dayjs(), "day")
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
            <Stack sx={{ alignItems: "center"}}>
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