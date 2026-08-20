import { Box, CircularProgress, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import PageHeader from "../components/PageHeader";
import EditChore from "../components/chores/EditChore";
import { EDITABLE_CHORE_FIELDS, EMPTY_CHORE_TEMPLATE } from "../constants/choreTemplate";
import { useState } from "react";
import { getUsers } from "../utils/userHelpers";
import { choreTemplateValidationErrors, getChores, getChoreFieldError } from "../utils/choreHelpers";
import { useQuery, useMutation } from "@tanstack/react-query";
import { createChore, updateChore } from "../utils/choreHelpers";
import { queryClient } from "../query/queryClient";
import { parseApiError } from "../utils/errorHelpers";



export default function Chores() {
  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [editChoreOpen, setEditChoreOpen] = useState(false);
  const [editChoreMode, setEditChoreMode] = useState(null);
  const [errors, setErrors] = useState({});

  const usersQuery = useQuery({
    queryKey: ["users"],
    queryFn: () => getUsers(),
  })

  const choreTemplatesQuery = useQuery({
    queryKey: ["choreTemplates"],
    queryFn: () => getChores(),
  })

  const users = usersQuery.data ?? [];
  const choreTemplates = choreTemplatesQuery.data ?? [];

  function openEditChore(choreTemplate) {
    setEditedChoreTemplate({
      ...choreTemplate,
      assignee: choreTemplate.assignee ? choreTemplate.assignee : "unassigned"
    });
    setEditChoreMode("edit");
    setErrors({});
    setEditChoreOpen(true);
  }

  function closeEditChore() {
    setErrors({});
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
    setErrors({});
    setEditChoreOpen(true);
  }

  function validateChore(chore) {
    const newErrors = {};
    const choreFields = EDITABLE_CHORE_FIELDS;

    choreFields.forEach(field => {
      const error = getChoreFieldError(field, chore[field]);

      if (error) {
        newErrors[field] = error;
      }
    });

    setErrors(newErrors)

    return Object.keys(newErrors).length === 0;
  }

  function handleSave() {
    if (!validateChore(editedChoreTemplate)) {
      return;
    }

    if (editChoreMode === "create") {
      newChoreMutation.mutate(editedChoreTemplate);
    }
    else {
      editChoreMutation.mutate(editedChoreTemplate);
    }
  }

  const newChoreMutation = useMutation({
    mutationFn: createChore,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["choreTemplates"],
      });
      closeEditChore();
    },
    onError: (error) => {
      handleChoreError(error);
    }
  });

  const editChoreMutation = useMutation({
    mutationFn: updateChore,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: ["choreTemplates"],
      });
      closeEditChore();
    },
    onError: (error) => {
      handleChoreError(error);
    }
  });

  function handleChoreError(error) {
    let parsed, validationError
    switch (error.status) {
      case 400:
        parsed = parseApiError(error)
        validationError = choreTemplateValidationErrors[parsed.message]
        setErrors(current => ({
          ...current,
          [validationError.field]: validationError.message,
        }))
        break;
      
      case 403:
        setErrors(current => ({
          ...current,
          "form": "You do not have permission to do this."
        }))
        break;

      case 404:
        setErrors(current => ({
          ...current,
          "form": "This resource no longer exists."
        }))  

        break;

      case 409:
        setErrors(current => ({
          ...current,
          "name": "Chore name already exists",
        }));
        break;
      
      default:
        setErrors(current => ({
          ...current,
          "form": "An unexpected error occurred. Please try again."}))
        break;
    }
  }

  function validateChoreField(field, value) {
    const error = getChoreFieldError(field, value);

    setErrors(current => ({
      ...current,
      [field]: error,
    }));
  }

  const isSaving = newChoreMutation.isPending || editChoreMutation.isPending


  if (usersQuery.isPending || choreTemplatesQuery.isPending) {
    return( 
      <Box sx ={{
        width: "100%",
        height: "100vh",
        display: "flex",
        justifyContent: "center",
        alignItems: "center"
      }}>
        <CircularProgress />
      </Box>
  )
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
          <ChoresTable choreTemplates={choreTemplates} openEditChore={openEditChore} editNewChore={editNewChore} users={users}/>
          {editedChoreTemplate && 
            <EditChore 
              open={editChoreOpen} 
              chore={editedChoreTemplate} 
              closeEditChore={closeEditChore} 
              editChoreMode={editChoreMode}
              onChoreDetailChange={onChoreDetailChange} 
              users={users}
              errors={errors}
              validateChoreField={validateChoreField}
              handleSave={handleSave}
              isSaving={isSaving}
            />}
      </Stack>
)
}