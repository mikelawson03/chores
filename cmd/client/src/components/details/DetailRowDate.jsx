import { Box, Stack, Typography } from "@mui/material";
import { DatePicker } from "@mui/x-date-pickers/DatePicker";
import dayjs from "dayjs";

export default function DetailRowDate({label, field, value, maxDate, onValueChange, validateChoreField, error, required = false}) {
    return(
        <Stack>
            <Stack direction="row" sx={{ alignItems: "center"}}>
                <Box sx={{width: 175}}>
                    <Typography variant="h5">
                        {label}
                        {required && (
                            <Box 
                                component = "span"
                                sx={{
                                    color: "text.secondary",
                                    fontSize: "0.8em",
                                    verticalAlign: "top",
                                }}
                            >
                                *
                            </Box>
                        )}
                    </Typography>
                </Box>
                <Box>
                    <DatePicker
                        value={dayjs(value) ?? null} 
                        size="small" 
                        sx={{width: 175}} 
                        onChange={(newValue) => onValueChange(field, newValue ? newValue.startOf("day").toISOString() : null)} 
                        onBlur={() => validateChoreField(field, value)}
                        maxDate={dayjs(maxDate)}
                        slotProps={{
                            popper: {
                                sx: {
                                    zIndex: (theme) => theme.zIndex.modal + 2,
                                }
                            }
                        }}
                    />
                </Box>
            </Stack>
            <Typography 
                variant="body2"
                sx={{
                    color: "error.main",
                }}
            >
                {error && error}
            </Typography>
        </Stack>
    )
}