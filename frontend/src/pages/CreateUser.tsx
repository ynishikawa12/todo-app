import { useState } from 'react';
import { Link as RouterLink} from 'react-router-dom';
import axios from 'axios';
import { Container, TextField, Button, Typography, Box, Alert } from '@mui/material';
import { API_BASE_URL } from '../consts/constants';
import { AlertMessage } from '../components/AlertMessage';

const CreateUser = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.post(`${API_BASE_URL}/users`, {
        name,
        email,
        password,
      });
      setMessage('User created successfully!');
    } catch (error) {
      setMessage('Failed to create user.');
    }
  };

  const LoginLink = () => {
    return (
      <Typography variant="body1" gutterBottom>
        Do you have an account ? <RouterLink to="/">Login</RouterLink>
      </Typography>
    );
  };
  return (
    <Container maxWidth="sm">
      <Box sx={{ mt: 4 }}>
        <Typography variant="h4" component="h1" gutterBottom>
          Create User
        </Typography>
        <LoginLink />
        <form onSubmit={handleSubmit}>
          <Box sx={{ mb: 2 }}>
            <TextField
              fullWidth
              label="Name"
              variant="outlined"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />
          </Box>
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
          <Button variant="contained" color="primary" type="submit">
            Create User
          </Button>
        </form>
        <AlertMessage message={message} />
      </Box>
    </Container>
  );
};

export default CreateUser;
