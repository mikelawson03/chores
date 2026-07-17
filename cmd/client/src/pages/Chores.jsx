import { Box, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import EditChore from "../components/chores/EditChore";
import { EMPTY_CHORE_TEMPLATE } from "../constants/choreTemplate";
import { useEffect, useState } from "react";
import { createChoreTemplate, deleteChoreTemplate, getChoreTemplates, updateChoreTemplate } from "../api/choreTemplates";

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

  async function createNewChore(){
    const newChore = await createChoreTemplate(editedChoreTemplate);
    setChoreTemplates([...choreTemplates, newChore]
    )
  }

  // To-Do:  Log save event
  async function saveChore() {
    const updatedChore = await updateChoreTemplate(editedChoreTemplate)
    setChoreTemplates(choreTemplates.map( template => {
      if (updatedChore.id === template.id) {
        return {
          ...updatedChore
        };
      }
      return template;
    }));
  }

  // TO-DO: 1. Make API Call to delete chore from DB
  // 2. Log delete event
  async function deleteChore() {
    await deleteChoreTemplate(editedChoreTemplate.id);
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