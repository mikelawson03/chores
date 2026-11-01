import { Box, Button, Stack, TextField, Typography } from "@mui/material"
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../auth/useAuth";


export default function Login() {

    const navigate = useNavigate();
    const { loginUser } = useAuth();

    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const [loginError, setLoginError] = useState("");

    const handleLogin = async (event) => {
        event.preventDefault();
        setLoginError("");

        try {
            await loginUser(username, password);
            navigate("/");
        } catch (err) {
            setLoginError(err.message)
        }
    }

    return (
        <Box 
            component="form" onSubmit={handleLogin} sx={{ width: "100%", height: "100vh", display: "flex", justifyContent:"center", alignItems:"center"}}
        >
            <Stack spacing={3} sx={{border: 1, borderRadius: 2, p:5, width: 300 }}>
                <Box>
                    <Typography>Username</Typography>
                    <TextField 
                        fullWidth
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                    />
                </Box>
                <Box>
                    <Typography>Password</Typography>
                    <TextField 
                        fullWidth
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                    />
                </Box>
                {loginError && (
                    <Typography variant="body2" color="error">
                        {loginError}
                    </Typography>
                )}
                <Button type="submit">Login</Button>
            </Stack>
        </Box>
    )
}