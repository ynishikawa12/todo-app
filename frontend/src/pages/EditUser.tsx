import { useState, useEffect } from 'react';
import axios from 'axios';
import { Container, TextField, Button, Typography, Box } from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';
import { API_BASE_URL, TODO_URL } from '../consts/constants';
import { AlertMessage } from '../components/AlertMessage';
import TodoAppBar from '../components/TodoAppBar';

const EditUser = () => {
  const { id } = useParams<{ id: string }>();
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [message, setMessage] = useState('');

  const navigate = useNavigate();

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const response = await axios.get(`${API_BASE_URL}/users/${id}`);
        const user = response.data;
        setName(user.name);
        setEmail(user.email);
        setPassword(user.password);
      } catch (error) {
        setMessage('Failed to fetch user data.');
      }
    };

    fetchUser();
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await axios.put(`${API_BASE_URL}/users/${id}`, {
        name,
        email,
        password,
      });
      setMessage('User updated successfully!');
      navigate(TODO_URL + '/' + id);
    } catch (error) {
      setMessage('Failed to update user.');
    }
  };

  return (
    <>
      {id ? <TodoAppBar id={id} /> : null}
      <Container maxWidth="sm">
        <Box sx={{ mt: 4 }}>
          <Typography variant="h4" component="h1" gutterBottom>
            Edit User
          </Typography>
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
              Update User
            </Button>
          </form>
          <AlertMessage message={message} />
        </Box>
      </Container>
    </>
  );
};

export default EditUser;
