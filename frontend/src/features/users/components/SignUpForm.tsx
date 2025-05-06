import { useActionState } from "react";
import { Box, TextField, Typography, Button } from "@mui/material";
import { signUpApi } from "../api/signUpApi";
import { ApiResultMessage } from "../../../components/ErrorMessage";
import { ApiResult } from "../../../types/types";

const username = "username";
const email = "email";
const password = "password";

const action = async (_: ApiResult, formData: FormData): Promise<ApiResult> => {
  return await signUpApi({
    username: formData.get(username) as string,
    email: formData.get(email) as string,
    password: formData.get(password) as string,
  });
};

export const SignUpForm = () => {
  const [result, signUp, isPending] = useActionState<ApiResult, FormData>(
    action,
    { success: false, message: "" }
  );

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mt={8}>
        <Typography variant="h4" gutterBottom>
          SignUp
        </Typography>
        <form action={signUp}>
          <TextField
            label="Username"
            variant="outlined"
            fullWidth
            margin="normal"
            name={username}
          />
          <TextField
            label="Email"
            variant="outlined"
            fullWidth
            margin="normal"
            name={email}
          />
          <TextField
            label="Password"
            type="password"
            variant="outlined"
            fullWidth
            margin="normal"
            name={password}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{ mt: 2 }}
            disabled={isPending}
          >
            {isPending ? "Loading..." : "SignUp"}
          </Button>
        </form>
      </Box>

      <ApiResultMessage result={result} />
    </>
  );
};
