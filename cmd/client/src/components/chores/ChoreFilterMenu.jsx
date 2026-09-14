import { Box, Button, Checkbox, Chip, List, ListItem, ListItemIcon, ListItemText, Popover, Stack, Typography } from "@mui/material";
import { CADENCES } from "../../constants/cadences";
import { HOUSEHOLD_USER_COLORS } from "../../constants/colorPalette";
import { grey } from "@mui/material/colors";

export default function FilterMenu({
    users,
    open,
    handleClose,
    anchorEl,
    hiddenUserIds,
    hiddenCadences,
    setHiddenUserIds,
    setHiddenCadences,
    toggleFilterItem,
    toggleAllFilters,
    resetFilters,
    showMyTasks
}){
    const userOptions=[
        ...users.map(user => ({
            id: user.user.id,
            displayName: user.displayName,
            colorOption: user.colorOption,
        })),
        {id: "", displayName: "Unassigned"},
    ]

    const userIds = users.map(user => user.user.id)
    const cadences = Object.keys(CADENCES);
    
    return (
        <Popover
            open={open}
            onClose={handleClose}
            anchorEl={anchorEl}
            anchorOrigin={{
                vertical: 'bottom',
                horizontal:'left',
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
            <Stack sx={{ p: 2, pl: 1 }}>
                <Box sx={{ob: 2}}>
                    <Stack
                        onClick={() => toggleAllFilters(setHiddenUserIds, userIds)}
                        direction="row"
                        sx={{
                            alignItems: "center",
                            p: 0,
                        }}
                    >
                        <Checkbox 
                            checked={hiddenUserIds.size === 0}
                            indeterminate={hiddenUserIds.size > 0 && hiddenUserIds.size < userOptions.length}
                        />
                        <Typography variant="h6">
                            Users
                        </Typography>
                    </Stack>
                    <List sx={{ pt: 0, pl: 4}}>
                        {userOptions.map(user => (
                            <ListItem
                                key={user.id}
                                onClick={() => toggleFilterItem(setHiddenUserIds, user.id)}
                                sx={{ p:0 }}
                            >
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenUserIds.has(user.id)}
                                    />
                                </ListItemIcon>
                                <ListItemText 
                                    sx={{
                                        bgcolor: HOUSEHOLD_USER_COLORS[user.colorOption]?.bgColor ?? grey[300],
                                        py: 0.75,
                                        px: 2,
                                        borderRadius: 1,
                                        color: HOUSEHOLD_USER_COLORS[user.colorOption]?.textColor ?? "#000",
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
                    <Stack
                        onClick={() => toggleAllFilters(setHiddenCadences, cadences)}
                        direction="row"
                        sx={{
                            alignItems: "center",
                            p: 0
                        }}
                    >
                        <Checkbox 
                            checked={hiddenCadences.size === 0}
                            indeterminate={hiddenCadences.size > 0 && hiddenCadences.size < cadences.length}
                        />
                        <Typography variant="h6">
                            Cadences
                        </Typography>
                    </Stack>
                    <List sx={{pt: 0, pl: 4}}>
                        {Object.entries(CADENCES).map(([cadence, cadenceConfig]) => (
                            <ListItem 
                                key={cadence} 
                                sx={{p: 0}}
                            >
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenCadences.has(cadence)}
                                        onChange={() => toggleFilterItem(setHiddenCadences, cadence)}
                                    />
                                </ListItemIcon>
                                <Chip 
                                    label={cadenceConfig.cadenceBadge}
                                    size="small"
                                    sx={{
                                        flexShrink: 0,
                                        bgcolor: cadenceConfig.backgroundColor,
                                        color: cadenceConfig.color,
                                        mr: 1,
                                    }}
                                />
                                <ListItemText 
                                    sx={{
                                        py: 0.75,
                                        borderRadius: 1,
                                    }}
                                    slotProps={{
                                        primary: {
                                            variant: "body2",
                                            sx: {
                                                fontWeight: 500,
                                            }
                                        }
                                    }}
                                    primary={cadenceConfig.label}
                                />
                            </ListItem>
                        ))}
                    </List>
                </Box>
                <Stack 
                    direction="row"
                    spacing={4}
                >
                    <Button
                        onClick={resetFilters}
                    >
                        Reset All
                    </Button>
                    <Button
                        onClick={showMyTasks}
                    >
                        Show My Tasks
                    </Button>
                </Stack>
            </Stack>

        </Popover>
    )
}