import axios from "axios";
import { API_LOGIN, API_URL } from "../../../consts/consts";
import { ApiResult } from "../../../types/types";

type LoginProps = {
  email: string;
  password: string;
};

export const loginApi = async (props: LoginProps): Promise<ApiResult> => {
  try {
    await axios.post(API_URL + API_LOGIN, props);
    // TODO: set authentication token
  } catch (error) {
    console.error(error);
    return { success: false, message: "ログインに失敗しました。" };
  }

  return { success: true, message: "ログインしました。" };
};
