import { Container } from "@mui/material";

export default function AppLayout({ children, maxWidth }) {
    return (
        <Container 
            maxWidth={maxWidth}
            sx = {{ height: "100vh", p: 0}}
        >
            {children}
        </Container>
    );
}