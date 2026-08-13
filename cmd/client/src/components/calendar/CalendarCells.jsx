import { Box } from "@mui/material";
import dayjs from "dayjs";
import CalendarDay from "./CalendarDay";

export default function CalendarCells({currentDate, days, maxDayItems, onOverflowClick}) {
    return(
        <Box sx={{
            flex: 1,
            minHeight: 0,
            display: "grid",
            gridTemplateColumns:"repeat(7, 1fr)",
            gridTemplateRows: "repeat(6, 1fr)",
        }}>
            
            {days.map(({ day, tasks }, index) => (
                <CalendarDay 
                    key={day.format("YYYY-MM-DD")}
                    currentDate={currentDate} 
                    day={day} 
                    tasks={tasks}
                    maxDayItems={maxDayItems}
                    onOverflowClick={onOverflowClick}
                    isLastColumn={(index + 1) % 7 === 0}
                    isLastRow={index > 35}
                    isCurrentMonth={day.isSame(currentDate, "month")}
                    isCurrentDate={day.isSame(dayjs(), "day")}
                />
            ))}
            

        </Box>   
    )
}