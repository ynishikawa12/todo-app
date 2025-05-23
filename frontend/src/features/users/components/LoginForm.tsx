import { useActionState } from "react";
import { TextField, Button, Typography, Box } from "@mui/material";
import { loginApi } from "../api/loginApi";
import { ApiResultMessage } from "../../../components/ErrorMessage";
import { ApiResult } from "../../../types/types";

const action = async (_: ApiResult, formData: FormData): Promise<ApiResult> => {
  return loginApi({
    email: formData.get("email") as string,
    password: formData.get("password") as string,
  });
};

export const LoginForm = () => {
  const [result, login, isPending] = useActionState<ApiResult, FormData>(
    action,
    { success: false, message: "" }
  );

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mt={8}>
        <Typography variant="h4" gutterBottom>
          Login
        </Typography>
        <form action={login}>
          <TextField
            label="Email"
            variant="outlined"
            fullWidth
            margin="normal"
            name="email"
          />
          <TextField
            label="Password"
            type="password"
            variant="outlined"
            fullWidth
            margin="normal"
            name="password"
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{ mt: 2 }}
            disabled={isPending}
          >
            {isPending ? "Loading..." : "Login"}
          </Button>
        </form>
      </Box>

      <ApiResultMessage result={result} />
    </>
  );
};
