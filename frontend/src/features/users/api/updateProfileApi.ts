import { API_URL, API_USER } from "../../../consts/consts";
import { ApiResult } from "../../../types/types";
import { User } from "../types/userType";

export const updateProfileApi = async (props: User): Promise<ApiResult> => {
  const body = {
    username: props.name,
    email: props.email,
  };
  try {
    await axios.put(API_URL + API_USER + "/" + props.id, body);
  } catch (error) {
    console.error(error);
    return { success: false, message: "プロフィールの更新に失敗しました。" };
  }

  return { success: true, message: "プロフィールを更新しました。" };
};
