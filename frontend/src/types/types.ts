export type ApiResult = {
  success: boolean;
  message: string;
};

export type ApiResponse<T> = {
  success: boolean;
  message: string;
  response: T;
};
