import { Checkbox, List, ListItem, ListItemIcon, ListItemText, Stack, Popover, Typography, Box, Chip, Button } from "@mui/material";
import { HOUSEHOLD_USER_COLORS } from "../constants/colorPalette";
import { CADENCES } from "../constants/cadences";
import { TASK_STATUSES } from "../constants/taskStatuses";
import { useFilterStore } from "../stores/filterStore";

export default function FilterMenu({ 
    users, 
    open, 
    handleClose, 
    anchorEl,
    filterConfig
}){
    const userIds = users.map(user => user.user.id);
    const cadences = Object.keys(CADENCES)
    const statuses = Object.keys(TASK_STATUSES)
    const resetFilters = useFilterStore(
        (state) => state.initializeFilters
    );

    const toggleFilterItem = useFilterStore(
        (state) => state.toggleFilterItem
    );

    const toggleAllFilters = useFilterStore(
        (state) => state.toggleAllFilters
    );

    const hiddenUserIds = useFilterStore(
        (state) => state.hiddenUserIds
      );

    const hiddenCadences = useFilterStore(
    (state) => state.hiddenCadences
    );
    
    const hiddenStatuses = useFilterStore(
    (state) => state.hiddenStatuses
    );

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
                    <Stack onClick={() => toggleAllFilters("users", userIds)} direction="row" sx={{alignItems: "center", p: 0}}>
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
                                onClick={() => toggleFilterItem("users", user.user.id)}
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
                    <Stack 
                        onClick={() => toggleAllFilters("cadences", cadences)} 
                        direction="row" 
                        sx={{alignItems: "center", p: 0}}
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
                            <ListItem key={cadence} sx={{p:0}}>
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenCadences.has(cadence)}
                                        onChange={() => toggleFilterItem("cadences", cadence)}
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
                                        // bgcolor: cadenceConfig.backgroundColor,
                                        py: 0.75,
                                        // px: 2,
                                        borderRadius: 1,
                                        // color: cadenceConfig.color,
                                    }}
                                    slotProps={{
                                        primary:{
                                            variant: "body2",
                                            sx: {
                                                fontWeight: 500,
                                            }
                                        }
                                    }}
                                    primary={cadenceConfig.label}
                                >
                                    
                                    
                                </ListItemText>
                            </ListItem>
                        ))}
                    </List>
                </Box>
                <Box>
                    <Stack 
                        onClick={() => toggleAllFilters("statuses", statuses)}
                        direction="row"
                        sx={{alignItems: "center", p:0}}
                    >
                        <Checkbox 
                            checked={hiddenStatuses.size === 0}
                            indeterminate={hiddenStatuses.size > 0 && hiddenStatuses.size < statuses.length}
                        />
                        <Typography variant="h6">
                            Statuses
                        </Typography>
                    </Stack>
                    <List sx={{pt: 0, pl: 4}}>
                        {Object.entries(TASK_STATUSES).map(([status, statusConfig]) => (
                            <ListItem key={status} sx={{p:0}}>
                                <ListItemIcon>
                                    <Checkbox 
                                        checked={!hiddenStatuses.has(status)}
                                        onChange={() => toggleFilterItem("statuses", status)}
                                    />
                                </ListItemIcon>
                                <ListItemText 
                                    sx={{
                                        py: 0.75,
                                        // px: 2,
                                        borderRadius:1,
                                    }}
                                    slotProps={{
                                        primary:{
                                            variant: "body2",
                                            sx: {
                                                fontWeight: 500,
                                            }
                                        }
                                    }}
                                    primary={statusConfig.label}
                                />
                            </ListItem>
                        ))}
                    </List>
                </Box>
                <Box>
                    <Button 
                        onClick={() => resetFilters(filterConfig)}
                    >Reset All</Button>
                </Box>
            </Stack>
        </Popover>
    )
}