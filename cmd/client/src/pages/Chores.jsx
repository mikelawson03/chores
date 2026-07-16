import { Box, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import ChoreDetails from "../components/chores/ChoreDetails";
import { useState } from "react";

export default function Chores() {
  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [choreDetailsOpen, setChoreDetailsOpen] = useState(false);
  

  const [choreTemplates, setChoreTemplates]=useState([
      {
        id: 1,
        name: "Make dinner",
        cadence: "daily",
        assignee: "",
        instructions: "",
        duration: 60,
        created_at: "2026-11-10T05:26:00",
        updated_at: "2026-11-10T05:26:00"
      },
      {
        id: 2,
        name: "Wash dishes",
        cadence: "daily",
        assignee: "",
        duration: 75,
        instructions: "",
        created_at: "2026-11-13T17:26:00",
        updated_at: "2026-07-13T17:26:00"
      },
      {
        id: 3,
        name: "Clean fish tank",
        cadence: "weekly",
        assignee: "Mike",
        duration: 20,
        instructions: "",
        created_at: "2026-06-05T17:26:00",
        updated_at: "2026-07-05T17:26:00"
      },
      {
        id: 4,
        name: "Mow yard",
        cadence: "weekly",
        assignee: "",
        duration: 90,
        instructions: "",
        created_at: "2026-07-14T16:11:00",
        updated_at: "2026-07-14T16:11:00"
      },
      {
        id: 5,
        name: "Feed dog",
        cadence: "daily",
        assignee: "",
        duration: 10,
        instructions: "Feed Ruby 1 cup of food in AM",
        created_at: "2026-07-15T13:48:00",
        updated_at: "2026-07-14T13:48:00"
      }
    ])

  function openChoreDetails(choreTemplate) {
    setEditedChoreTemplate({...choreTemplate});
    setChoreDetailsOpen(true);
  }

  function closeChoreDetails(choreTemplate) {
    setChoreDetailsOpen(false);
  }

  function onChoreDetailChange(field, value) {
    setEditedChoreTemplate(previous => ({
      ...previous,
      [field]: value,
    }))
  }

  // To-Do: 1. Make API Call to update chore in DB
  // 2: Log save event
  function saveChore(editedChore) {
    setChoreTemplates(choreTemplates.map( template => {
      if (editedChore.id === template.id) {
        return {
          ...editedChore
        };
      }
      return template;
    }));
  }

  // TO-DO: 1. Make API Call to delete chore from DB
  // 2. Log delete event
  function deleteChore(editedChore) {
    setChoreTemplates(previousChores => 
      previousChores.filter(chore => chore.id !== editedChore.id)
    );
  }

  return (
      <Stack spacing={4} sx={{
          flex: 1,
          alignItems: "center",
        }}
      >
        <Box sx={{
            width: "100%", 
            pb: 3, 
            borderBottom: 1, 
            borderColor: "divider", 
            display:"flex", 
            flexDirection: "column", 
            alignItems: "center"
          }}
        >
          <PageHeader title="Chore Management" />
        </Box>
        <Stack direction="column" sx={{flex: 1, width: "95%"}}>
          <ChoresTable choreTemplates={choreTemplates} openChoreDetails={openChoreDetails}/>
        </Stack>
          {editedChoreTemplate && 
            <ChoreDetails 
              open={choreDetailsOpen} 
              chore={editedChoreTemplate} 
              closeChoreDetails={closeChoreDetails} 
              onChoreDetailChange={onChoreDetailChange} 
              saveChore={saveChore}
              deleteChore={deleteChore}
            />}
      </Stack>
)
}