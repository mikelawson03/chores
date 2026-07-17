import { Box, MenuItem, Stack, TextField, Typography } from "@mui/material";


export default function DetailRowNumber({label, field, value, onValueChange, units}) {
    return(
        <Stack direction="row" sx={{ alignItems: "center"}}>
            <Box sx={{width: 175}}>
                <Typography variant="h5">
                    {label}
                </Typography>
            </Box>
            <Box>
                <TextField 
                    type="number" 
                    value={value} 
                    size="small" 
                    sx={{width: 90}} 
                    onChange={
                        (event) => onValueChange(field, Number(event.target.value))
                    } 
                    slotProps={{
                        htmlInput: {
                            min: 1,
                            max: 1440,
                            step: 1,
                        }
                    }}
                />
                
            </Box>
            {units && <Typography sx={{p:1}}>{units}</Typography>}
        </Stack>
    )
}