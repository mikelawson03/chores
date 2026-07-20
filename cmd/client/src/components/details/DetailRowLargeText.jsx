import { Box, MenuItem, Stack, TextField, Typography } from "@mui/material";


export default function DetailRowLargeText({label, field, value, onValueChange, units}) {
    return(
        <Stack direction="column" >
            <Typography variant="h5" sx={{ pb: 1}}>
                {label}
            </Typography>
            <TextField  
                variant="outlined"
                multiline={true}
                rows={5}
                value={value} 
                size="small" 
                
                onChange={(event) => onValueChange(field, event.target.value)} 
            />
        </Stack>
    )
}