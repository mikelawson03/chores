import { Box, CircularProgress, Stack } from "@mui/material";
import ChoresTable from "../components/chores/ChoresTable";
import EditChore from "../components/chores/EditChore";
import { EDITABLE_CHORE_FIELDS, EMPTY_CHORE_TEMPLATE } from "../constants/choreTemplate";
import { useState } from "react";
import { getUsers } from "../utils/userHelpers";
import { choreTemplateValidationErrors, getChores, getChoreFieldError } from "../utils/choreHelpers";
import { useQuery, useMutation } from "@tanstack/react-query";
import { createChore, updateChore } from "../utils/choreHelpers";
import { queryClient } from "../query/queryClient";
import { parseApiError } from "../utils/errorHelpers";
import ChoresToolbar from "../components/chores/ChoresToolbar";
import { useAuth } from "../auth/useAuth";



export default function Chores() {
  const { user } = useAuth();

  const [editedChoreTemplate, setEditedChoreTemplate] = useState(null);
  const [editChoreOpen, setEditChoreOpen] = useState(false);
  const [editChoreMode, setEditChoreMode] = useState(null);
  const [templateErrors, setTemplateErrors] = useState({});
  const [hiddenUserIds, setHiddenUserIds] = useState(new Set());
  const [hiddenCadences, setHiddenCadences] = useState(new Set());

  const resetFilters = () => {
    setHiddenUserIds(new Set());
    setHiddenCadences(new Set());
  }

  const toggleFilterItem = (setter, value) => {
    setter(current => {
      const next = new Set(current);

      if (next.has(value)) {
        next.delete(value);
      } else {
        next.add(value);
      }
      return next;
    });
  }

  const toggleAllFilters = (setter, values) => {
    setter(current => {
      if (current.size === 0) {
        return new Set(values);
      } else {
        return new Set();
      }
    });
  }

  const usersQuery = useQuery({
    queryKey: ["users"],
    queryFn: () => getUsers(),
  })

  const choreTemplatesQuery = useQuery({
    queryKey: ["choreTemplates"],
    queryFn: () => getChores(),
  })

  const users = usersQuery.data ?? [];
  const userIds = users.map(user => user.user.id);
  const showMyTasks = () => {
          const idsToHide = userIds.filter(userId => userId !== user.user.id);
          setHiddenUserIds(new Set(["", ...idsToHide]));
      };
  
  const choreTemplates = choreTemplatesQuery.data ?? [];

  let filteredTemplates=choreTemplates
  filteredTemplates=filteredTemplates.filter(template => !hiddenUserIds.has(template.assignee));
  filteredTemplates=filteredTemplates.filter(template => !hiddenCadences.has(template.cadence));

  function openEditChore(choreTemplate) {
    setEditedChoreTemplate({
      ...choreTemplate,
      assignee: choreTemplate.assignee ? choreTemplate.assignee : "unassigned"
    });
    setEditChoreMode("edit");
    setTemplateErrors({});
    setEditChoreOpen(true);
  }

  function closeEditChore() {
    setTemplateErrors({});
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
    setTemplateErrors({});
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

    setTemplateErrors(newErrors)

    return Object.keys(newErrors).length === 0;
  }

  function handleSave() {
    console.log(editedChoreTemplate)
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
        setTemplateErrors(current => ({
          ...current,
          [validationError.field]: validationError.message,
        }))
        break;
      
      case 403:
        setTemplateErrors(current => ({
          ...current,
          "form": "You do not have permission to do this."
        }))
        break;

      case 404:
        setTemplateErrors(current => ({
          ...current,
          "form": "This resource no longer exists."
        }))  

        break;

      case 409:
        setTemplateErrors(current => ({
          ...current,
          "name": "Chore name already exists",
        }));
        break;
      
      default:
        setTemplateErrors(current => ({
          ...current,
          "form": "An unexpected error occurred. Please try again."}))
        break;
    }
  }

  function validateChoreField(field, value) {
    const error = getChoreFieldError(field, value);

    setTemplateErrors(current => ({
      ...current,
      [field]: error,
    }));
    return;
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
      <Stack spacing={2} sx={{
          flex: 1,
          alignItems: "center",
        }}
      >
        <ChoresToolbar 
          editNewChore={editNewChore}
          hiddenUserIds={hiddenUserIds}
          hiddenCadences={hiddenCadences}
          setHiddenUserIds={setHiddenUserIds}
          setHiddenCadences={setHiddenCadences}
          resetFilters={resetFilters}
          toggleFilterItem={toggleFilterItem}
          toggleAllFilters={toggleAllFilters}
          users={users}
          showMyTasks={showMyTasks}
        />
        <ChoresTable 
          choreTemplates={filteredTemplates} 
          openEditChore={openEditChore} 
          editNewChore={editNewChore} 
          users={users}
        />
        {editedChoreTemplate && 
          <EditChore 
            open={editChoreOpen} 
            chore={editedChoreTemplate} 
            closeEditChore={closeEditChore} 
            editChoreMode={editChoreMode}
            onChoreDetailChange={onChoreDetailChange} 
            users={users}
            errors={templateErrors}
            validateChoreField={validateChoreField}
            handleSave={handleSave}
            isSaving={isSaving}
          />}
      </Stack>
)
}