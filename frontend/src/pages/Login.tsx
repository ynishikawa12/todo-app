import { Container } from "@mui/material";
import { LoginForm } from "../features/users/components/LoginForm";

export const Login = () => {
  return (
    <Container maxWidth="sm">
      <LoginForm />
    </Container>
  );
};
