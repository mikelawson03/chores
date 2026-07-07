import { Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";

export default function StagingSection({title, tasks, onToggleComplete, openTaskDetails}) {
    return (
    <Stack direction="column">
        <Typography variant="h6" gutterBottom>
            {title}
        </Typography>
        <Stack direction="row" sx={{ pb: 2, flexWrap: "wrap", gap: 1 }}>
            {tasks.map( task => (
                <PlannerTaskCard
                    key={task.id}
                    task={task}
                    onToggleComplete={onToggleComplete}
                    openTaskDetails={openTaskDetails}
                    width={175}
                    />
            ))}
        </Stack>
    </Stack>
    )
}