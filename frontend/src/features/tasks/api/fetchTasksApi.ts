import { ApiResponse } from "../../../types/types";
import { Task } from "../types/taskType";
import axios from "axios";
import { API_TASK, API_URL } from "../../../consts/consts";

export const fetchTodosApi = async (): Promise<ApiResponse<Task[] | null>> => {
  try {
    const response = await axios.get(API_URL + API_TASK);
    const tasks = response.data as Task[];
    return {
      response: tasks,
      success: true,
      message: "タスクを取得しました。",
    };
  } catch (error) {
    console.error(error);
    return {
      response: null,
      success: false,
      message: "タスクの取得に失敗しました。",
    };
  }
};
