import { Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";
import DraggableTask from "../dragAndDrop/DraggableTask";

export default function StagingSection({title, tasks, toggleTaskComplete }) {
    return (
    <Stack direction="column">
        <Typography variant="h6" gutterBottom>
            {title}
        </Typography>
        <Stack direction="row" sx={{ pb: 2, flexWrap: "wrap", gap: 1 }}>
            {tasks.map( task => (
                <DraggableTask task={task} key={task.id}>
                    <PlannerTaskCard
                        key={task.id}
                        task={task}
                        toggleTaskComplete={toggleTaskComplete}
                        width={175}
                        />
                </DraggableTask>
            ))}
        </Stack>
    </Stack>
    )
}