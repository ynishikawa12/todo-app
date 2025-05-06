import { Container } from "@mui/material";
import { TagForm } from "../features/tags/components/TagForm";
import { TagList } from "../features/tags/components/TagList";

export const Tag = () => {
  return (
    <Container maxWidth="sm">
      <TagForm />
      <TagList />
    </Container>
  );
};
