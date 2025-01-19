import axios from "axios";
import { useState } from "react";
import { API_BASE_URL, TODO_URL } from "../consts/constants";
import { Box, Container } from "@mui/system";
import { Button, TextField, Typography } from "@mui/material";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import Link from "@mui/material/Link";
import { AlertMessage } from "../components/AlertMessage";

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axios.post(`${API_BASE_URL}/login`, {
        email,
        password,
      });
      setMessage('Login successfully!');
      // TODO: アプリ画面に遷移する
      navigate(TODO_URL + '/' + response.data.id);
    } catch (error) {
      setMessage('Failed to login.');
    }
  };

  const SignUpLink = () => {
    return (
      <Typography variant="body1" gutterBottom>
        Don't have an account? <Link component={RouterLink} to="/create-user">Sign up</Link>
      </Typography>
    );
  };

  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 4 }}>
        <Typography variant="h4" component={"h1"} gutterBottom>
          Login
        </Typography>
        <SignUpLink />
        <form onSubmit={handleSubmit}>
          <Box sx={{ mb: 2 }}>
            <TextField
              fullWidth
              label="Email"
              variant="outlined"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </Box>
          <Box sx={{ mb: 2 }}>
            <TextField
              fullWidth
              label="Password"
              variant="outlined"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </Box>
          <Button type="submit" variant="contained" color="primary">Login</Button>
        </form>
        <AlertMessage message={message} />
      </Box>
    </Container>
  )
}

export default Login;