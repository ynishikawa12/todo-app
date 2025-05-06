import { useState, useEffect } from "react";
import { ApiResponse } from "../../../types/types";
import { Tag } from "../types/tagType.ts";
import { fetchTagsApi } from "../api/fetchTagsApi.ts";

export const useTags = (): ApiResponse<Tag[] | null> => {
  const [tags, setTags] = useState<ApiResponse<Tag[] | null>>({
    response: null,
    success: false,
    message: "",
  });
  useEffect(() => {
    fetchTagsApi().then((response) => {
      setTags(response);
    });
  }, []);

  return tags;
};
