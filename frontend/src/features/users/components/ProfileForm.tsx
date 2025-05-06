import { Box, Button, Typography } from "@mui/material";
import { useActionState } from "react";
import { ApiResultMessage } from "../../../components/ErrorMessage";
import { updateProfileApi } from "../api/updateProfileApi";
import { ApiResult } from "../../../types/types";
import { useParams } from "react-router-dom";
import { useUser } from "../hooks/useUser";
import { EditableUserField } from "./EditableUserField";

const action = async (_: ApiResult, formData: FormData): Promise<ApiResult> => {
  return updateProfileApi({
    id: Number(formData.get("id")) || 0,
    name: formData.get("name") as string,
    email: formData.get("email") as string,
  });
};

export const ProfileForm = () => {
  const [result, updateProfile, isPending] = useActionState<
    ApiResult,
    FormData
  >(action, { success: false, message: "" });

  const { userId } = useParams<{ userId: string }>();

  const userResponse = useUser({ id: userId });
  const user = userResponse.response;

  return (
    <>
      <Box display="flex" flexDirection="column" alignItems="center" mt={8}>
        <Typography variant="h4" gutterBottom>
          Profile
        </Typography>
        <form action={updateProfile}>
          <input type="hidden" name="id" value={userId} />
          <EditableUserField
            label="Name"
            fieldName="name"
            value={user?.name || ""}
          />
          <EditableUserField
            label="Email"
            fieldName="email"
            value={user?.email || ""}
          />
          <Button
            type="submit"
            variant="contained"
            color="primary"
            sx={{ mt: 2 }}
            disabled={isPending}
          >
            {isPending ? "Loading..." : "Update"}
          </Button>
        </form>
      </Box>

      <ApiResultMessage result={result} />
    </>
  );
};
