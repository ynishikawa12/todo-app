import axios from "axios";
import { API_URL, API_TAG } from "../../../consts/consts";
import { ApiResult } from "../../../types/types";

export const createTagApi = async ({
  name,
}: {
  name: string;
}): Promise<ApiResult> => {
  try {
    await axios.post(API_URL + API_TAG, { name });
  } catch (error) {
    console.error(error);
    return { success: false, message: "タグの作成に失敗しました。" };
  }

  return { success: true, message: "タグを作成しました。" };
};
