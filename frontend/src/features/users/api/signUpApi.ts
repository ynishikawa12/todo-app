import { API_URL, API_USER } from "../../../consts/consts";
import axios from "axios";
import { ApiResult } from "../../../types/types";

type SignUpProps = {
  username: string;
  email: string;
  password: string;
};

export const signUpApi = async (props: SignUpProps): Promise<ApiResult> => {
  try {
    await axios.post(API_URL + API_USER, props);
  } catch (error) {
    console.error(error);
    return { success: false, message: "サインアップに失敗しました。" };
  }

  return { success: true, message: "サインアップに成功しました。" };
};
