import axios from "axios";
import { API_TAG, API_URL } from "../../../consts/consts";
import { ApiResponse } from "../../../types/types";
import { Tag } from "../types/tagType";

export const fetchTagsApi = async (): Promise<ApiResponse<Tag[] | null>> => {
  try {
    const response = await axios.get(API_URL + API_TAG);
    const tags = response.data as Tag[];
    return {
      response: tags,
      success: true,
      message: "タグを取得しました。",
    };
  } catch (error) {
    console.error(error);
    return {
      response: null,
      success: false,
      message: "タグの取得に失敗しました。",
    };
  }
};
