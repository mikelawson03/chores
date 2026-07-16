import { Box, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import EditChore from "../components/chores/EditChore";
import { EMPTY_CHORE_TEMPLATE } from "../constants/choreTemplate";
import { useEffect, useState } from "react";
import dayjs from "dayjs";
import { getChoreTemplates } from "../api/choreTemplates";

export default function Chores() {
  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [editChoreOpen, setEditChoreOpen] = useState(false);
  const [editChoreMode, setEditChoreMode] = useState(null);

  const [choreTemplates, setChoreTemplates]=useState([])

  useEffect(() => {
    async function loadChores() {
      const chores = await getChoreTemplates();
      setChoreTemplates(chores);
    }
    loadChores();
  },[])

  function openEditChore(choreTemplate) {
    setEditedChoreTemplate({...choreTemplate});
    setEditChoreMode("edit");
    setEditChoreOpen(true);
  }

  function closeEditChore(choreTemplate) {
    setEditChoreOpen(false);
  }

  function onChoreDetailChange(field, value) {
    setEditedChoreTemplate(previous => ({
      ...previous,
      [field]: value,
    }))
  }

  function editNewChore() {
    setEditedChoreTemplate({...EMPTY_CHORE_TEMPLATE});
    setEditChoreMode("create");
    setEditChoreOpen(true);
  }

  function createNewChore(){
    setChoreTemplates([...choreTemplates,
      {...editedChoreTemplate,
        created_at: dayjs().toISOString(),
        updated_at: dayjs().toISOString(),
        id: Math.max(...choreTemplates.map(chore => chore.id)) + 1
      }]
    )
  }

  // To-Do: 1. Make API Call to update chore in DB
  // 2: Log save event
  function saveChore() {
    setChoreTemplates(choreTemplates.map( template => {
      if (editedChoreTemplate.id === template.id) {
        return {
          ...editedChoreTemplate
        };
      }
      return template;
    }));
  }

  // TO-DO: 1. Make API Call to delete chore from DB
  // 2. Log delete event
  function deleteChore() {
    setChoreTemplates(previousChores => 
      previousChores.filter(chore => chore.id !== editedChoreTemplate.id)
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
          <ChoresTable choreTemplates={choreTemplates} openEditChore={openEditChore} editNewChore={editNewChore}/>
        </Stack>
          {editedChoreTemplate && 
            <EditChore 
              open={editChoreOpen} 
              chore={editedChoreTemplate} 
              closeEditChore={closeEditChore} 
              createNewChore={createNewChore}
              editChoreMode={editChoreMode}
              onChoreDetailChange={onChoreDetailChange} 
              saveChore={saveChore}
              deleteChore={deleteChore}
            />}
      </Stack>
)
}