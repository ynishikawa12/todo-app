import { useActionState } from "react";
import { Box, Typography } from "@mui/material";
import { Button, TextField } from "@mui/material";
import { ApiResult } from "../../../types/types";
import { createTagApi } from "../api/createTagApi";
import { ApiResultMessage } from "../../../components/ErrorMessage";

const tagName = "tagName";

const action = async (_: ApiResult, fromData: FormData): Promise<ApiResult> => {
  return await createTagApi({ name: fromData.get(tagName) as string });
};

export const TagForm = () => {
  const [result, createTag, isPending] = useActionState<ApiResult, FormData>(
    action,
    { success: false, message: "" }
  );

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mt={8}>
        <Typography variant="h4" gutterBottom>
          Create Tag
        </Typography>
        <form action={createTag}>
          <Box display="flex" flexDirection={"row"} alignItems="center">
            <TextField
              label="Tag"
              variant="outlined"
              fullWidth
              margin="normal"
              name={tagName}
            />
            <Button
              type="submit"
              variant="contained"
              color="primary"
              sx={{ ml: 2 }}
              disabled={isPending}
            >
              {isPending ? "Loading..." : "Create"}
            </Button>
          </Box>
        </form>
      </Box>

      <ApiResultMessage result={result} />
    </>
  );
};
