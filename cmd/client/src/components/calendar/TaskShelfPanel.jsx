import { Stack } from "@mui/material";
import { useTaskStore } from "../../stores/taskStore";
import DraggableTask from "../dragAndDrop/DraggableTask";
import ShelfTask from "./ShelfTask";

export default function TaskShelfPanel({ tasks }) {
    const openTaskDetails = useTaskStore(
        (state) => state.openTaskDetails
    );

    return(
        <Stack spacing={1} 
            sx={{pt: 2}}
        >
            {tasks.map( task => 
                <DraggableTask task={task} key={task.id}>
                    {(isDragging) => <ShelfTask 
                        task={task}
                        isDragging={isDragging}
                        openTaskDetails={openTaskDetails}
                    />}
                </DraggableTask>
            )}
        </Stack>
    )
}