import { Box, Stack } from "@mui/material";
import CalendarGrid from "./CalendarGrid";
import TaskShelf from "./TaskShelf";

export default function CalendarContent({ currentDate, days, maxDayItems, onOverflowClick, openTaskDetails, monthlyTasks, weeklyTasks }) {
    return (
        <Stack direction="row" sx={{
            flex: 1,
            minHeight: 0
        }}>
            <CalendarGrid 
                currentDate={currentDate} 
                days={days} 
                maxDayItems={maxDayItems} 
                openTaskDetails={openTaskDetails}
                onOverflowClick={onOverflowClick}
            />
            <TaskShelf 
                weeklyTasks={weeklyTasks}
                monthlyTasks={monthlyTasks}
                openTaskDetails={openTaskDetails}
            />
        </Stack>
    )
}