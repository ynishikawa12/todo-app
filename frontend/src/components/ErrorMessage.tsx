import { Box, Typography } from "@mui/material";
import { ApiResult } from "../types/types";

export const ApiResultMessage = ({ result }: { result: ApiResult }) => {
  return (
    <>
      {result.success ? (
        <Box display="flex" flexDirection="column" alignItems="center" mt={2}>
          <Typography color="success">{result.message}</Typography>
        </Box>
      ) : (
        <Box display="flex" flexDirection="column" alignItems="center" mt={2}>
          <Typography color="error">{result.message}</Typography>
        </Box>
      )}
    </>
  );
};
