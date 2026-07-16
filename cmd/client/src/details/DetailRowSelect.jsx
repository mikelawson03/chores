import { Box, MenuItem, Select, Stack, Typography } from "@mui/material";

export default function DetailRowSelect({label, field, value, options, onValueChange}) {
    return(
        <Stack direction="row" sx={{ alignItems: "center"}}>
            <Box sx={{width: 175}}>
                <Typography variant="h5">
                    {label}
                </Typography>
            </Box>
            <Box>
                <Select value={value} options={options} size="small" sx={{width: 175}} onChange={(event) => onValueChange(field, event.target.value)}>
                    {options.map(option => (
                        <MenuItem key={option.value} value={option.value}>
                            {option.label}
                        </MenuItem>
                    ))}
                </Select>
            </Box>
        </Stack>
    )
}