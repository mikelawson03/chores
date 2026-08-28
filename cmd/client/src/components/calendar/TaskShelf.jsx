import { Box } from "@mui/material";
import TaskShelfPanel from "./TaskShelfPanel";
import { useState } from "react";
import Tabs from '@mui/material/Tabs';
import Tab from '@mui/material/Tab';

export default function TaskShelf({ weeklyTasks, monthlyTasks }) {
    const [activeTab, setActiveTab] = useState(0);

    const handleTabChange = (event, newValue) => {
        setActiveTab(newValue);
    };

    return (
        <Box sx={{
            ml: 6, 
            p: 1, 
            width: 260, 
            border: 1,
            borderColor: "divider",
            borderRadius: 1
            }} 
        >
            <Tabs value={activeTab} onChange={handleTabChange} variant="fullWidth" >
                <Tab label={`weekly (${weeklyTasks.length})`} sx={{px: 0}} />
                <Tab label={`monthly (${monthlyTasks.length})`} sx={{px: 0}} />
            </Tabs>
            {activeTab === 0 && (
                <TaskShelfPanel
                    title="Weekly Tasks"
                    tasks={weeklyTasks}
                />
            )}

            {activeTab === 1 && (
                <TaskShelfPanel 
                    title="Monthly Tasks"
                    tasks={monthlyTasks}
                />
            )}
        </Box>
    )
}