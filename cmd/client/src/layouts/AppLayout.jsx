import { Box, Container, Stack } from "@mui/material";
import Sidebar from "../components/Sidebar";

export default function AppLayout({ children, maxWidth }) {
    return (
        <Container 
            maxWidth={false}
            sx = {{ height: "100vh", p: 0}}
        >
            <Stack direction="row">
                <Sidebar />
                <Box sx={{ flexGrow: 1 }}>
                    {children}
                </Box>
            </Stack>
            {/* {children} */}
        </Container>
    );
}