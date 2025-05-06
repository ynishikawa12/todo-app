import { Box, Typography } from "@mui/material";
import { useTags } from "../hooks/useTags";

export const TagList = () => {
  const tagsRes = useTags();
  const tags = tagsRes.response;

  const DisplayTags = () => {
    if (!tags) {
      return null;
    }
    return (
      <>
        {tags.map((tag) => (
          <Typography key={tag.id} variant="h6">
            {tag.name}
          </Typography>
        ))}
      </>
    );
  };

  return (
    <Box display="flex" flexDirection="column" alignItems="center" mt={1}>
      <DisplayTags />
    </Box>
  );
};
