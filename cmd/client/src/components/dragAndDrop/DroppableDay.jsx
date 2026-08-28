import { useDroppable } from "@dnd-kit/react";
import { Box } from "@mui/material";

export default function DroppableDay({ day, children }) {
    const { ref, isDropTarget } = useDroppable({
        id: day.format("YYYY-MM-DD"),
    });

    return (
        <Box 
            ref={ref}
            sx={{
                flex: 1,
                border: isDropTarget? 1 : 0,
                borderColor: "primary.main",
                transition: "border-color 0.15s ease",
            }}
        >
            {children}
        </Box>
    )
}