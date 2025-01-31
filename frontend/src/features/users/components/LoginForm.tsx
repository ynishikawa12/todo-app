import { TextField, Button, Container, Typography, Box } from "@mui/material";
import { useLogin } from "../hooks/useLogin";

const LoginForm = () => {
  const {
    email,
    setEmail,
    password,
    setPassword,
    loading,
    error,
    handleSubmit,
  } = useLogin();

  return (
    <Container maxWidth="sm">
      <Box display="flex" flexDirection="column" alignItems="center" mt={8}>
        <Typography variant="h4" gutterBottom>
          Login
        </Typography>
        <TextField
          label="Email"
          variant="outlined"
          fullWidth
          margin="normal"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <TextField
          label="Password"
          type="password"
          variant="outlined"
          fullWidth
          margin="normal"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
        <Button
          variant="contained"
          color="primary"
          onClick={handleSubmit}
          sx={{ mt: 2 }}
          disabled={loading}
        >
          {loading ? "Loading..." : "Login"}
        </Button>
      </Box>

      {error && (
        <Box display="flex" flexDirection="column" alignItems="center" mt={2}>
          <Typography color="error">{error}</Typography>
        </Box>
      )}
    </Container>
  );
};

export default LoginForm;
