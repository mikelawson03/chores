import { Checkbox, List, ListItem, ListItemIcon, ListItemText, Stack, Popover, Typography, Box } from "@mui/material";
import { HOUSEHOLD_USER_COLORS } from "../constants/colorPalette";
import { CADENCES } from "../constants/cadences";

export default function FilterMenu({ users, hiddenUserIds, toggleUser, hiddenCadences, toggleCadence, open, handleClose, anchorEl, hideAllUsers, showAllUsers, hideAllCadences, showAllCadences }){
    function handleUsersHeaderClick() {
        if (hiddenUserIds.size === 0) {
            hideAllUsers();
        } else {
            showAllUsers();
        }

    };

    function handleCadencesHeaderClick() {
        if (hiddenCadences.size === 0) {
            hideAllCadences();
        } else {
            showAllCadences();
        }
    };

    return (
        <Popover 
          open={open}
          onClose={handleClose}
          anchorEl={anchorEl}
          anchorOrigin={{
            vertical: 'bottom',
            horizontal: 'left',
          }}
          slotProps={{
            paper: {
                sx: {
                    width: 350,
                    px: 1,
                },
            },
          }}
        >
            <Stack sx={{ p:2, pl: 1 }}>
                <Box sx={{pb: 2}}>
                    <Stack onClick={handleUsersHeaderClick} direction="row" sx={{alignItems: "center", p: 0}}>
                        <Checkbox 
                            checked={hiddenUserIds.size === 0}
                            indeterminate={hiddenUserIds.size > 0 && hiddenUserIds.size < users.length}
                        />
                        <Typography variant="h6">
                            Users
                        </Typography>
                    </Stack>
                    <List sx={{pt: 0, pl: 4}}>
                        {users.map(user => (
                            <ListItem 
                                key={user.user.id} 
                                onClick={() => toggleUser(user.user.id)}
                                sx={{p:0,}}>
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenUserIds.has(user.user.id)}
                                    />
                                </ListItemIcon>
                                <ListItemText 
                                    sx={{
                                        bgcolor: HOUSEHOLD_USER_COLORS[user.colorOption].bgColor, 
                                        py: 0.75, 
                                        px: 2,
                                        borderRadius: 1,
                                        color: HOUSEHOLD_USER_COLORS[user.colorOption].textColor,
                                    }}
                                    slotProps={{
                                        primary: {
                                            variant: "body2",
                                            sx: {
                                                fontWeight: 500,
                                            }
                                        }
                                    }}
                                    primary={user.displayName}
                                />
                            </ListItem>
                        ))}
                    </List>
                </Box>
                <Box>
                    <Stack onClick={handleCadencesHeaderClick} direction="row" sx={{alignItems: "center", p: 0}}>
                        <Checkbox 
                            checked={hiddenCadences.size === 0}
                            indeterminate={hiddenCadences.size > 0 && hiddenCadences.size < Object.keys(CADENCES).length}
                        />
                        <Typography variant="h6">
                            Cadences
                        </Typography>
                    </Stack>
                    <List sx={{pt: 0, pl: 4}}>
                        {Object.entries(CADENCES).map(([cadence, cadenceConfig]) => (
                            <ListItem key={cadence} sx={{p:0}}>
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenCadences.has(cadence)}
                                        onChange={() => toggleCadence(cadence)}
                                    />
                                </ListItemIcon>
                                <ListItemText 
                                    sx={{
                                        bgcolor: cadenceConfig.backgroundColor,
                                        py: 0.75,
                                        px: 2,
                                        borderRadius: 1,
                                        color: cadenceConfig.color,
                                    }}
                                    slotProps={{
                                        primary:{
                                            variant: "body2",
                                            sx: {
                                                fontWeight: 500
                                            }
                                        }
                                    }}
                                    primary={cadenceConfig.label}
                                />
                            </ListItem>
                        ))}
                    </List>
                </Box>
            </Stack>
        </Popover>
    )
}