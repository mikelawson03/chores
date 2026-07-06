import { Stack, Typography } from "@mui/material";
import PlannerTaskCard from "./PlannerTaskCard";

export default function StagingSection({title, tasks, onToggleComplete}) {
    return (
    <Stack direction="column">
        <Typography variant="h6" gutterBottom>
            {title}
        </Typography>
        <Stack direction="row" sx={{ pb: 2, flexWrap: "wrap", gap: 1 }}>
            {tasks.map( task => (
                <PlannerTaskCard
                    key={task.id}
                    id={task.id}
                    title={task.title}
                    assignee={task.assignee}
                    duration={task.duration}
                    completed={task.completed}
                    onToggleComplete={onToggleComplete}
                    width={175}
                    />
            ))}
        </Stack>
    </Stack>
    )
}