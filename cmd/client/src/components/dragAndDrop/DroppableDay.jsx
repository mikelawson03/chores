import { useDroppable } from "@dnd-kit/react";
import { Box } from "@mui/material";

export default function DroppableDay({ day, children }) {
    const { ref } = useDroppable({
        id: day.format("YYYY-MM-DD"),
    });

    return (
        <Box 
            ref={ref}
            sx={{
                flex: 1,
            }}
        >
            {children}
        </Box>
    )
}