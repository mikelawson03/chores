import { Container } from "@mui/material";

export default function AppLayout({ children, maxWidth }) {
    return (
        <Container maxWidth={maxWidth}>
            {children}
        </Container>
    );
}