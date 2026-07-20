import { Box, Container, Stack } from "@mui/material";
import Sidebar from "../components/Sidebar";

export default function AppLayout({ children, maxWidth }) {
    return (
        <Container 
            maxWidth={false}
            sx = {{ minHeight: "100vh", p: 0, display: "flex", flexDirection: "column"}}
        >
            <Stack direction="row" sx= {{flex: 1, minHeight: 0}}>
                <Sidebar />
                <Box sx={{ overflow: "hidden", flex: 1, minHeight: 0 }}>
                    {children}
                </Box>
            </Stack>
            {/* {children} */}
        </Container>
    );
}