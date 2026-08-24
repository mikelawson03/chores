import { useMutation } from "@tanstack/react-query";
import { queryClient } from "../query/queryClient";
import { toggleTaskCompletion } from "../utils/assignmentHelpers";
import { useAuth } from "../auth/useAuth";
import { Checkbox } from "@mui/material";
import { useNotificationStore } from "../stores/notificationStore";

export default function CompletionCheckbox({ checked, taskId }) {
    const { user } = useAuth();

    const showErrorNotification = useNotificationStore(
        (state) => state.showNotification
    )

    const toggleTaskCompletionMutation = useMutation({
        mutationFn: toggleTaskCompletion,
        onSuccess: () => {
        queryClient.invalidateQueries({
            queryKey: ["assignments", user.id],
        });
        },
        onError: (error) => {
        handleToggleCompletionError(error);
        }
    });

    function handleToggleCompletionError(error) {
        switch (error.status) {
            case 403:
                showErrorNotification("You do not have permission to modify this assignment");
                break;
            case 404:
                showErrorNotification("This assignment no longer exists.")
                break;
            default:
                showErrorNotification("An unexpected error occurred. Please try again")
                break;
        }
    }

    return(
        <Checkbox 
            checked={checked}
            size="medium" 
            onChange={() => toggleTaskCompletionMutation.mutate(taskId) } 
            onClick={(event) => {event.stopPropagation();}}
            sx={{ pr: 2, pl: 0 }}
        />
    )
}