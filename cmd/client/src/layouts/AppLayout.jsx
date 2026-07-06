import { Container } from "@mui/material";

export default function AppLayout({ children }) {
    return (
        <Container maxWidth="lg">
            {children}
        </Container>
    );
}