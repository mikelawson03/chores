import { Box, MenuItem, Select, Stack, Typography } from "@mui/material";

export default function DetailRowSelect({label, field, value, options, onValueChange, validateChoreField, error, required = false}) {
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
                    <Select 
                        value={value ?? ""} 
                        size="small" 
                        sx={{width: 175}} 
                        onChange={(event) => onValueChange(field, event.target.value)} 
                        onBlur={() => validateChoreField(field, value)}
                        // renderValue={(selected) => {
                        //     const option = options.find(option => option.value === selected);
                        //     return option?.label ?? "";
                        // }}
                    >
                        {options.map(option => (
                            <MenuItem key={option.value} value={option.value}>
                                {option.label}
                            </MenuItem>
                        ))}
                    </Select>
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