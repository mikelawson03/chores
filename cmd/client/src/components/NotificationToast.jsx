import { Alert, Snackbar } from "@mui/material"
import { useNotificationStore } from "../stores/notificationStore"

export default function NotificationToast() {
    const notificationOpen = useNotificationStore(
        (state) => state.notificationOpen
    )

    const notificationMessage = useNotificationStore(
        (state) => state.notificationMessage
    )

    const closeNotification = useNotificationStore(
        (state) => state.closeNotification
    )

    return(
        <Snackbar 
            open={notificationOpen}
            autoHideDuration={5000}
            onClose={closeNotification}
            anchorOrigin={{ vertical: 'top', horizontal: 'center' }}
        >
            <Alert
                onClose={closeNotification}
                severity="error"
            >
                {notificationMessage}
            </Alert>
        </Snackbar>
    )
}