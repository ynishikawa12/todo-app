import { useEffect, useState } from "react";
import { ApiResponse } from "../../../types/types";
import { User } from "../types/userType";
import { fetchUserApi } from "../api/fetchUserApi";

export const useUser = ({ id }: { id: string | undefined }) => {
  const [user, setUser] = useState<ApiResponse<User | null>>({
    response: null,
    success: false,
    message: "",
  });
  useEffect(() => {
    if (!id) return;
    fetchUserApi(id).then((response) => {
      setUser(response);
    });
  }, [id]);

  return user;
};
