import { Box, Stack, Typography } from "@mui/material";

export default function DetailRow({label, value}) {
    return(
        <Stack direction="row" sx={{ alignItems: "center"}}>
            <Box sx={{width: 175}}>
                <Typography variant="h5">
                    {label}
                </Typography>
            </Box>
            <Box>
                <Typography variant="body1">
                    {value}
                </Typography>
            </Box>
        </Stack>
    )
}