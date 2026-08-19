import { Box, Stack, TextField, Typography } from "@mui/material";


export default function DetailRowNumber({label, field, value, onValueChange, validateChoreField, error, units, required = false}) {
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
                                    verticalAlign: "top"
                                }}
                            >
                                *
                            </Box>
                        )}
                    </Typography>
                </Box>
                <Box>
                    <TextField 
                        type="number" 
                        value={value} 
                        size="small" 
                        sx={{width: 90}} 
                        onChange={(event) => {
                            const newValue = Number(event.target.value);
                            validateChoreField(field, newValue);
                            onValueChange(field, newValue);
                        }} 
                        onBlur={() => validateChoreField(field, value)}
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