import axios from "axios";
import { API_URL } from "../../../consts/consts";
import { User } from "../types/userType";
import { ApiResponse } from "../../../types/types";

export const fetchUserApi = async (
  userId: string
): Promise<ApiResponse<User | null>> => {
  try {
    const response = await axios.get(API_URL + `/users/${userId}`);
    const user = response.data as User;
    return {
      response: user,
      success: true,
      message: "ユーザー情報を取得しました。",
    };
  } catch (error) {
    console.error(error);
    return {
      response: null,
      success: false,
      message: "ユーザー情報の取得に失敗しました。",
    };
  }
};
